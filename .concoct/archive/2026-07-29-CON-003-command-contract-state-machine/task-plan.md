---
id: CON-003
title: Define the command contract and workflow state machine
roadmap-id: CON-003
status: implementation-complete
created: 2026-07-29
updated: 2026-07-29
capability-impact:
  type: update
  ids:
    - CAP-001
  rationale: Formalizes the existing file-based workflow as a complete command and state-transition contract without implementing CLI behavior.
---

# Task Plan

## Goal

Define the normative contract for Concoct's initial command surface and workflow state machine in the repository-root `doc/` product-documentation directory so later CLI and prompt-rendering work can be implemented without inventing command behavior, transition rules, artifact ownership, or recovery guidance.

## Context

CON-003 is the contract-setting task between the existing manual, file-based workflow and the planned executable CLI. The roadmap names seven initial commands and requires each command to define its inputs, valid starting states, artifact interactions, selected persona, prompt behavior, resulting state, failures, and next commands.

The deliverables are:

- `doc/command-reference.md`
- `doc/state-machine.md`

These documents are intended to guide CON-005 through CON-009. They must describe intended behavior clearly while distinguishing it from behavior already implemented by the current shell initializer.

## Why this matters

Current workflow behavior is distributed across the roadmap, skill, personas, prompts, capabilities, and legacy documentation. Those sources agree on the overall role sequence but use different levels of state detail and contain stale paths and persona names. A single explicit contract is needed before CLI state detection and deterministic prompt rendering can be implemented consistently.

## Current state

- `CAP-001` provides a durable file-based workflow contract. The new canonical command-reference and state-machine content has been authored, but both documents currently sit under `.concoct/docs/`, which conflicts with the repository's source-documentation layout.
- `CAP-002` provides manual role-transition prompts; commands do not yet render them.
- The roadmap defines the initial commands and a high-level state flow, including the review-remediation loop.
- `.codex/skills/concoct/SKILL.md` and the current personas define the canonical artifacts, role ownership, review outcomes, archival rules, and a more detailed durable workflow progression.
- `cmd/concoct/concoct` is a legacy bootstrap shell script, not the planned command-oriented CLI. From its checked-in location it resolves template and persona paths incorrectly and cannot complete initialization.
- Repository guidance and legacy documentation contain stale root-level paths and older persona names. These are evidence of drift, not alternative product direction.
- Existing reusable product documentation lives under `doc/`; `.concoct/` is the living project-local workflow area. The two CON-003 documents must be relocated before the task is reviewable.
- There are no automated tests in the repository.

## Target state

- A command reference gives every initial command a complete, internally consistent contract.
- A state-machine document defines states in terms of observable repository artifacts, valid transitions, review outcomes, remediation loops, invalid and ambiguous states, and recovery guidance.
- Both documents live under the repository-root `doc/` directory alongside Concoct's other reusable product documentation; no source documentation remains under `.concoct/docs/`.
- The documents agree with each other and explicitly map the roadmap's high-level state vocabulary to the more detailed artifact-backed workflow phases where necessary.
- Later implementation tasks can use the contract without deciding product behavior themselves.
- Documentation does not imply that the planned CLI commands are already implemented.

## Design constraints

- Preserve the initial command surface from CON-003:
  - `concoct init <project>`
  - `concoct status`
  - `concoct roadmap`
  - `concoct plan <roadmap-id>`
  - `concoct code`
  - `concoct review`
  - `concoct archive`
