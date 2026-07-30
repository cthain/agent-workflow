---
id: CON-007
title: Implement active task planning
roadmap-id: CON-007
status: implementation-complete
created: 2026-07-30
updated: 2026-07-30
remediates-review: review-01.md
git:
  enabled: true
  trunk: main
  task-branch: concoct/con-007-implement-active-task-planning
  base: 302fb230813be6d08aa6b8a8cd942b3ed9245fa7
  status: active
capability-impact:
  type: add
  ids:
    - CAP-008
  rationale: Adds the validated roadmap-to-active-task planning transition on top of the existing prompt-rendering and Git task-isolation capabilities.
---

# Task Plan

## Goal

Complete `concoct plan <roadmap-id>` so one deterministically selected, eligible roadmap item can enter the durable `planned` workflow state through an implementation-ready Task Planner session, without overwriting active work or treating prompt rendering as completed planning.

## Context

CON-006 delivered deterministic Task Planner prompt rendering and the existing CLI already validates the ready state, roadmap ID syntax, item status, and outstanding roadmap dependencies. CON-015 added clean Git-boundary checks, deterministic task-branch creation, rollback on rendering failure, and emitted Git metadata for the planner to preserve.

CON-007 must complete the remaining planning contract: validate declared capability prerequisites against accepted capability truth, preserve safe empty-placeholder behavior, give the planner all required context, validate the resulting plan and notes before activating only the selected roadmap item, and leave durable Git-backed transition evidence that `concoct status` recognizes as `planned`.

## Why this matters

Rendering guidance is useful, but a workflow cannot safely proceed to implementation until roadmap intent, accepted prerequisites, active artifacts, roadmap status, and Git identity agree. This task closes that gap so the first state-changing role transition has deterministic entry checks, recoverable failures, and observable completion evidence.

## Current state

- `concoct plan` discovers the project, requires one roadmap ID, calls `workflow.ValidatePlanItem`, and renders the Task Planner handoff.
- `ValidatePlanItem` checks identifier syntax, roadmap validity, selected status `planned`, and outstanding roadmap dependencies, but it does not parse or validate `Capability prerequisites`.
- `workflow.Detect` already distinguishes empty template placeholders from populated active artifacts and validates populated plan metadata, matching notes, selected roadmap status `active`, capability impact, task status, and Git metadata.
- Git-backed invocation already creates and checks out the deterministic task branch from a clean attached HEAD and rolls it back when rendering or output fails.
- The Task Planner persona owns creation of `task-plan.md` and `notes.md`, selected-item activation after validation, and the complete planning-transition commit; the CLI does not execute planner judgment.
- Existing unit, CLI, prompt golden, Git, and integration tests cover much of the entry and state-detection foundation, but not the complete CON-007 prerequisite and role-completion contract.

## Target state

- Planning eligibility includes deterministic validation of every declared capability prerequisite against an active capability record; missing, duplicate, malformed, inactive, or otherwise unaccepted prerequisite evidence produces an actionable error before branch creation or output.
- Documented capability limitations are surfaced in the Task Planner context so the planner can decide whether they are compatible with the selected outcome; semantic compatibility remains planner judgment rather than a guessed parser rule.
- The generated prompt identifies the selected roadmap item and all canonical planning inputs, including relevant accepted capability and archive evidence, while retaining deterministic output.
- A planner may replace only recognized empty placeholders, writes both schema-valid active artifacts, and changes only the selected roadmap item from `planned` to `active` after the pair validates.
- Git-backed planning preserves the exact recorded trunk, task branch, and base, validates the checked-out branch, and commits the complete planning transition before development. Failure before a valid transition does not leave false workflow advancement or overwrite user work.
- `concoct status` reports `planned` and recommends `concoct code` after successful planning; invalid or partial evidence remains diagnosable.

## Design constraints

- Preserve `AGENTS.md`, the normative command/state-machine documents, and Task Planner role ownership; the CLI renders and validates context but does not author an implementation plan or claim semantic readiness.
- Extend the existing workflow, prompt, CLI, and Git boundaries rather than introducing a second parser, renderer, or planning state authority.
- Treat `.concoct/capabilities.md` as accepted product truth and roadmap `Capability prerequisites` as references to that truth, distinct from unresolved `Depends on` ordering.
- Keep prompt output deterministic and preserve existing create-only `--output` and Git rollback behavior.
- Never overwrite populated current artifacts, reviews, arbitrary files, or unrecognized placeholder content.
- Activate only CON-007 after both active artifacts validate; do not mutate unrelated roadmap items or capabilities during implementation.
- Keep source and embedded template counterparts aligned when shared workflow guidance or assets change.
- Preserve the executable mode of `cmd/concoct/concoct.sh`.

