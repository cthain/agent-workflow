package workflow

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestDetectStates(t *testing.T) {
	tests := []struct {
		name, taskStatus, reviewStatus, extra string
		want                                  State
	}{
		{"ready", "", "", "", Ready},
		{"planned", "planned", "", "", Planned},
		{"in progress", "implementation-in-progress", "", "", InProgress},
		{"complete", "implementation-complete", "", "", Complete},
		{"changes requested", "implementation-complete", "changes-requested", "", ChangesRequested},
		{"approved", "implementation-complete", "approved", "", Approved},
		{"blocked", "implementation-complete", "blocked", "", Blocked},
		{"remediation in progress", "implementation-in-progress", "changes-requested", "remediates-review: review-01.md\n", InProgress},
		{"remediation complete", "implementation-complete", "changes-requested", "remediates-review: review-01.md\n", Complete},
		{"blocked route code", "implementation-in-progress", "blocked", "blocked-review-resolution:\n  review: review-01.md\n  route: code\n  recorded-by: developer\n  evidence:\n    - evidence.md\n", InProgress},
		{"blocked route review", "implementation-complete", "blocked", "blocked-review-resolution:\n  review: review-01.md\n  route: review\n  recorded-by: task-planner\n  evidence:\n    - evidence.md\n", Complete},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := fixture(t, tt.taskStatus, tt.reviewStatus, tt.extra)
			r := Detect(root)
			if r.State != tt.want {
				t.Fatalf("state = %s, want %s; diagnostics: %v", r.State, tt.want, r.Diagnostics)
			}
		})
	}
}

func TestDetectArchivedGitTask(t *testing.T) {
	extra := "git:\n  enabled: true\n  trunk: trunk\n  task-branch: concoct/app-001-demo\n  base: abc123\n  archive-commit: def456\n  status: archived\n"
	root := fixture(t, "implementation-complete", "approved", extra)
	got := Detect(root)
	if got.State != Archived || got.Next != "concoct integrate" {
		t.Fatalf("got %#v", got)
	}
}

func TestDetectInterruptedDeliveryIsNotReady(t *testing.T) {
	root := fixture(t, "", "", "")
	write(t, filepath.Join(root, ".git/concoct/integrations/APP-001.yaml"), "phase: delivered\n")
	got := Detect(root)
	if got.State != Integrated || got.Next != "concoct integrate --continue" {
		t.Fatalf("got %#v", got)
	}
}

func TestDetectInvalidEvidence(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(t *testing.T, root string)
	}{
		{"missing notes", func(t *testing.T, r string) { os.Remove(filepath.Join(r, ".concoct/current/notes.md")) }},
		{"roadmap mismatch", func(t *testing.T, r string) {
			write(t, filepath.Join(r, ".concoct/current/task-plan.md"), task("OTHER-001", "planned", ""))
		}},
		{"unknown status", func(t *testing.T, r string) {
			write(t, filepath.Join(r, ".concoct/current/task-plan.md"), task("APP-001", "mystery", ""))
		}},
		{"unknown roadmap status", func(t *testing.T, r string) {
			write(t, filepath.Join(r, ".concoct/roadmap.md"), roadmap("mystery"))
		}},
		{"missing task title", func(t *testing.T, r string) {
			path := filepath.Join(r, ".concoct/current/task-plan.md")
			data, _ := os.ReadFile(path)
			write(t, path, strings.Replace(string(data), "title: Demo task\n", "", 1))
		}},
		{"missing task created", func(t *testing.T, r string) {
			path := filepath.Join(r, ".concoct/current/task-plan.md")
			data, _ := os.ReadFile(path)
			write(t, path, strings.Replace(string(data), "created: 2026-01-01\n", "", 1))
		}},
		{"missing task updated", func(t *testing.T, r string) {
			path := filepath.Join(r, ".concoct/current/task-plan.md")
			data, _ := os.ReadFile(path)
			write(t, path, strings.Replace(string(data), "updated: 2026-01-01\n", "", 1))
		}},
		{"missing impact rationale", func(t *testing.T, r string) {
			path := filepath.Join(r, ".concoct/current/task-plan.md")
			data, _ := os.ReadFile(path)
			write(t, path, strings.Replace(string(data), "  rationale: Demo impact\n", "", 1))
		}},
		{"missing review created", func(t *testing.T, r string) {
			path := filepath.Join(r, ".concoct/current/review-01.md")
			data, _ := os.ReadFile(path)
			write(t, path, strings.Replace(string(data), "created: 2026-01-01\n", "", 1))
		}},
		{"missing review persona", func(t *testing.T, r string) {
			path := filepath.Join(r, ".concoct/current/review-01.md")
			data, _ := os.ReadFile(path)
			write(t, path, strings.Replace(string(data), "persona: reviewer\n", "", 1))
		}},
		{"review gap", func(t *testing.T, r string) {
			os.Rename(filepath.Join(r, ".concoct/current/review-01.md"), filepath.Join(r, ".concoct/current/review-02.md"))
		}},
		{"stale remediation", func(t *testing.T, r string) {
			write(t, filepath.Join(r, ".concoct/current/task-plan.md"), task("APP-001", "implementation-in-progress", "remediates-review: review-02.md\n"))
		}},
		{"unsafe blocker evidence", func(t *testing.T, r string) {
			write(t, filepath.Join(r, ".concoct/current/task-plan.md"), task("APP-001", "implementation-in-progress", "blocked-review-resolution:\n  review: review-01.md\n  route: code\n  recorded-by: developer\n  evidence: [../outside]\n"))
		}},
		{"delivered active task", func(t *testing.T, r string) { write(t, filepath.Join(r, ".concoct/roadmap.md"), roadmap("delivered")) }},
		{"malformed task metadata", func(t *testing.T, r string) {
			write(t, filepath.Join(r, ".concoct/current/task-plan.md"), "---\nid: [\n---\n")
		}},
		{"multiple review outcomes", func(t *testing.T, r string) {
			path := filepath.Join(r, ".concoct/current/review-01.md")
			data, _ := os.ReadFile(path)
			write(t, path, strings.Replace(string(data), "`changes-requested`", "`changes-requested` and `approved`", 1))
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := fixture(t, "implementation-complete", "changes-requested", "")
			tt.mutate(t, root)
			got := Detect(root)
			if got.State != Invalid || len(got.Diagnostics) == 0 {
				t.Fatalf("got %#v", got)
			}
		})
	}
}