- Treat `handoff`, `abandon`, and `doctor` as optional later commands, not requirements or hidden dependencies of the happy path.
- Keep the contract agent-neutral and compatible with the canonical role boundaries.
- Define workflow state from durable repository evidence rather than conversation history or agent claims.
- Treat `AGENTS.md` as canonical project guidance, `capabilities.md` as accepted current truth, and `roadmap.md` as intended future behavior.
- Preserve transactional archive ordering: active artifacts are cleared only after archive writes and required reconciliations succeed.
- Preserve review history and support repeated `changes-requested → code → review` cycles.
- Each role command must cover its incoming context and outgoing handoff; normal operation must not require a separate `handoff` command.
- Use lowercase hyphenated Markdown filenames and repository-relative paths.
- Treat `doc/` as the canonical location for reusable Concoct product documentation. Reserve `.concoct/` for living project-local workflow state, personas, prompts, roadmap, capabilities, reviews, and archives.
- Preserve the existing normative content while relocating the two documents. Change content only where a path, link, or reference must be corrected for the new location, or where inspection identifies a separate defect that prevents the existing acceptance criteria from being met.
- Limit implementation to documentation. Do not alter the CLI, templates, prompts, personas, roadmap, or capabilities in this task.

## Non-goals

- Implementing or restructuring any CLI command.
- Repairing the legacy initializer or deciding its compatibility-wrapper disposition.
- Implementing prompt rendering or direct agent execution.
- Adding artifact parsers, validators, state-detection code, or automated recovery.
- Correcting every stale path or persona reference in existing repository documentation.
- Adding optional `handoff`, `abandon`, or `doctor` contracts beyond identifying them as future work.
- Updating capability truth before implementation is reviewed and archived.
- Specifying internal Go packages, types, libraries, or command-framework architecture for later tasks.

## Working assumptions

- The roadmap's command list and outcome are authoritative product scope.
- The roadmap's high-level states can be reconciled with the skill's finer workflow phases by defining one normative artifact-backed model and documenting aliases or projections, without changing the intended user journey.
- A command may leave state unchanged when it only reports status, renders a prompt, or fails before a valid mutation.
- “Prompt produced” means deterministic, inspectable handoff content; direct agent launch remains later scope.
- State detection must account for the latest sequential `review-NN.md` artifact and its exactly one outcome: `approved`, `changes-requested`, or `blocked`.
- The current shell initializer is implementation evidence and compatibility context, not the normative contract for the future `init` subcommand.

## Risks and open questions

- Existing sources use overlapping state vocabularies such as `roadmapped`, `planned`, `implemented`, `reviewed`, `implementation-in-progress`, and `implementation-complete`. The documents must choose and explain a single normative model while retaining traceability to roadmap terminology.
- `reviewed` alone is not actionable without a review outcome. The state model must represent outcome-bearing review states explicitly enough to determine whether `code`, `archive`, or escalation is next.
- Files can be missing, malformed, contradictory, or left behind after partial operations. The contract must distinguish invalid or ambiguous repository state from ordinary command misuse and provide non-destructive recovery guidance.
- `roadmap` and role commands primarily render prompts today, while later roadmap items may add mutations. The reference must state precisely which effects belong to the initial contract without claiming unimplemented behavior.
- Repository and installed-project layouts must not be conflated. Contract language should use project-root-relative canonical artifacts and identify initialization as the operation that creates them.

No unresolved question requires a Product Owner decision before implementation. Local terminology and document organization may be resolved by the developer as long as all roadmap behavior and acceptance criteria remain intact.

## Implementation phases

### Phase 1 — Reconcile authoritative inputs

Status: `complete`

- Compare CON-003 with the canonical skill, personas, handoff prompts, capability limitations, archive rules, and later dependent roadmap items.
- Inventory every state-bearing artifact and metadata field used to identify roadmap selection, task progress, latest review, review outcome, approval, and archival readiness.
- Establish a consistent vocabulary and record how the roadmap's high-level labels map to artifact-backed states.

### Phase 2 — Define the workflow state machine

Status: `complete`

- Create the normative state and transition reference. This content was initially written at `.concoct/docs/state-machine.md` and must finish at `doc/state-machine.md` under the relocation phase below.
- Define each supported state through observable artifact evidence, including uninitialized, ready, roadmap-ready/roadmapped, planned, implementation in progress or complete, review outcomes, and return to ready after archival.
- Define valid transitions, commands that trigger them, state-preserving operations, remediation loops, blocked review routing, and terminal results.
- Define invalid and ambiguous combinations, failure atomicity, precedence rules where evidence conflicts, and actionable recovery expectations.
- Include a compact transition table or equivalent representation that covers the complete happy path and repeated review-remediation cycles.