## Non-goals

- Directly launching or embedding a planning agent in the CLI.
- Automatically generating semantically complete task plans or deciding whether a capability limitation is product-compatible.
- Implementing `concoct code`, review-loop mutations, archive automation, or integration changes beyond compatibility required by planning.
- Changing roadmap prioritization, CON-007 product intent, or unrelated capability records.
- Broad schema redesign, generalized Markdown parsing, concurrent/stacked active tasks, worktrees, or provider integrations.
- Reworking the accepted deterministic prompt and Git task lifecycle except where CON-007 validation or completion evidence requires it.

## Working assumptions

- `Status: active` capability entries are accepted truth; other statuses must not satisfy a prerequisite.
- Capability limitation compatibility cannot be inferred reliably from Markdown alone, so deterministic parsing should validate identity/status and the prompt should direct the planner to inspect limitations.
- The existing documented empty template forms are the only replaceable placeholders; absent files may be created, while non-placeholder content is user work and must be refused.
- The role-completion commit is part of successful Git-backed planning, as required by the Task Planner persona and optional Git lifecycle contract.
- CAP-008 is the next capability identifier available at planning time; the Archivist must confirm or reconcile it against repository truth before acceptance.

## Risks and open questions

- Capability records are human-readable Markdown rather than a fully structured schema. Parsing must be narrow, deterministic, and reject ambiguity without imposing undocumented formatting constraints.
- Planning spans CLI entry and agent-authored completion, so tests must clearly separate command-side guarantees from persona/transition evidence and must not imply that rendering alone changes state.
- A failure between writing active artifacts, activating the roadmap item, and committing could leave an invalid intermediate state. Implementation should validate staged content and use ordering/rollback or explicit recovery diagnostics that preserve user-authored planning work.
- Existing prompt golden files will change if capability or archive context is rendered differently; updates must be intentional and reviewed as observable contract changes.
- No unresolved product decision blocks planning: the accepted contract already assigns semantic readiness to the Task Planner and structural/state validation to repository tooling.

## Implementation phases

### Phase 1 — Confirm schemas and remaining contract gaps

Status: `complete`

- Trace CON-007 requirements through `doc/command-reference.md`, `doc/state-machine.md`, the Task Planner persona, workflow detection, prompt selection, CLI invocation, Git operations, templates, and existing tests.
- Define the narrow accepted capability-record and roadmap-prerequisite parsing rules from checked-in artifact conventions, including duplicate/missing/inactive diagnostics.
- Enumerate safe planning boundaries for absent files, recognized empty placeholders, populated task artifacts, current reviews, dirty Git state, and partial planning evidence.

### Phase 2 — Complete eligibility and planner context

Status: `complete`

- Extend the shared workflow model to parse declared capability prerequisites and validate them against accepted capability records without conflating them with roadmap delivery dependencies.
- Return stable, actionable failures before any Git branch or output mutation for malformed, missing, duplicate, or inactive prerequisite evidence.
- Expose the selected prerequisite records, limitations, and deterministic relevant history through the existing prompt context/rendering path so the planner can perform semantic compatibility checks.
- Preserve exact-output determinism and update golden evidence only for deliberate contract additions.

### Phase 3 — Complete safe planning establishment

Status: `complete`

- Implement or refine the planning-completion boundary used by the Task Planner workflow so only absent or recognized empty placeholders may become active artifacts.
- Validate the candidate `task-plan.md` and `notes.md` together, including matching roadmap ID, planned task status, capability impact, observable content, and exact Git metadata when enabled, before selected-item activation is accepted.
- Ensure only the selected item moves from `planned` to `active`, and that partial/interrupted outcomes remain either recoverable without data loss or explicitly `invalid` with actionable diagnostics.
- For Git-backed tasks, verify the recorded branch/base identity and commit the complete planning transition once; retain the non-Git unbranched path.

### Phase 4 — Add contract-level verification

Status: `complete`

