---
id: CON-028
title: Recommend the next project action
roadmap-id: CON-028
status: implementation-complete
created: 2026-07-30
updated: 2026-07-30
remediates-review: review-01.md
capability-impact:
  type: add
  ids:
    - CAP-009
  rationale: Adds a read-only Product Owner recommendation command and makes it the single ready-state entry point.
git:
  enabled: true
  trunk: main
  task-branch: concoct/con-028-recommend-the-next-project-action
  base: 23dc20a7393ff83f7e4967c5da184cc5319bc849
  status: active
---

# Task Plan

## Goal

Implement `concoct next` as the single command recommended from valid `ready`
state. The command must assemble authoritative project evidence and render a
deterministic Product Owner prompt that recommends one valid next action without
selecting work, creating a task, or mutating lifecycle evidence.

## Context

Concoct currently reports `concoct roadmap or concoct plan <roadmap-id>` in
ready state, and initialization, integration, documentation, templates, skills,
and personas repeat variants of that unexplained choice. Existing workflow
inspection and prompt rendering already provide the structural foundation, but
there is no explicit decision step between reaching `ready` and choosing the
next supported workflow action.

## Current state

- `internal/workflow.Detect` validates ready state and reports two possible next
  commands without explaining which is appropriate.
- `internal/prompt` renders role-aware prompts for `roadmap`, `plan`, `code`,
  `review`, and `archive`; it has deterministic archive discovery and full
  golden-output coverage.
- `internal/cli` provides shared create-only `--output` handling for prompt
  commands, but has no `next` command.
- Initialization and successful integration return to ready state and publish
  legacy roadmap-or-plan guidance.
- The only currently supported work origins are roadmap planning and human
  product input through roadmap reconciliation. Candidate future origin types,
  including bugs, must not be presented as executable paths.

## Target state

- Every structurally valid ready repository has exactly one recommended command:
  `concoct next`.
- `concoct next` renders a complete, byte-deterministic Product Owner prompt
  from validated roadmap, capability, dependency, prerequisite, archive, and
  supported-origin evidence.
- The prompt requires one evidence-backed recommendation outcome: plan one
  structurally eligible roadmap item, perform supported roadmap/product intake
  or reconciliation, resolve a named blocker or inconsistency, or acknowledge
  that no actionable work is recorded.
- Rendering remains advisory and read-only; the Product Owner performs semantic
  prioritization and the existing `roadmap` and `plan` commands retain their
  distinct mutation and selection boundaries.
- Source and installed templates consistently teach the new ready-state entry
  point.

## Design constraints

- Reuse `internal/workflow` as the authority for structural state and planning
  eligibility; do not create a competing parser or infer acceptance from code,
  priority, or implementation existence.
- Separate deterministic evidence assembly from Product Owner judgment. The
  rendered prompt may expose ordered evidence, but the CLI must not claim it has
  selected the next work item.
- Permit `next` only in valid `ready` state. Invalid or contradictory evidence
  must retain actionable diagnostics rather than being normalized into a prompt.
- Recommendations must never describe an item with unresolved dependencies or
  missing/inactive capability prerequisites as immediately plannable.
- Use only accepted, currently executable work-origin contracts. Structure the
  evidence boundary so future accepted origins can be added without changing
  the command's advisory role.
- Preserve deterministic ordering, complete prompt visibility, stdout behavior,
  create-only output safety, project discovery, and non-mutation guarantees.
- Keep source workflow assets and their embedded `templates/` counterparts in
  sync, and preserve the agent-neutral contract.
- Preserve existing Git task isolation and integration behavior.

## Non-goals

- Do not auto-select, activate, or create a task.
- Do not edit roadmap, capabilities, current-task, bug, or archive artifacts as
  a side effect of `next`.
- Do not implement candidate task-origin, bug-lifecycle, workflow-policy, direct
  agent execution, archive, code, or review roadmap work.
- Do not replace `status`, `roadmap`, or `plan`, or merge their responsibilities
  into `next`.
- Do not introduce remote services, configuration-driven ranking, or a general
  workflow graph.
- Do not broadly rewrite unrelated documentation or persona content.

## Working assumptions

- Product Owner judgment can be expressed as prompt guidance over deterministic
  repository evidence without a new durable recommendation artifact.
- A roadmap item is structurally plannable only when the existing planning
  eligibility rules accept its status, dependencies, and capability
  prerequisites; priority remains advisory Product Owner evidence.
- Human product input and roadmap maintenance are the only supported alternative
  origins at this capability boundary. Other roadmap-described origins remain
  unavailable until their own contracts are accepted.
- `CAP-009` is the next available capability identifier and will describe the
  observable next-action recommendation behavior if the implementation is
  approved and archived.

## Risks and open questions

- Evidence assembly could duplicate workflow parsing and drift from `plan` or
  `status`; keep shared eligibility logic authoritative and test agreement.
- A prompt that merely dumps the roadmap would leave Product Owner choice
  underspecified; define explicit outcome and evidence requirements in the
  source and template Product Owner guidance.
- Priority ordering can look like automated selection. Output and tests must
  distinguish deterministic presentation from semantic recommendation.
- Updating all ready-state guidance spans executable output, normative docs,
  reusable skills, personas, and embedded templates; parity and stale-string
  checks are required.
- No product decision remains unresolved. Local representation of the evidence
  model and prompt sections is a Developer decision within these constraints.

## Implementation phases

### Phase 1 — Define the next-action evidence contract

Status: `complete`

- Extend the shared workflow inspection boundary with a deterministic,
  read-only view of roadmap items, priorities, dependencies, accepted
  prerequisite status and limitations, relevant archive provenance, and the
  currently supported work origins.
