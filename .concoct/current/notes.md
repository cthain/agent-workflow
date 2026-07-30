# Notes

## Planning summary

CON-007 is ready for implementation. Its product outcome remains coherent, its declared dependencies are `None`, and CAP-001, CAP-005, CAP-006, and CAP-007 are active accepted prerequisites. The work completes the roadmap-to-task transition by building on previously delivered prompt rendering and Git isolation rather than recreating them.

## Confirmed findings

- The planning command already requires a stable roadmap ID, ready workflow state, selected status `planned`, and satisfied outstanding roadmap dependencies.
- `internal/cli` already creates a deterministic Git task branch before rendering, emits trunk/branch/base, refuses unsafe repository states through `internal/gitrepo`, and rolls back branch creation after rendering or output failure.
- `internal/prompt` already selects the Task Planner persona and canonical handoff, lists authorized planning artifacts, and states that rendered output does not complete role work.
- `internal/workflow` already recognizes template placeholders, validates paired populated plan/notes artifacts, requires the selected roadmap item to be `active`, validates capability impact and Git metadata, and maps valid task status `planned` to the next command `concoct code`.
- `ValidatePlanItem` parses roadmap status and `Depends on` but does not parse `Capability prerequisites` or confirm those references against `.concoct/capabilities.md`.
- The command reference assigns semantic plan creation and limitation-compatibility judgment to the Task Planner; automatic plan authorship or direct agent execution is not part of accepted behavior.
- The current branch and HEAD exactly match the rendered prompt: `concoct/con-007-implement-active-task-planning` at `302fb230813be6d08aa6b8a8cd942b3ed9245fa7`, with trunk `main`.

## Decisions

- Treat CON-007 as completion of an existing partial command surface. Preserve CON-006 deterministic rendering and CON-015 Git lifecycle behavior.
- Validate prerequisite identity and accepted status structurally, but surface documented limitations to the planner for semantic compatibility judgment.
- Use the existing workflow detector/parser boundary as the state authority; do not create an independent planning-state model.
- Declare a new CAP-008 because successful planning is a new observable transition rather than an internal-only refactor.
- Include Git-backed transition-commit evidence in scope because the current Task Planner persona explicitly requires the complete planning transition to be committed before development.

## Assumptions and risks

- Active capability status represents accepted current truth. The Developer must confirm the exact narrow Markdown parsing convention and reject ambiguity consistently.
- The repository does not yet expose a standalone CLI subcommand that marks planner role completion. Implementation should respect the agent-driven handoff model and must not imply that initial prompt rendering alone changes workflow artifacts.
- Multi-file planning establishment can be interrupted. Ordering and diagnostics must preserve user-authored work and prevent false `planned` state rather than silently deleting partial artifacts.
- Golden prompt output is an observable contract; any update needs explicit test evidence.

## Relevant history

- CON-003 established the normative planning command and artifact-backed state machine.
- CON-005 delivered CLI discovery, initialization, status detection, and artifact validation foundations.
- CON-006 delivered deterministic prompt rendering and planner persona selection.
- CON-015 delivered deterministic Git task branches, exact planning metadata, safe boundary checks, archival/integration lifecycle, and non-Git fallback.

## Task Planner handoff

- Current state: `planned` once these artifacts and selected roadmap activation validate.
- Work completed: repository inspection, prerequisite/readiness confirmation, scope reconciliation with prior deliveries, implementation phases, acceptance criteria, verification plan, capability impact, and exact Git metadata.
- Work remaining: implement CON-007 according to `task-plan.md`; no source code was changed during planning.
- Known review risks: capability Markdown parsing, semantic-vs-structural prerequisite responsibility, safe placeholder replacement, interrupted transition behavior, exact prompt output, and Git/non-Git parity.
- Checks run during planning: verified branch `concoct/con-007-implement-active-task-planning`, HEAD/base `302fb230813be6d08aa6b8a8cd942b3ed9245fa7`, clean starting worktree, roadmap eligibility, active capability prerequisites, relevant archived deliveries, source/test boundaries, and normative command/state contracts.
- Expected next role: Developer.
- Recommended next command: `concoct code`.