- Add focused workflow tests for capability prerequisite parsing and accepted, missing, inactive, duplicate, malformed, and limitation-bearing records.
- Add CLI/Git tests proving failures occur before branch creation, collision/dirty/detached protections remain intact, rendering/output rollback still works, and successful prompt metadata is exact.
- Add end-to-end planning fixtures covering placeholder replacement, populated-artifact refusal, selected-only activation, valid `planned` status, partial-transition diagnostics, Git transition commit evidence, and non-Git behavior.
- Confirm existing roadmap, code, review, archive, integration, status, and prompt modes remain compatible.

### Phase 5 — Reconcile guidance and prepare independent review

Status: `complete`

- Update user and workflow documentation, personas/handoffs, templates, and golden prompts only where delivered planning behavior changes the observable contract.
- Keep source/template counterparts synchronized and search for stale planning rules, paths, persona names, and capability references.
- Record implementation decisions, verification, known risks, and capability impact in notes; inspect the complete diff and leave a fresh reviewer handoff.

## Acceptance criteria

1. `concoct plan <roadmap-id>` selects exactly one parseable `planned` item and refuses missing, malformed, unknown, ineligible, active-conflicting, or dependency-blocked inputs without overwriting workflow artifacts.
2. Every declared capability prerequisite must resolve uniquely to accepted active capability truth before planning begins; failures name the item and prerequisite and recommend an actionable next step.
3. Capability limitations relevant to the selected item are included in deterministic planner context, and the prompt makes semantic compatibility a Task Planner readiness decision.
4. Prompt rendering retains the required persona, selected item, exact canonical inputs, authorized updates, readiness rules, Git metadata when applicable, and next transition; rendering alone does not establish `planned` state.
5. Planning creates absent artifacts or replaces only recognized empty placeholders. Populated or ambiguous current artifacts and current reviews are preserved and produce clear conflict diagnostics.
6. The completed task plan and notes satisfy the canonical artifact schema, preserve the selected roadmap ID and declared capability impact, and include implementation-ready acceptance and verification guidance.
7. Only after both active artifacts validate does the selected roadmap item become `active`; unrelated roadmap items remain byte-for-byte unchanged apart from unavoidable document-level metadata explicitly required by schema.
8. Git-backed planning records `main`, `concoct/con-007-implement-active-task-planning`, and base `302fb230813be6d08aa6b8a8cd942b3ed9245fa7` exactly for this task, validates branch identity, and commits the complete planning transition without duplicate or false advancement. Non-Git planning remains supported.
9. After successful completion, `concoct status` reports `planned` with next action `concoct code`; representative partial or contradictory evidence reports `invalid` with recovery guidance.
10. Existing prompt, workflow-state, Git isolation/integration, and initialization behavior remains passing, and changed shared assets are synchronized with their embedded template counterparts.

## Verification

- Run `gofmt` on changed Go files.
- Run focused tests for `internal/workflow`, `internal/prompt`, `internal/cli`, and `internal/gitrepo` during development.
- Run `go test -count=1 ./...`, `go vet ./...`, and `go build ./cmd/concoct`.
- Run `bash -n cmd/concoct/concoct.sh` and confirm the wrapper remains executable.
- Exercise `concoct plan` in temporary Git repositories for successful planning entry, invalid prerequisites, active conflicts, dirty/detached/collision inputs, output failure rollback, and resulting `planned` status/transition commit evidence.
- Exercise the equivalent planning establishment in a temporary non-Git project.
- Run the repository initialization check required by `AGENTS.md` against a temporary parent; confirm root files, dotfiles, nested templates, personas, planning directories, bootstrap prompt, Git initialization, staged files, and no generated commit.
- Compare every changed source workflow asset with its `templates/` counterpart.
- Run `git diff --check` and search for stale branding, persona names, planning paths, and obsolete command behavior.

## Capability impact

Expected `add`: CAP-008, an observable active-task-planning capability that validates roadmap and accepted capability eligibility, renders complete planner context, safely establishes schema-valid active artifacts, activates only the selected roadmap item, preserves optional Git task identity, and exposes the resulting planned state.

CAP-001, CAP-005, CAP-006, and CAP-007 remain prerequisites and may need wording updates only if the accepted implementation materially changes their documented limitations or relationships; final reconciliation belongs to review and archival.

## Handoff expectations

The Developer should begin by confirming the parser and transition gaps against repository reality, then keep code changes focused on CON-007. Before review, update phase statuses honestly, record durable decisions and test results in `notes.md`, inspect the complete diff, and provide a reviewer handoff listing files changed, checks run, known risks, unresolved work, capability impact, and suggested focus on prerequisite semantics, no-overwrite guarantees, transition atomicity, and Git/non-Git parity.