- Represent structural eligibility and blockers explicitly enough that prompt
  rendering cannot present blocked work as plannable.
- Keep malformed or contradictory canonical evidence on the existing invalid
  path with useful recovery context.

### Phase 2 — Render the Product Owner recommendation prompt

Status: `complete`

- Add `next` to the prompt role selection and use the Product Owner persona plus
  a canonical source/template handoff asset suited to recommendation rather
  than roadmap mutation.
- Render all required authoritative evidence in stable order and require exactly
  one supported recommendation outcome, its rationale, cited durable evidence
  and blockers, and the exact follow-up command when available.
- Preserve the existing deterministic stdout/create-only output contract and
  make `next` explicitly authorize no workflow-artifact writes.

### Phase 3 — Wire the CLI and ready-state transitions

Status: `complete`

- Add `concoct next [--output <path>]` to CLI dispatch, usage, argument
  validation, nested discovery, and shared prompt-output behavior.
- Change ready-state status, initialization/bootstrap output, successful
  integration output, and other executable ready transitions to recommend
  exactly `concoct next`.
- Preserve existing command eligibility and Git lifecycle behavior outside the
  changed recommendation boundary.

### Phase 4 — Reconcile the portable contract

Status: `complete`

- Update `README.md`, command reference, state-machine/workflow documentation,
  the Concoct skill, applicable personas and handoffs, and their installed
  template counterparts to distinguish `status`, `next`, `roadmap`, and `plan`.
- Remove stale ready-state instructions that offer an unexplained choice while
  retaining `roadmap` and `plan` as explicit follow-up actions selected by the
  Product Owner.
- Keep repository-owned and embedded template assets semantically aligned.

### Phase 5 — Verify and prepare review

Status: `complete`

- Add focused workflow tests for eligibility/blocker evidence and invalid input.
- Add full-output golden prompt cases covering an eligible roadmap item,
  supported non-planning Product Owner work, roadmap reconciliation, a specific
  blocker, and no actionable recorded work.
- Add CLI tests for ready-state restriction, nested discovery, argument errors,
  byte-identical stdout/file output, create-only collision refusal, and complete
  workflow non-mutation.
- Exercise initialization and Git integration ready-state outputs, source/template
  parity, documentation guidance, and the repository's standard Go, shell, and
  diff checks.
- Update the active artifacts with honest implementation status, results,
  durable decisions, and a reviewer handoff; do not modify capability truth.

## Acceptance criteria

- `concoct status` in valid ready state, successful initialization, and
  successful integration recommend exactly `concoct next` as the next workflow
  command.
- `concoct next` succeeds only in structurally valid ready state and rejects
  other lifecycle states or invalid canonical evidence without mutation.
- Unchanged repository evidence produces byte-identical prompt bytes on
  repeated runs and through stdout versus a newly created `--output` file;
  existing output destinations are not overwritten.
- The prompt contains authoritative roadmap, accepted capability, dependency,
  prerequisite, relevant archive, and supported-origin evidence in stable order
  and clearly identifies structural eligibility and blockers.
- Product Owner guidance requires exactly one recommendation covering each
  supported outcome class: plan an eligible roadmap item, address supported
  product input or roadmap reconciliation, resolve a specific blocker or
  inconsistency, or report no actionable recorded work.
- Every recommendation explains why it is next, cites durable evidence and
  blockers, and supplies the exact existing follow-up command when one applies.
- No recommendation treats priority, implementation presence, passing checks,
  unresolved dependencies, or unsupported/inactive capability prerequisites as
  acceptance or planning eligibility.
- Invoking `next` changes no roadmap, capability, current-task, bug, archive, or
  other workflow artifact and does not create a Git branch.
- Help, documentation, source and template skills/personas/handoffs, bootstrap
  guidance, and ready-state transitions consistently explain the boundaries of
  `status`, `next`, `roadmap`, and `plan`.
- Tests demonstrate all recommendation outcome classes, state validation,
  evidence ordering, non-mutation, output safety, initialization/integration
  guidance, and source/template parity.

## Verification

- `go test -count=1 ./...`
- `go vet ./...`
- `go build ./cmd/concoct`
- `bash -n cmd/concoct/concoct`
- Confirm `cmd/concoct/concoct` remains executable.
- Initialize a project under a temporary parent, confirming root files,
  dotfiles, nested templates, personas, planning directories, Git repository,
  bootstrap prompt, staged files, no generated commit, and `concoct next`
  ready-state guidance.
- Run `concoct next` and `concoct status` from the generated project and a nested
  directory, verifying deterministic output and no workflow mutation.
- Exercise or extend integration tests to confirm successful delivery returns
  ready state with `concoct next` as the sole recommendation.
- Compare task-relevant source workflow assets with their installed template
  counterparts.
- `git diff --check`
- Search source, templates, skills, personas, help, and documentation for stale
  ready-state roadmap-or-plan recommendations and obsolete command counts.

## Capability impact

Expected impact is `add`: on acceptance, add `CAP-009` for deterministic,
read-only, Product Owner next-action recommendation. Existing CAP-001, CAP-005,
CAP-006, and CAP-008 remain prerequisites and should be updated only if the
delivered behavior materially changes their accepted descriptions.

## Handoff expectations

The Developer should begin by setting task status to
`implementation-in-progress`, validate the evidence model against current
workflow parsers and tests, and keep changes limited to `CON-028`. Before review,
record files changed, decisions, checks, source/template parity, known risks,
skipped work, and final capability impact in `notes.md`; set status to
`implementation-complete`; and recommend `concoct review`.