func TestLaterReviewSupersedesRecoveryMetadata(t *testing.T) {
	tests := []struct {
		name, firstOutcome, extra string
	}{
		{"remediation", "changes-requested", "remediates-review: review-01.md\n"},
		{"blocked resolution", "blocked", "blocked-review-resolution:\n  review: review-01.md\n  route: review\n  recorded-by: developer\n  evidence:\n    - evidence.md\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := fixture(t, "implementation-complete", tt.firstOutcome, tt.extra)
			write(t, filepath.Join(root, ".concoct/current/review-02.md"), reviewFile(2, "approved"))
			got := Detect(root)
			if got.State != Approved {
				t.Fatalf("state = %s, want %s; diagnostics: %v", got.State, Approved, got.Diagnostics)
			}
		})
	}
}

func TestHistoricalRemediationDoesNotBypassLatestBlockedResolution(t *testing.T) {
	root := fixture(t, "implementation-in-progress", "changes-requested", "remediates-review: review-01.md\nblocked-review-resolution:\n  review: review-02.md\n  route: code\n  recorded-by: developer\n  evidence:\n    - evidence.md\n")
	write(t, filepath.Join(root, ".concoct/current/review-02.md"), reviewFile(2, "blocked"))

	got := Detect(root)
	if got.State != InProgress {
		t.Fatalf("state = %s, want %s; diagnostics: %v", got.State, InProgress, got.Diagnostics)
	}
}

func TestDiscoverableStatusDoesNotMutate(t *testing.T) {
	root := fixture(t, "planned", "", "")
	before := snapshot(t, root)
	_ = Detect(root)
	after := snapshot(t, root)
	if before != after {
		t.Fatal("Detect modified project files")
	}
}

func TestInspectNextActionEvidenceUsesPlanEligibility(t *testing.T) {
	root := fixture(t, "", "", "")
	write(t, filepath.Join(root, ".concoct/roadmap.md"), "---\nversion: 1\nproject: demo\nupdated: 2026-01-01\n---\n# Roadmap\n## APP-001 — Eligible\n- Status: `planned`\n- Priority: `high`\n- Depends on: `none`\n## APP-002 — Blocked\n- Status: `planned`\n- Priority: `critical`\n- Depends on: APP-003\n")
	evidence, err := InspectNextActionEvidence(root)
	if err != nil {
		t.Fatal(err)
	}
	if len(evidence.RoadmapItems) != 2 || !evidence.RoadmapItems[0].Eligible || evidence.RoadmapItems[1].Eligible || !strings.Contains(evidence.RoadmapItems[1].Blocker, "unsatisfied dependency APP-003") {
		t.Fatalf("evidence = %#v", evidence)
	}
	if evidence.RoadmapItems[0].Priority != "high" || len(evidence.SupportedOrigins) != 2 {
		t.Fatalf("evidence = %#v", evidence)
	}
}

