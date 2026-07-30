# Notes

## Planning summary

`CON-015` is ready as one cross-cutting Git lifecycle task. Dependencies `CON-003` and `CON-005` are delivered, the roadmap records no remaining product decisions, and no populated active task or review existed. Branch identity, expanded state, archival, integration, recovery, prompts, templates, docs, and tests must agree to avoid false delivery or lost work.

## Confirmed findings

- `internal/cli/cli.go` supports `init`, `status`, `roadmap`, `plan`, `code`, and `review`; only `init` mutates project state.
- `internal/project` invokes Git only for initialization; no reusable Git boundary exists.
- `internal/workflow` recognizes artifact-backed states from `ready` through review outcomes, but no Git/archive/integration metadata or states.
- `internal/prompt` derives exact reads/writes from durable state; Git context must remain validated and deterministic.
- Command/state docs currently specify archive returning directly to `ready`; `CON-015` deliberately changes this contract.
- Tests use synthetic fixtures and real temporary Git repositories. The embedded `templates/` tree is a product contract and must track applicable source changes.
- `CON-007` depends on this item but needs Git planning hooks; `CON-009` owns general archival but this item explicitly needs Git archival/integration. Implement only the cross-cutting behavior required here.

## Repository state at planning

- `.concoct/current/` contained only `.gitkeep`.
- User changes already existed in `.gitignore` and the roadmap metadata date; they were preserved.
- The repository is on `main`, but implementation must not assume that name.

## Planning decisions

- Use provisional `CAP-007`, the next unused stable capability ID; do not change capability truth before accepted archival.
- Keep provider/PR integration, semantic conflict resolution, worktrees, concurrent tasks, generalized downstream-command work, and unrelated cleanup out of scope.
- This task begins under the existing unbranched workflow because the generated prompt authorized only active artifacts and the selected roadmap status; it implements branch isolation for subsequent tasks.

## Risks

- Role work and CLI transition calls occur separately; commit ownership and retry rules must prevent incomplete work from appearing complete.
- Final bookkeeping after squash integration must put delivery evidence on trunk before task-branch deletion or active cleanup.
- Conflict recovery modifies the trunk index/worktree and must distinguish unrelated Git operations.
- Push policy must preserve local/non-interactive use.
- Non-Git fallback needs dedicated fixtures because initialization always creates Git.

## Planning validation — 2026-07-30

- Inspected required inputs, `CON-015`, archive summaries, relevant CLI/project/workflow/prompt code and tests, templates, README, and workflow contracts.
- Confirmed dependencies, eligibility, product readiness, and absence of an active-task conflict.
- No implementation checks were run because the Task Planner changes planning artifacts only.
- `go run ./cmd/concoct status` reported `CON-015`, phase/task status `planned`, capability impact `add`, and next command `concoct code`.
- `go run ./cmd/concoct code` successfully rendered the Developer prompt from the planned state without workflow mutation.
- `git diff --check` passed.

## Task Planner handoff — 2026-07-30

- Current state: `planned`.
- Completed: implementation-ready plan, repository-grounded constraints, transaction risks, verification scenarios, and capability impact.
- Remaining: implement all phases, preserve technical decisions/results here, and produce the Developer-to-Reviewer handoff.
- Highest risks: transactional recovery, final bookkeeping, destructive Git guards, and source/template drift.
- Artifacts: `.concoct/current/task-plan.md`, `.concoct/current/notes.md`, and selected `CON-015` roadmap status after validation.
- Expected role: Developer.
- Recommended command: `concoct code`.

## Development progress — 2026-07-30

### Decisions and implemented foundations

- Added a small `internal/gitrepo` command boundary with root detection, clean
  checks, branch/ref/HEAD inspection, operation guards, squash primitives,
  commit/reset/checkout/delete operations, and deterministic task-branch
  normalization.