### Phase 3 — Define every command contract

Status: `complete`

- Create the reference for all seven initial commands. This content was initially written at `.concoct/docs/command-reference.md` and must finish at `doc/command-reference.md` under the relocation phase below.
- For every command, state purpose, invocation and required inputs, valid starting states, files read, persona selected, files created or updated, prompt produced, resulting state, failure conditions, and recommended next commands.
- Make read-only, prompt-rendering, mutating, and transactional effects explicit.
- Cross-reference the state-machine definitions instead of duplicating or contradicting transition logic.
- State clearly that the documents define intended behavior and do not assert current CLI implementation.

### Phase 4 — Complete initial contract validation

Status: `complete`

- Check every command contract against every applicable state and transition.
- Verify that invalid transitions have actionable errors and do not imply partial mutation.
- Trace the happy path from uninitialized through archive back to ready.
- Trace at least one repeated changes-requested remediation cycle and the blocked-review escalation path.
- Search the new documents for stale canonical paths, obsolete persona names, unsupported optional commands, and accidental claims of implemented CLI behavior.
- Run Markdown and repository hygiene checks available in the repository, update task status honestly, and add the developer-to-reviewer handoff to `notes.md`.

This phase records the completed validation of the normative content at its original, incorrect location. Its completion is retained as task history and does not imply that the corrected deliverables are ready for review.

### Phase 5 — Relocate reusable product documentation

Status: `complete`

- Move `.concoct/docs/command-reference.md` to `doc/command-reference.md` without rewriting the normative command contract.
- Move `.concoct/docs/state-machine.md` to `doc/state-machine.md` without rewriting the normative state-machine contract.
- Update repository-relative links and references affected by the move, including cross-links between the two documents and references in the active task artifacts.
- Do not edit `.concoct/roadmap.md`; its old deliverable paths are Product Owner-owned roadmap content and are explicitly outside this implementation pass.
- Remove `.concoct/docs/` if it is empty after both documents are moved.
- Inspect the moved documents for a separate defect, but limit any content correction to defects that prevent the existing CON-003 acceptance criteria from being met. Record any such correction in `notes.md`.

### Phase 6 — Reverify placement and prepare review

Status: `complete`

- Repeat the relevant content-completeness and scenario checks against the documents in `doc/`.
- Verify that links and repository-relative references resolve from the new locations.
- Verify that `.concoct/docs/` contains no reusable source documentation and is absent if empty.
- Confirm the final diff preserves the prior normative content apart from required path/reference corrections or separately recorded defects.
- Run repository hygiene checks, update the task to `implementation-complete`, and replace the superseded reviewer handoff in `notes.md` with a fresh handoff that includes the relocation work and verification evidence.

### Phase 7 — Remediate blocked-review transition contract

Status: `complete`

- Define structured, artifact-backed resolution evidence that preserves the latest blocked review while allowing an authorized return to implementation or review.
- Assign resolution-record ownership for Product Owner, Task Planner, Developer, and human-originated blocker routes without adding an optional command.
- Align detection precedence, `code` and `review` starting states, invalid-state diagnostics, transition tables, and explicit blocked-review traces.
- Record Review 01 Finding 1's disposition, re-run contract checks, and prepare a fresh reviewer handoff.

## Acceptance criteria