func TestInspectPlanEligibilityValidatesCapabilityPrerequisites(t *testing.T) {
	capability := func(status string) string {
		return fmt.Sprintf("---\nversion: 1\nproject: demo\nupdated: 2026-01-01\n---\n# Capabilities\n## CAP-001 — Foundation\n- Status: `%s`\n- Archive: `.concoct/archive/2026-01-01-CAP-001-foundation/`\n### Limitations\n\n- Planner must inspect this limitation.\n", status)
	}
	setup := func(t *testing.T, prerequisite, capabilities string) string {
		root := fixture(t, "", "", "")
		write(t, filepath.Join(root, ".concoct/roadmap.md"), fmt.Sprintf("---\nversion: 1\nproject: demo\nupdated: 2026-01-01\n---\n# Roadmap\n## APP-001 — Demo\n- Status: `planned`\n- Depends on: `none`\n- Capability prerequisites: %s\n", prerequisite))
		write(t, filepath.Join(root, ".concoct/capabilities.md"), capabilities)
		return root
	}

	t.Run("accepted with limitation context", func(t *testing.T) {
		got, err := InspectPlanEligibility(setup(t, "CAP-001", capability("active")), "APP-001")
		if err != nil {
			t.Fatal(err)
		}
		if len(got.Prerequisites) != 1 || !strings.Contains(got.Prerequisites[0].Limitations, "inspect this limitation") || len(got.Prerequisites[0].Archives) != 1 {
			t.Fatalf("eligibility = %#v", got)
		}
	})

	tests := []struct {
		name, prerequisite, capabilities string
		want                             []string
	}{
		{"missing", "CAP-002", capability("active"), []string{"CAP-002 is missing"}},
		{"inactive", "CAP-001", capability("limited"), []string{"status limited"}},
		{"duplicate declaration", "CAP-001, CAP-001", capability("active"), []string{"duplicate capability prerequisite CAP-001"}},
		{"malformed declaration", "cap-one", capability("active"), []string{"malformed Capability prerequisite cap-one"}},
		{"duplicate record", "CAP-001", capability("active") + "\n## CAP-001 — Again\n- Status: `active`\n", []string{"roadmap item APP-001", "duplicate capability CAP-001", "correct .concoct/capabilities.md before retrying planning"}},
		{"missing record status", "CAP-001", strings.Replace(capability("active"), "- Status: `active`\n", "", 1), []string{"roadmap item APP-001", "CAP-001 missing Status", "correct .concoct/capabilities.md before retrying planning"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := InspectPlanEligibility(setup(t, tt.prerequisite, tt.capabilities), "APP-001")
			if err == nil {
				t.Fatal("expected eligibility error")
			}
			for _, want := range tt.want {
				if !strings.Contains(err.Error(), want) {
					t.Fatalf("error = %v, want fragment %q", err, want)
				}
			}
		})
	}
}

func fixture(t *testing.T, status, reviewStatus, extra string) string {
	t.Helper()
	root := t.TempDir()
	write(t, filepath.Join(root, "AGENTS.md"), "# Agents\n")
	write(t, filepath.Join(root, ".concoct/roadmap.md"), roadmap(map[bool]string{true: "active", false: "planned"}[status != ""]))
	write(t, filepath.Join(root, ".concoct/capabilities.md"), "---\nversion: 1\nproject: demo\nupdated: 2026-01-01\n---\n# Capabilities\n")
	if status != "" {
		write(t, filepath.Join(root, ".concoct/current/task-plan.md"), task("APP-001", status, extra))
		notes := "# Notes\n\nImplementation evidence and handoff to reviewer. Blocker disposition.\n"
		if extra != "" && strings.Contains(extra, "remediates") {
			notes += "Finding 1 fixed.\n"
		}
		write(t, filepath.Join(root, ".concoct/current/notes.md"), notes)
	} else {
		write(t, filepath.Join(root, ".concoct/current/task-plan.md"), "# task-plan.md\n")
		write(t, filepath.Join(root, ".concoct/current/notes.md"), "# notes.md\n_Add decisions here._\n_Record meaningful verification results here._\n")
	}
	if reviewStatus != "" {
		write(t, filepath.Join(root, ".concoct/current/review-01.md"), reviewFile(1, reviewStatus))
	}
	write(t, filepath.Join(root, "evidence.md"), "resolved\n")
	return root
}
func roadmap(status string) string {
	return fmt.Sprintf("---\nversion: 1\nproject: demo\nupdated: 2026-01-01\n---\n# Roadmap\n## APP-001 — Demo\n- Status: `%s`\n", status)
}
func task(id, status, extra string) string {
	return fmt.Sprintf("---\nid: %s\ntitle: Demo task\nroadmap-id: %s\nstatus: %s\ncreated: 2026-01-01\nupdated: 2026-01-01\n%scapability-impact:\n  type: add\n  ids: [CAP-001]\n  rationale: Demo impact\n---\n# Task\n", id, id, status, extra)
}
func reviewFile(number int, status string) string {
	return fmt.Sprintf("---\ntask-id: APP-001\nreview: %d\nstatus: %s\ncreated: 2026-01-01\npersona: reviewer\n---\n# Review\n## Outcome\n\n`%s`\n\n## Findings\n\n### Finding 1 — example\n", number, status, status)
}
func write(t *testing.T, path, body string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(body), 0644); err != nil {
		t.Fatal(err)
	}
}
func snapshot(t *testing.T, root string) string {
	t.Helper()
	var b strings.Builder
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			t.Fatal(err)
		}
		if info.Mode().IsRegular() {
			data, _ := os.ReadFile(path)
			fmt.Fprintf(&b, "%s:%x\n", path, data)
		}
		return nil
	})
	return b.String()
}
