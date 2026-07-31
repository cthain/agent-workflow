# Notes

## Planning summary

- Readiness: `ready` before this planning transition.
- `CON-029` exists with status `planned`, has no roadmap dependencies, and
  defines one coherent documentation outcome.
- The prompt-recorded task branch is checked out at the exact recorded base;
  the worktree was clean and no active task artifacts conflicted with planning.
- The planned change is limited to a human-first root README and has no expected
  runtime or capability-truth impact.

## Confirmed findings

- `README.md` currently introduces installed workflow files and repository
  layout before a usable bootstrap or end-to-end journey.
- The README already contains accurate pieces for building the CLI,
  initialization, status, prompt rendering, Git task isolation/integration,
  source layout, development checks, naming, and repository rename guidance.
  The primary defect is narrative order, onboarding clarity, and actor/effect
  distinction rather than missing low-level documentation.
- `go run ./cmd/concoct help` exposes `init`, `status`, `next`, `roadmap`,
  `plan`, `code`, `review`, `archive`, `integrate`, and `help`.
- `doc/workflow.md`, `doc/command-reference.md`, `doc/state-machine.md`, and
  `doc/multi-agent-workflow.md` provide detailed supporting material suitable
  for links from a concise README.
- The repository currently provides source-build and source-wrapper usage; no
  packaged installation channel should be invented in this task.

## Capability prerequisite compatibility

- `CAP-001`: compatible. The README can explain the durable role contract while
  preserving its limitation that humans or agents supply semantic judgment and
  role outcomes.
- `CAP-005`: compatible. Initialization and read-only status provide the
  supported entry point; metadata validation remains scoped to Concoct schemas.
- `CAP-006`: compatible. The quick start must say role commands render
  deterministic guidance, do not launch agents, do not persist role outcomes,
  and use create-only output paths when that option is discussed.
- `CAP-007`: compatible. A representative Git-backed loop is supported if the
  README preserves exact branch/integration boundaries and human-attested
  semantic conflict resolution; provider workflows and concurrent tasks remain
  outside the story.

## Decisions

- Make the root README's sequence product promise → fit and limitations → quick
  start → conceptual workflow/supporting links → installed/source structure →
  contributor guidance.
- Use the current Git-backed happy path as the representative journey because
  it covers the roadmap item's required integration endpoint and present
  accepted capabilities.
- Treat the review remediation loop and integration conflict recovery as short
  branches with links to detailed documentation.
- Preserve current contributor facts, but place them after user-facing material
  and consolidate duplication where possible.
- Record capability impact as `none`; improved onboarding does not change
  accepted observable behavior.

## Risks

- The `plan` command both renders planner guidance and establishes a safe task
  branch in Git-backed projects; oversimplified prompt-only wording would be
  inaccurate.
- A command transcript without actor annotations could imply autonomous role
  execution. The implementation should pair commands with concise explanations
  of the subsequent human/agent work.
- Existing detailed documentation contains some historical wording, so the
  README should rely on current command and capability truth when descriptions
  differ and should avoid expanding this task into a broad docs reconciliation.

## Relevant history

- CON-003 established the artifact-backed workflow and the separation between
  rendered guidance and completed role work.
- CON-005 delivered initialization and read-only status.
- CON-006 delivered deterministic role-aware prompt rendering.
- CON-015 delivered safe Git task isolation, archival, and local integration.
- CON-028 made `concoct next` the single ready-state recommendation and should
  be reflected early in the user journey.

## Handoff

- Current state: `planned` after the roadmap status and paired active artifacts
  validate.
- Work completed: repository, documentation, command-surface, archive, and
  prerequisite inspection; implementation-ready README plan; acceptance and
  verification criteria.
- Work remaining: rewrite `README.md`, verify all commands, claims, section
  order, and links, update these artifacts with results, and prepare independent
  review.
- Unresolved questions: none requiring Product Owner action; exact prose and
  presentation are Developer decisions within the plan.
- Expected next role: Developer.
- Recommended next command: `concoct code`.

## Implementation results