- `doc/command-reference.md` exists and documents all seven initial commands.
- Every command contract explicitly covers purpose, required inputs, valid starting states, files read, selected persona, files created or updated, prompt produced, resulting state, failure conditions, and recommended next commands.
- `doc/state-machine.md` defines every supported workflow state using observable repository evidence rather than conversational state.
- The state-machine document identifies valid, invalid, ambiguous, and state-preserving transitions and assigns each valid transition to an initial command.
- The documented happy path proceeds from uninitialized through initialization, roadmap work, planning, implementation, review, approval, and archival back to ready without requiring `concoct handoff`.
- The model supports one or more `changes-requested → code → review` remediation cycles while preserving sequential review artifacts.
- The model routes `blocked` review outcomes to the responsible role or human without treating them as approval or silently mutating state.
- `concoct archive` is valid only with approval evidence, preserves the complete task and review history, reconciles capability and roadmap records, and clears active artifacts only after durable archive completion.
- Invalid transitions and malformed or contradictory artifact states have clear, actionable, non-destructive errors or recovery guidance.
- The command reference and state machine use consistent state names, artifact paths, persona names, outcomes, and next-command recommendations.
- Optional later commands are not required by any initial happy path or recovery path.
- The documents distinguish normative planned behavior from the currently limited shell initializer and do not claim that the command surface is implemented.
- The normative content previously written under `.concoct/docs/` is preserved except for required path/reference corrections or a separately identified and recorded acceptance-blocking defect.
- Repository-relative links and references within the moved documents and active task artifacts resolve to `doc/command-reference.md` and `doc/state-machine.md` as applicable.
- No reusable source documentation remains under `.concoct/docs/`, and the directory is removed if empty.
- No source code, tests, templates, prompts, personas, roadmap items, capabilities, or archived artifacts are changed as part of implementation.

## Verification

- Build a command-completeness matrix covering all seven commands and the ten required contract fields; confirm there are no gaps.
- Trace the documented state table for:
  - the full happy path;
  - at least one repeated remediation loop;
  - an approved archive transition;
  - a blocked review outcome;
  - representative invalid and ambiguous artifact combinations.
- Cross-check new paths and persona names against `.codex/skills/concoct/SKILL.md`, `.concoct/personas/`, and `.concoct/prompts/`.
- Cross-check state and command claims against CON-003 and the acceptance expectations of dependent items CON-005 through CON-009.
- Run `git diff --check`.
- Use `rg` to check the new documents for obsolete paths and persona names discovered during inspection.
- Compare the relocated files with their pre-relocation content and confirm that differences are limited to necessary path/reference corrections or separately recorded acceptance-blocking defects.
- Search the repository for `.concoct/docs/command-reference.md` and `.concoct/docs/state-machine.md`; classify any remaining match and update every in-scope reference. The Product Owner-owned roadmap match is expected to remain unchanged under this task's explicit non-goal and must be recorded in `notes.md` rather than silently edited.
- Confirm `doc/command-reference.md` links to `doc/state-machine.md` successfully from its new location.
- Confirm `.concoct/docs/` is absent when empty and contains no reusable source documentation in all cases.
- If a Markdown linter is available without adding dependencies, run it; otherwise record the manual structure and link checks in `notes.md`.
- Inspect the final diff to confirm only the relocated deliverable documents and active planning artifacts changed, with deletions at the old documentation paths represented as moves where Git can detect them.

## Capability impact

Expected impact: `update` to `CAP-001`.

If accepted and archived, Concoct's durable file-based workflow contract will include a normative command reference and artifact-backed state machine. This task does not add an implemented CLI capability and does not change `CAP-002`'s existing manual prompt behavior. The reviewer and archivist must validate the final capability wording against delivered documentation.

## Handoff expectations

The developer should:

- read `AGENTS.md`, `.concoct/personas/developer.md`, `.concoct/capabilities.md`, this plan, these notes, and the relevant archive summary before editing;
- keep implementation limited to the two CON-003 documentation deliverables plus honest updates to active task artifacts;
- move the existing documents from `.concoct/docs/` to `doc/`, update affected in-scope links and references, and remove the old directory if empty;
- preserve the completed normative contract and avoid broad rewriting during relocation;
- preserve product-level command and transition behavior while making local terminology and presentation decisions explicit in `notes.md`;
- record any contradiction that cannot be reconciled without changing product intent and stop for the appropriate owner rather than inventing behavior;
- verify placement and link correctness in addition to rechecking the command/state matrix and scenario traces;
- finish with updated phase statuses and a reviewer handoff covering files changed, decisions, checks, risks, skipped work, capability impact, and suggested review focus;
- recommend `concoct review` when implementation is complete.