## Implementation findings and decisions

- Extended the existing roadmap parser rather than adding a second planning
  model. `Capability prerequisites` remains distinct from `Depends on` and is
  validated only for the selected planned item.
- Added a narrow capability-record parser that rejects duplicate records and
  missing statuses, requires each selected prerequisite to resolve uniquely to
  `Status: active`, and preserves limitations plus archive-summary provenance
  for rendering.
- Kept semantic limitation compatibility with the Task Planner. The prompt
  reports structurally accepted prerequisites, tells the planner whether
  limitations are documented, names available provenance summaries, and
  explicitly assigns the compatibility decision to that role.
- Preserved the accepted agent-driven completion boundary. The command still
  does not author or activate task artifacts; existing workflow detection,
  placeholder recognition, selected-item agreement, and Git identity checks
  remain the authority for a completed planning transition.
- Prerequisite validation occurs before `CreateTaskBranch`, so malformed,
  duplicate, missing, and inactive evidence cannot create a task branch or
  emit a false planning transition.

## Verification results

- `gofmt` on changed Go files — passed.
- `go test -count=1 ./...` — passed across CLI, Git, integration, project,
  prompt, and workflow packages.
- `go vet ./...` — passed.
- `go build ./cmd/concoct` — passed; the generated local binary was moved to
  trash after verification.
- `bash -n cmd/concoct/concoct.sh` and executable-mode check — passed.
- Source/template comparisons for the changed Task Planner persona and handoff
  prompt — passed.
- Temporary initialization produced a ready Git repository with staged root,
  dotfile, nested template, persona, prompt, current planning, bootstrap, and
  executable wrapper assets and no commit. The expected `git log` check exited
  nonzero because an unborn `main` has no commits, confirming the no-commit
  condition.
- `git diff --check` — passed before the final artifact update.

## Handoff to reviewer

### Implemented

- Added deterministic capability-prerequisite parsing and validation to
  planning eligibility.
- Added accepted prerequisite, limitation, and archive provenance context to
  the Task Planner prompt.
- Added workflow, prompt, CLI/Git boundary, and golden-output coverage.
- Reconciled command/state documentation and synchronized planner guidance with
  embedded templates.

### Key decisions

- Only `active` capability records satisfy prerequisites.
- Capability identity/status is structural; limitation compatibility remains
  Task Planner judgment.
- Prompt rendering remains non-mutating apart from the already accepted Git
  task-branch creation boundary.

### Files changed

- `internal/workflow/workflow.go` and tests.
- `internal/prompt/render.go`, tests, and plan golden output.
- `internal/cli/cli_test.go`.
- `doc/command-reference.md` and `doc/state-machine.md`.
- Source and template Task Planner persona and Product Owner-to-Task Planner
  handoff.
- Developer-owned task plan and notes.

### Verification

- Full Go tests, vet, build, shell syntax/mode, template parity,
  initialization, and diff whitespace checks passed as recorded above.

### Known risks

- Capability parsing intentionally follows the checked-in heading/status and
  section conventions; future schema changes must update this parser.
- Archive provenance is included when a capability record names a conventional
  `.concoct/archive/<task>/` path. Semantic archive relevance remains
  deliberately conservative.

### Skipped or unresolved work

- No standalone role-completion mutation was added because agent-authored plan,
  notes, selected-item activation, and transition commit remain the accepted
  workflow contract and direct agent execution is out of scope.
- No product ambiguity or unresolved implementation blocker remains.

### Capability impact

- Expected `add`: CAP-008, completing the observable validated
  roadmap-to-active-task planning transition. Final capability reconciliation
  remains Archivist-owned.

### Suggested review focus