- Rewrote the root README in the planned order: product promise, fit and
  maturity boundary, representative quick start, detailed references,
  installed contract, and contributor material.
- Made the command/actor boundary explicit at every role transition. The quick
  start separately identifies Git-backed `plan` branch creation and the
  mutating `integrate` transaction.
- Preserved source-build and compatibility-wrapper instructions without
  introducing a package or release channel.
- Kept the implementation documentation-only; no runtime, template, detailed
  documentation, roadmap, review, or capability-ledger files changed.

## Handoff to reviewer

### Implemented

- Added a human-first opening that identifies the intended developer, durable
  workflow value, role boundaries, agent-neutral scope, and current limits.
- Added an end-to-end Git-backed journey from `concoct init hello-world`
  through initial commit, next-action selection, roadmap intake, planning,
  implementation, independent review/remediation, archival, and integration.
- Moved installed structure, source-build instructions, repository layout,
  development checks, naming guidance, and repository-rename guidance after
  the user journey.
- Added focused links to the workflow, command reference, state machine, and
  multi-agent documentation.

### Key decisions

- Used short actor/effect explanations after each command instead of a bare
  transcript, so prompt rendering cannot reasonably be read as autonomous role
  execution.
- Kept conflict recovery concise while naming human semantic responsibility,
  `integrate --continue`, and `integrate --abort`.
- Described create-only `--output` behavior once in the workflow reference
  section rather than repeating it at every prompt command.

### Files changed

- `README.md`
- `.concoct/current/task-plan.md`
- `.concoct/current/notes.md`

### Verification

- `go run ./cmd/concoct help` passed; every displayed Concoct command in the
  README is present in the current command surface.
- `go run ./cmd/concoct status` reported `implementation-in-progress` and the
  expected task branch before this completion update.
- Manually traced the quick start against `doc/command-reference.md`,
  `doc/state-machine.md`, and CAP-001, CAP-005, CAP-006, and CAP-007.
- Checked every README Markdown link with a filesystem existence test; none
  were missing.
- Searched the README for autonomous-agent and semantic-correctness claims;
  only the intentional limitations statement matched.
- `git diff --check` passed before the final artifact update and is included in
  the final verification pass.
- Inspected the complete README diff for narrative order, command effects,
  terminology, and documentation-only scope.

### Known risks

- The quick start assumes an already available `concoct` binary, then points
  source-checkout users to the later build section. This is intentional because
  the repository does not provide a packaged installation channel.
- Detailed documents remain the authority for invalid states and uncommon
  recovery branches; the README deliberately stays at journey level.

### Skipped or unresolved work

- `go test ./...` and `go vet ./...` were not run because only Markdown content
  and workflow artifacts changed, as allowed by the task's proportional
  verification plan.
- No unresolved implementation work or product ambiguity remains.

### Capability impact

- `none`; the README reframes CAP-001, CAP-005, CAP-006, and CAP-007 without
  changing observable product behavior or accepted capability truth.

### Suggested review focus

- Check the actor/effect wording around `plan`, all rendered role prompts,
  archival, and `integrate` against the current command contract.
- Assess whether the opening and quick-start progression give a new user enough
  value and maturity context before introducing repository internals.
- Confirm that the concise push/conflict language does not overstate automated
  semantic validation.

Current state: `implementation-complete`.

Expected next role: Reviewer.

Recommended next command: `concoct review`.

## Archival handoff

- Archive validation passed for active task `CON-029`, roadmap metadata, the
  recorded Git branch and base, accepted implementation presence, required
  artifacts, approving `review-01.md`, and the unused archive destination.
- Preserved accepted task, notes, and review evidence with a durable summary at
  `.concoct/archive/2026-07-31-CON-029-human-first-readme/`.
- Capability impact remains `none`; the README improves presentation of
  CAP-001, CAP-005, CAP-006, and CAP-007 without changing accepted behavior, so
  `.concoct/capabilities.md` remains unchanged.
- Delivery remains pending integration; current task evidence stays intact and
  CON-029 remains active until `concoct integrate` succeeds.
- Expected next transition: `concoct integrate`.