- Added optional task-plan `git` metadata and workflow recognition for
  `archived`, `integrating`, and `integrated`. Prompt entry validates clean
  checkout identity for Git-backed active tasks.
- Added an Archivist prompt and conditional Git archival guidance. Git-backed
  archival preserves current and active roadmap evidence; non-Git behavior is
  explicitly retained.
- Added `concoct integrate`, `--continue`, and `--abort` with local recovery
  evidence under `.git/concoct/integrations/`, recorded-trunk squash integration,
  staged human conflict continuation, exact pre-integration reset, final
  delivery/current cleanup, and accepted task-branch deletion.
- The archival commit is an ancestor of the task-branch HEAD rather than the
  HEAD itself: a commit cannot contain its own hash, so a retry-safe follow-up
  metadata commit records the archival hash.

### Verification

- `go test ./...` passed, including a real-repository local integration and
  cleanup test plus deterministic branch normalization coverage.
- `go vet ./...`, `go build ./cmd/concoct`, `bash -n cmd/concoct/concoct.sh`,
  and `git diff --check` passed before the latest documentation refinements.
- Temporary initialization at `/tmp/tmp.vqYP2yYjXF/generated` confirmed root,
  dotfile, nested skill/persona, bootstrap prompt, Git initialization, staged
  files, no commit, executable wrapper, and `ready` status. It remains outside
  the repository because shell cleanup was blocked by the execution policy.

### Work remaining before review

- Add real-repository conflict/continue, abort/retry, collision, drift, dirty,
  and interrupted-bookkeeping tests; harden idempotent continuation after each
  possible integration checkpoint.
- Implement matching-upstream prompt/auto-push policy without making remote
  availability part of local success.
- Complete source/template persona and command/state contract reconciliation,
  golden coverage for the archive prompt, and stale-contract searches.
- Re-run the complete verification suite and inspect the final diff. Task state
  remains `implementation-in-progress`; it is not ready for independent review.

### Developer continuation handoff

- Current state: `implementation-in-progress`.
- Next role: Developer.
- Recommended command: `concoct code`.

## Implementation completion — 2026-07-30

The continuation closed every previously recorded gap:

- Recovery records now persist prepared, squashed, integrated, and delivered
  checkpoints. `--continue` discovers local recovery independently of current
  task files and resumes without duplicating either commit; status refuses to
  report `ready` while recovery remains.
- Real-repository tests cover conflict resolution, unresolved/staged guards,
  exact abort and retry, dirty input, checkout drift, interruption after the
  integration commit, task-branch deletion, trunk checkout, and final cleanup.
- Matching upstreams prompt by default; confirmed and declined paths are
  covered. `.concoct/config.yaml` `git.auto-push: true` is explicit opt-in and
  covered. Missing/non-matching upstreams preserve local success.
- Git-backed `concoct plan` now validates eligibility and repository safety,
  creates the deterministic collision-free task branch, emits exact trunk/base
  metadata, and rolls back branch creation on render/output failure. In-project
  prompt output is rejected so the new branch stays clean.
- Role guidance defines clean, retry-safe planning, implementation, review, and
  archival commit boundaries. Source and embedded template counterparts match.
- Archive prompt golden coverage and Git identity fields in status were added.

## Handoff to reviewer

### Implemented

- Optional Git-backed planning, role validation, archival, squash integration,
  conflict continuation, exact abort, interruption recovery, delivery cleanup,
  and upstream policy, while retaining non-Git role behavior.

### Key decisions

- Committed task-plan Git metadata records trunk, task branch, immutable base,
  archive commit, integration fields, and lifecycle status.
- The archive commit is recorded by a follow-up metadata commit because a Git
  commit cannot contain its own hash; integration validates it as an ancestor
  of the clean task-branch HEAD.
- `.git/concoct/integrations/*.yaml` is local recovery evidence only and remains
  authoritative until validated delivery/cleanup removes it.
- Integration uses two commits on the exact recorded trunk: squash result, then
  delivery bookkeeping. Recovery checkpoints make both retry-safe.