- Confirm malformed/duplicate/missing/inactive prerequisite diagnostics, the
  pre-branch failure boundary, deterministic limitation/provenance rendering,
  and that the implementation does not claim rendered guidance is completed
  planning.

- Current state: `implementation-complete`.
- Expected next role: Reviewer.
- Recommended next command: `concoct review`.

## Review 01 remediation

### Finding 1 — Malformed capability-record errors lack required planning context

- Disposition: `fixed`.
- `InspectPlanEligibility` now wraps capability-ledger parse diagnostics with
  the selected roadmap item and an actionable instruction to correct
  `.concoct/capabilities.md` before retrying planning.
- The original diagnostic remains intact, so duplicate capability records and
  missing capability statuses continue to name the affected prerequisite.
- Focused tests now require the roadmap item, affected capability diagnostic,
  and recovery instruction for both cases.
- The validation location is unchanged: malformed capability evidence is still
  rejected before CLI task-branch creation or output.

## Remediation verification results

- `gofmt` on `internal/workflow/workflow.go` and
  `internal/workflow/workflow_test.go` — passed.
- `go test -count=1 ./internal/workflow` — passed.
- `go test -count=1 ./...` — passed across all packages.
- `go vet ./...` — passed.
- `go build -o /tmp/concoct-con-007-remediation ./cmd/concoct` — passed; the
  temporary build output was removed after verification.
- `bash -n cmd/concoct/concoct.sh` and executable-mode check — passed.
- `git diff --check` — passed before this artifact update.

## Handoff to reviewer after remediation

### Implemented

- Enriched malformed capability-ledger planning errors with the selected
  roadmap item and a concrete retry prerequisite.
- Strengthened duplicate-record and missing-status tests to cover the complete
  acceptance-level diagnostic contract.

### Key decisions

- Wrapped diagnostics at the planning-eligibility boundary instead of changing
  the reusable capability parser's file-specific diagnostics.
- Preserved validation ordering and all existing safe pre-branch behavior.

### Files changed

- `internal/workflow/workflow.go`
- `internal/workflow/workflow_test.go`
- `.concoct/current/task-plan.md`
- `.concoct/current/notes.md`

### Verification

- Focused and full Go tests, vet, build, shell syntax/mode, and diff whitespace
  checks passed as recorded above.

### Known risks

- No new risk identified; this remediation changes error context only.

### Skipped or unresolved work

- None. Review 01 contained one finding, and it is fixed.

### Capability impact

- Unchanged expected `add`: CAP-008. Final reconciliation remains
  Archivist-owned after approval.

### Suggested review focus

- Confirm both malformed-ledger errors contain the selected roadmap item, the
  affected capability diagnostic, and an actionable correction, while still
  failing before branch creation.

- Current state: `implementation-complete` remediating `review-01.md`.
- Expected next role: Reviewer.
- Recommended next command: `concoct review`.

## Archival handoff

- Current state: `review-approved`; archival evidence is being committed on the recorded Git task branch.
- Accepted outcome: CON-007 adds validated capability-prerequisite eligibility and deterministic capability limitation/provenance context to the existing roadmap-to-active-task planning workflow.
- Approving review: `review-02.md` (`approved`).
- Capability reconciliation: add CAP-008; CAP-001, CAP-005, CAP-006, and CAP-007 remain prerequisite capabilities rather than being replaced.
- Verification: `go test -count=1 ./...`, `go vet ./...`, `bash -n cmd/concoct/concoct.sh`, executable-mode validation, and `git diff --check 302fb230813be6d08aa6b8a8cd942b3ed9245fa7..HEAD` passed immediately before archival.
- Archive destination: `.concoct/archive/2026-07-30-CON-007-active-task-planning/`.
- Delivery remains pending integration; current task evidence must remain intact and CON-007 must remain active until `concoct integrate` succeeds.
- Expected next transition: `concoct integrate`.
