---
task-id: CON-007
review: 1
status: changes-requested
created: 2026-07-30
persona: reviewer
---

# Review 01

## Outcome

`changes-requested`

## Summary

The implementation is focused, preserves the accepted prompt-rendering and Git
boundaries, and independently passes the full Go test, vet, build, shell,
template-parity, and diff-whitespace checks. Capability prerequisites are
validated before task-branch creation and accepted limitation/provenance context
is rendered deterministically.

One observable error-path requirement remains incomplete. Two malformed
capability-ledger cases reject planning without the selected roadmap item or an
actionable recovery recommendation, contrary to acceptance criterion 2.

## Acceptance criteria assessment

- Criteria 1 and 3–10 are satisfied by the implementation, existing workflow
  behavior, focused tests, and repository evidence.
- Criterion 2 is partially satisfied: missing and inactive prerequisite errors
  name the roadmap item and prerequisite and provide a next action, but duplicate
  capability records and missing capability statuses do not.

## Findings

### Finding 1 — Malformed capability-record errors lack required planning context

- Severity: `minor`
- Status: `unresolved`
- Evidence: `InspectPlanEligibility` returns `invalid capabilities: ... duplicate
  capability CAP-001` or `invalid capabilities: ... CAP-001 missing Status`
  immediately after `parseCapabilities`. These messages do not identify the
  selected roadmap item and do not recommend a corrective next step. The focused
  tests in `internal/workflow/workflow_test.go` assert only the shorter fragments,
  so the acceptance-level diagnostic contract is not covered.
- Impact: A failed `concoct plan <roadmap-id>` does not consistently tell the
  operator which planning attempt was rejected or how to recover. This violates
  acceptance criterion 2 for two explicitly required malformed/duplicate cases,
  although the safe pre-branch failure boundary itself is correct.
- Required action: Make duplicate-record and missing-status prerequisite failures
  name the selected roadmap item and affected prerequisite and recommend an
  actionable correction, then strengthen tests to assert that context. Preserve
  the current behavior that rejects the input before branch creation or output.

## Verification performed

- `go test -count=1 ./...` — passed.
- `go vet ./...` — passed.
- `go build -o /tmp/concoct-review-build ./cmd/concoct` — passed.
- `bash -n cmd/concoct/concoct.sh` — passed.
- `git diff --check 302fb230813be6d08aa6b8a8cd942b3ed9245fa7..HEAD` — passed.
- Source/template comparison for the changed Task Planner persona and handoff —
  passed.
- `go run ./cmd/concoct status` — reported `implementation-complete` and
  recommended `concoct review` before this artifact was created.
- Inspected the complete committed diff from the recorded task base through
  implementation commit `5016c34`, including relevant workflow, prompt, CLI
  tests, guidance, golden output, task artifacts, and archive summaries.

## Capability impact assessment

The declared `add` impact for CAP-008 is accurate if the diagnostic gap is
remediated and subsequently approved. CAP-001, CAP-005, CAP-006, and CAP-007
remain prerequisites rather than capabilities replaced by this task.

## Scope, compatibility, and documentation assessment

The implementation remains within CON-007 scope. It extends the existing
workflow parser and prompt renderer, leaves role-completion authorship with the
Task Planner, and does not broaden Git mutation behavior. Documentation and
embedded template counterparts are synchronized. No unrelated source changes or
compatibility regressions were identified.

## Prior finding disposition

No prior reviews exist for this active task.

## Handoff

- Current state: review changes requested.
- Work remaining: resolve Finding 1 and add diagnostic assertions.
- Known risks: keep malformed prerequisite rejection ahead of task-branch and
  output mutation while enriching errors.
- Expected next role: Developer.
- Recommended next command: `concoct code`.
