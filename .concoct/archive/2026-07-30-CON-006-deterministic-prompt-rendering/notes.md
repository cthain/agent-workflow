# Notes

## Planning summary

- Readiness result: `ready`.
- CON-006 exists with status `planned`, priority `high`, and a coherent outcome
  covering deterministic prompt rendering for `roadmap`, `plan`, `code`, and
  `review`.
- Dependencies CON-003 and CON-005 are delivered, independently approved,
  archived, and reflected in CAP-001 and CAP-005.
- `.concoct/current/` contained only `.gitkeep` before planning; no active task
  or current review conflicted with this plan.
- Repository code, manual prompts, embedded templates, tests, and normative
  documents support the roadmap assumptions. No unresolved product decision
  requires Product Owner action before implementation.

## Confirmed findings

- `internal/cli/cli.go` currently dispatches `init`, `status`, and help through
  a small explicit switch and uses `project.Discover` for repository commands.
- `internal/workflow/workflow.go` strictly parses state-bearing metadata and
  detects ready, planned, implementation, review, remediation, blocked-review
  recovery, and invalid states without mutation.
- `workflow.Report` contains status-oriented fields but does not expose all
  roadmap, archive, and ownership context required by complete prompt
  rendering; a read-only context extension is likely required.
- The repository and embedded template both contain the manual transition
  prompt corpus. These files already encode canonical role inputs, ownership
  boundaries, completion requirements, and next transitions.
- CON-003's command reference separates prompt rendering from role completion:
  a successfully rendered prompt is state-preserving and is not evidence that
  the selected role performed its work.
- CON-007 owns active task establishment, CON-008 owns code/review state
  transitions and review persistence, CON-010 owns direct agent execution, and
  CON-015 owns Git-backed task isolation. CON-006 must not absorb them.
- Existing workflow tests provide reusable fixtures for state and recovery
  modes, while project tests demonstrate nested discovery, embedded template
  operation, Git inspection, and non-mutation patterns.

## Decisions

- Propose `CAP-006` for the new executable prompt-rendering capability.
- Treat the checked-in manual prompts, personas, and normative command/state
  documents as the durable content contract. Implementation may refactor how
  they are consumed, but must not establish divergent role rules.
- Require deterministic stable ordering for archive summaries, reviews, and
  any other discovered collection, plus byte-identical stdout and file output.
- Require explicit prompt modes for normal development, continued work,
  changes-requested remediation, and both validated blocked-review recovery
  routes where the corresponding command is eligible.
- Keep `plan` state-preserving in CON-006. It validates and renders a Task
  Planner handoff; CON-007 later owns artifact creation and roadmap mutation.
- Keep `review` state-preserving in CON-006. It may name the expected next
  review path, but CON-008 later owns safe allocation and artifact creation.
- Leave the internal package/file layout and safe output-write mechanism to the
  Developer, constrained by explicitness, non-mutation, reproducibility, and
  avoidance of a general template framework.

## Assumptions

- Exact repository-relative file lists can be derived deterministically from
  the selected role, current state, roadmap dependencies, task metadata,
  reviews, and archive metadata without new product policy.
- Conservative inclusion of a relevant archive summary is preferable to an
  opaque heuristic that can omit required historical context; the chosen rule
  must remain deterministic and tested.
- Golden fixtures are appropriate because prompt wording and ordering are part
  of this task's observable output contract.
- The CLI can expose sufficient validated rendering context without making
  workflow metadata structs a broad public API.

## Risks and open questions

- Prompt drift is the primary maintainability risk because prompt content is
  present in repository-owned files and embedded templates. Review should
  confirm one authoritative content path or an enforced synchronization
  strategy.
- Archive relevance is not represented by a dedicated index. The Developer
  must document and test a deterministic evidence-based selection rule.
- Output-path collision and partial-write behavior are not prescribed by the
  roadmap. They remain technical choices, but silent destructive overwrite or
  partial output would violate the plan's safety criteria.
- The current CLI reports detected state but does not expose a command-specific
  eligibility API. Rendering must not duplicate a conflicting transition
  model.