- Accepted task-branch deletion uses explicit forced deletion only after both
  commits and optional push handling; abort hard-resets only the recorded trunk
  to its recorded pre-integration head.

### Files changed

- Git/integration implementation and tests: `internal/gitrepo/`,
  `internal/integration/`.
- CLI, prompt, workflow state and tests: `internal/cli/`, `internal/prompt/`,
  `internal/workflow/`.
- User and normative docs: `README.md`, `doc/command-reference.md`,
  `doc/state-machine.md`, `doc/workflow.md`.
- Source/template workflow contracts: relevant `.codex/skills/`, personas, and
  handoff prompts under both repository and `templates/` trees.

### Verification

- `gofmt -w internal` passed.
- `go test ./...` passed.
- `go vet ./...` passed.
- `go build ./cmd/concoct` passed.
- `bash -n cmd/concoct/concoct.sh` and executable inspection passed.
- `git diff --check` passed.
- Real temporary initialization at `/tmp/tmp.8MuaZvRiHX/generated` confirmed
  root/dot/nested files, personas, prompts, bootstrap guidance, staged files,
  Git repository, no commit, and `ready` status. It remains outside this repo
  because command-policy cleanup was unavailable.
- Source/template diffs for changed skills, personas, and prompts were empty.

### Known risks

- `--abort` is intentionally destructive only to the recorded trunk checkout;
  reviewers should independently verify target resolution and recovery guards.
- Push confirmation reads standard input in the CLI; declined or unavailable
  input safely defaults to no push after successful local delivery.
- Worktrees and concurrent tasks remain explicit non-goals.

### Skipped or unresolved work

- None within `CON-015` scope.

### Capability impact

- Still `add CAP-007`; capability truth must be updated only after approval and
  archival/integration.

### Suggested review focus

- Exercise interrupted integration at each recovery phase and confirm status
  never returns `ready` early.
- Inspect branch deletion and hard-reset target guards, collision rollback, and
  dirty/checkout-drift refusal.
- Verify Git archival cannot mark delivery and that non-Git archival remains
  available.
- Compare every changed source persona/prompt/skill with its template copy.

Current state: `implementation-complete`.
Recommended next command: `concoct review`.

## Review 01 remediation — 2026-07-30

### Finding dispositions

- Finding 1 — `fixed`. Abort now validates the recorded trunk HEAD before any
  reset, refuses later-phase recovery unless HEAD is the exact recorded
  integration or delivery commit, rejects unrelated Git operations, and uses
  phase-appropriate worktree guards. Dirty abort, unrelated committed trunk
  work, and later-phase continuation from an unexpected HEAD have dedicated
  real-repository tests.
- Finding 2 — `fixed`. A `prepared` record is now recoverable while the task
  branch is still checked out. Continue validates the clean task branch,
  archive ancestry, and unchanged trunk before completing checkout and squash;
  abort validates the same pre-mutation state and removes only the recovery
  record. Both paths have real-repository tests.

Staged recovery validation also rejects unstaged, unmerged, or untracked
changes before an integration commit, preventing unrelated files from being
swept into delivery bookkeeping.

### Verification

- `gofmt -w internal/integration/integration.go internal/integration/integration_test.go` passed.
- `go test ./...` passed.
- `go vet ./...` passed.
- `go build ./cmd/concoct` passed.
- `bash -n cmd/concoct/concoct.sh` passed.
- `git diff --check` passed.
- `cmd/concoct/concoct.sh` remains executable.

## Handoff to reviewer — Review 01 remediation

### Implemented

- Exact phase/HEAD guards for continuation and abort.
- Safe prepared-phase recovery from the task branch.
- Worktree and operation guards around destructive reset and resumed commits.
- The four requested real-repository recovery scenarios.

### Key decisions

- `prepared` on the task branch is treated as pre-mutation evidence: abort may
  remove it only after validating clean checkout identity, archive ancestry,
  and an unchanged trunk.