- No product-level question is open. Escalate only if satisfying the task would
  require agent execution, workflow mutation, a new role, a changed ownership
  boundary, or omission of roadmap-required prompt content.

## Relevant history

- `.concoct/archive/2026-07-29-CON-003-command-contract-state-machine/summary.md`
  records approval of deterministic, artifact-backed command and state
  contracts, including remediation and blocked-review recovery.
- `.concoct/archive/2026-07-29-CON-005-go-cli-foundation/summary.md` records the
  accepted Go CLI, installed template distribution, project discovery, strict
  validation, full state detection, and read-only status behavior on which
  CON-006 depends.
- CON-005's structured review-finding identifier idea remains follow-up work and
  is not a CON-006 requirement.

## Handoff to developer

- Current state: `planned` for CON-006 after active-artifact creation and the
  corresponding roadmap status transition.
- Work completed: readiness validation, repository and archive inspection,
  scope boundaries, implementation phases, acceptance criteria, verification,
  and capability-impact proposal.
- Work remaining: implement deterministic prompt context, composition, command
  dispatch, safe output, tests, and focused documentation; then record results
  and prepare an independent reviewer handoff.
- Key constraints: prompt-only and state-preserving behavior, stable bytes and
  ordering, exact role ownership, no review mutation, no agent execution, and
  no absorption of CON-007/008/010/015.
- Suggested eventual review focus: command/state eligibility, remediation and
  blocked-review modes, deterministic archive/review selection, prompt asset
  drift, output-file safety, non-mutation, golden coverage, and compatibility
  with `init`, `status`, installed binaries, and nested discovery.
- Expected next role: Developer.
- Recommended next command: `concoct code`.

## Implementation findings and decisions

- Reused `workflow.Detect` as the state-validation authority and added a narrow
  read-only prompt context instead of introducing a second state model.
- Prompt composition reads the project-owned persona and manual handoff asset
  at render time. Installed templates therefore remain the authoritative prompt
  corpus and no duplicate embedded wording or general template engine was
  introduced.
- Roadmap planning eligibility requires a `planned` item and delivered declared
  dependencies. Dependency parsing accepts the checked-in plain and template
  backtick forms.
- Archive summaries are selected conservatively from roadmap/task identifiers
  and sorted, while prior reviews are taken from the already validated,
  sequential review set. Every emitted collection is sorted before rendering.
- `--output` uses create-only semantics and removes a newly created partial file
  after a write or close failure. Existing files are never silently replaced.
- The untracked `cmd/concoct/concoct` executable predated Developer work and was
  preserved unchanged. Verification used a separate temporary build.

## Verification results

- `gofmt` completed on all changed Go files.
- `go test ./...` passed, including role/mode selection, repeatability,
  dependency rejection, argument validation, nested discovery, stdout/file byte
  equality, non-mutation, and output-collision safety.
- `go vet ./...` passed.
- A Linux CLI build to a temporary path succeeded; help and `code` rendering
  worked from the project root and `internal/prompt`, with byte-identical repeat
  and nested-directory output.
- `bash -n cmd/concoct/concoct.sh` passed and the wrapper remains executable.
- `git diff --check` passed.
- A focused stale-name and agent-invocation search found no newly introduced
  stale persona spelling or direct-agent behavior.

## Handoff to reviewer

### Implemented

- Added deterministic prompt rendering for `roadmap`, `plan`, `code`, and
  `review`, including normal, continuation, remediation, post-remediation, and
  validated blocked-review recovery modes.
- Added ready-state plan-item/dependency validation, sorted archive and review
  context, next-review naming, explicit ownership/outcome/validation sections,
  stdout rendering, and safe `--output` creation.
- Added automated prompt-mode and CLI integration coverage and documented the
  prompt-only command behavior.

### Key decisions

- State authority remains in `internal/workflow`; `internal/prompt` consumes a
  narrow validated context and the repository's canonical prompt assets.
- Output files are create-only rather than overwrite-capable, providing a clear
  safe-failure contract without mutating workflow artifacts.
- Archive relevance uses deterministic identifier matching with conservative
  inclusion; review ordering uses the validated sequence.

### Files changed

- `internal/workflow/workflow.go`
- `internal/prompt/render.go`
- `internal/prompt/render_test.go`
- `internal/cli/cli.go`
- `internal/cli/cli_test.go`
- `README.md`
- `.concoct/current/task-plan.md`
- `.concoct/current/notes.md`

### Verification

- Passed: `gofmt`, `go test ./...`, `go vet ./...`, temporary Linux build,
  root/nested render comparisons, `bash -n cmd/concoct/concoct.sh`, executable
  wrapper check, targeted stale-content search, and `git diff --check`.

### Known risks

- Archive relevance is intentionally conservative because no archive index
  exists; reviewers should confirm the selected summaries are sufficient and
  not excessive for each role.
- Canonical handoff assets are appended verbatim beneath generated context, so
  future asset wording changes intentionally change rendered bytes.

### Skipped or unresolved work

- No workflow mutations, direct agent execution, review persistence, archive
  command, or Git task lifecycle work was added; those remain assigned to later
  roadmap items.
- No generated prompt snapshots were committed. Full-output determinism is
  tested by repeated byte comparison, while role and mode contract sections are
  asserted directly.

### Capability impact

- The plan's proposed `CAP-006` addition remains accurate: executable,
  deterministic, role-aware prompt rendering is now implemented. Capability
  truth remains unchanged pending approval and archival.

### Suggested review focus

- Validate command/state eligibility, blocked and remediation mode selection,
  plan dependency handling, archive relevance/order, exact ownership language,
  output collision/non-mutation behavior, and alignment with the manual prompt
  assets and command contract.

- Expected next role: Reviewer.
- Recommended next command: `concoct review`.

## Review 01 remediation

### Finding 1 — Required golden prompt coverage is absent

- Disposition: fixed.
- Added nine checked-in full-output fixtures under `internal/prompt/testdata/`
  for `roadmap`, `plan`, initial and continuing development, review
  remediation, blocked-review recovery to development, initial review,
  post-remediation review, and blocked-review recovery to review.
- Updated `TestRenderRolesAndModesDeterministically` to compare every rendered
  prompt byte-for-byte with its golden fixture while retaining repeatability,
  role/mode marker, and required-section assertions.
- Retained an explicit `UPDATE_GOLDEN=1` regeneration path so intentional
  prompt-contract changes require a visible fixture diff for review.

## Handoff to reviewer after Review 01

### Implemented

- Resolved the sole Review 01 finding with stable, full-byte golden coverage
  across all four commands and all materially distinct code/review modes.

### Key decisions

- Reused the existing nine-case role/mode matrix as the golden matrix so the
  fixtures protect exactly the state selection already exercised by focused
  assertions.
- Kept golden regeneration opt-in; normal test runs only compare checked-in
  expected output and fail on any drift.

### Files changed in remediation

- `internal/prompt/render_test.go`
- `internal/prompt/testdata/*.golden` (nine fixtures)
- `.concoct/current/task-plan.md`
- `.concoct/current/notes.md`

### Checks run

- `gofmt -w internal/prompt/render_test.go` — passed.
- `go test ./internal/prompt` with `UPDATE_GOLDEN=1` — passed while creating
  the initial reviewed fixtures.
- `go test ./...` without fixture regeneration — passed.
- `go vet ./...` — passed.
- `bash -n cmd/concoct/concoct.sh` — passed.
- `test -x cmd/concoct/concoct.sh` — passed.
- `git diff --check` — passed.

### Known risks

- Golden fixtures intentionally make any canonical prompt-asset wording change
  require reviewed fixture updates; this is the observable contract protection
  requested by the plan and Review 01.

### Skipped or unresolved work

- No Review 01 findings remain unresolved.
- No workflow mutation, direct agent execution, review persistence, archival,
  or later-roadmap work was added.

### Capability impact

- The proposed `CAP-006` addition remains accurate and unchanged. Capability
  truth remains untouched pending independent approval and archival.

### Suggested review focus

- Confirm each golden file represents the intended full output for its named
  role/mode and that a normal test run detects arbitrary byte drift.

- Current state: `implementation-complete` after Review 01 remediation.
- Expected next role: Reviewer.
- Recommended next command: `concoct review`.