- On the trunk, the permitted HEAD is derived strictly from the recovery phase:
  pre-integration, integration, or delivery commit.
- Conflict/squash abort permits transaction-owned index state but refuses
  untracked or unstaged additions; committed phases require a fully clean tree.

### Files changed

- Remediation code and tests: `internal/integration/integration.go`,
  `internal/integration/integration_test.go`.
- Durable status and handoff: `.concoct/current/task-plan.md`,
  `.concoct/current/notes.md`.

### Verification

- Full Go test, vet, and build checks passed, along with shell syntax,
  executable-mode inspection, formatting, and diff whitespace validation.

### Known risks

- Human-resolved conflict contents remain intentionally human-attested; the
  guard verifies repository shape and prevents unrelated unstaged/untracked
  material, but does not judge semantic resolution.

### Skipped or unresolved work

- None within either Review 01 finding.

### Capability impact

- Remains `add CAP-007`; capability truth is unchanged pending approval and
  archival/integration.

### Suggested review focus

- Re-run the new destructive-operation tests and inspect the phase-to-HEAD
  mapping before each reset or resumed commit.
- Verify both prepared recovery paths preserve the task branch and archive
  evidence when validation refuses.

Current state: `implementation-complete`.
Recommended next command: `concoct review`.

## Review 02 remediation — 2026-07-30

### Finding disposition

- Finding 1 — `fixed`. Pre-commit abort now derives the transaction-owned path
  set from the recorded archive commit's changes since its merge base with the
  pre-integration trunk. Any staged, unstaged, unmerged, or untracked status
  outside that set is refused before `reset --hard`, preserving both user work
  and recovery evidence. Real-repository tests cover a staged new file and a
  staged modification to an unrelated tracked file and verify their index
  contents remain intact after refusal.

### Verification

- `gofmt -w internal/gitrepo/git.go internal/integration/integration.go internal/integration/integration_test.go` passed.
- `go test ./...` passed.
- `go vet ./...` passed.
- `go build ./cmd/concoct` passed.
- `bash -n cmd/concoct/concoct.sh` and executable-mode checks passed.
- A temporary initialization at `/tmp/tmp.CJ609OkZYU/generated` confirmed the
  generated Git repository, staged root/dot/nested templates, personas,
  planning directories, bootstrap prompt, no commit, and `ready` state.
- `git diff --check` passed; the stale-contract search produced only expected
  current, historical, wrapper, and explicit non-Git lifecycle references.

## Handoff to reviewer — Review 02 remediation

### Implemented

- Transaction-path validation before destructive pre-commit abort.
- Regression coverage for staged additions and unrelated tracked modifications.

### Key decisions

- Transaction ownership is bounded to the archive side of the merge-base diff,
  matching the paths the recorded squash integration may legitimately alter.
- Conflict-resolution contents on those paths remain human-attested as planned;
  abort refuses every changed path outside that bounded set.

### Files changed

- `internal/gitrepo/git.go`
- `internal/integration/integration.go`
- `internal/integration/integration_test.go`
- `.concoct/current/task-plan.md`
- `.concoct/current/notes.md`

### Verification

- Full Go test, vet, and build checks passed, along with shell syntax,
  executable-mode, initialization smoke, and diff whitespace checks.

### Known risks

- Conflict-resolution contents on transaction-owned paths remain intentionally
  human-attested; abort now rejects every changed path outside that boundary.

### Skipped or unresolved work

- None within Review 02 Finding 1.

### Capability impact

- Remains `add CAP-007`; capability truth is unchanged pending approval and
  archival/integration.

### Suggested review focus

- Re-run both staged-work tests and confirm abort refuses before reset while the
  recovery record, index entry, and worktree content remain present.
- Inspect the merge-base path boundary against conflict and clean squash cases.

Current state: `implementation-complete`.
Recommended next command: `concoct review`.
