---
id: CON-006
title: Implement deterministic prompt rendering
roadmap-id: CON-006
status: implementation-complete
created: 2026-07-29
updated: 2026-07-30
capability-impact:
  type: add
  ids:
    - CAP-006
  rationale: Adds executable, role-aware generation of complete and reproducible workflow prompts from validated repository state.
---

# Task Plan

## Goal

Implement deterministic, inspectable prompt rendering for `concoct roadmap`,
`concoct plan <roadmap-id>`, `concoct code`, and `concoct review` without
launching an agent or treating rendered output as completed role work.

## Context

CON-003 defines the normative command and artifact-backed state contracts.
CON-005 delivered the Go CLI, project discovery, strict artifact validation,
and deterministic state detection for `init` and `status`. Concoct also ships a
manual prompt corpus under `.concoct/prompts/` and the embedded project
template. This task connects those accepted contracts and assets to the CLI so
users can obtain the correct role handoff directly from repository state.

## Why this matters

Users currently select and supply transition prompts manually. Executable
rendering makes the prompt-only workflow consistent and testable while
preserving Concoct's agent-neutral boundary and leaving later commands to own
workflow mutation and direct execution.

## Current state

- The Go CLI recognizes only `init`, `status`, and help.
- Project discovery and `workflow.Detect` already expose the validated current
  state and key active-task/review fields needed for command eligibility.
- Manual role-transition prompts exist in both the repository and embedded
  templates, but no command selects or renders them.
- CAP-002 describes manual prompts; CAP-005 describes the executable CLI and
  state reporting. No current capability renders role prompts.

## Target state

- The four roadmap commands validate their arguments and allowed starting
  states, select the intended persona and mode, and emit a complete prompt to
  stdout by default.
- `--output <path>` writes the identical rendered bytes to the requested file
  through a safe, explicit file-output path.
- Rendering is deterministic for the same repository state and command input,
  including stable ordering of archive summaries and review files.
- Prompts name the exact repository-relative inputs to read, exact artifact
  ownership boundaries, detected state, expected outcome, validation and
  completion requirements, and recommended next transition.
- Rendering and validation failures do not mutate workflow artifacts or
  otherwise establish a new workflow state.

## Design constraints

- Follow `doc/command-reference.md` and `doc/state-machine.md`; durable
  artifacts, not conversation history or rendered output, determine state.
- Reuse the delivered project-discovery and workflow-validation boundaries
  rather than introducing a second state model.
- Keep command handling and rendering small, explicit, and free of a
  framework-heavy CLI architecture.
- Keep prompts agent-neutral and fully inspectable.
- Preserve the manual prompt corpus and embedded template as product assets;
  avoid divergent role rules between rendered and manually supplied prompts.
- Use repository-relative paths in rendered guidance and stable ordering for
  every discovered input.
- Treat invalid evidence or invalid command transitions as actionable errors
  with no output-file or workflow mutation.
- Preserve existing `init`, `status`, wrapper, and installed-binary behavior.

## Non-goals

- No direct agent invocation, configuration, supervision, or permission
  management from CON-010.
- No implementation of role-completion mutations: roadmap editing, active-task
  creation, task-status transitions, review allocation/writes, or archival.
- No Git task-branch lifecycle or integration behavior from CON-015.
- No general-purpose template engine or arbitrary Markdown rendering system.
- No prompt snapshot committed as routine generated output.
- No broad cleanup of stale template persona references or empty writer
  personas unless a directly consumed CON-006 prompt cannot render correctly
  without a narrowly scoped correction.
- No structured review-finding identifier scheme; remediation rendering uses
  the accepted review and notes contract.

## Working assumptions

- The checked-in manual transition prompts, personas, command reference, and
  state-machine contract provide enough product guidance to compose the
  rendered prompts without a new prompt language.
- Output formatting and safe file-write mechanics are technical design choices
  so long as stdout and file output are byte-identical, deterministic, and do
  not silently destroy unrelated data.
- Relevant archive summaries can be selected deterministically from explicit
  roadmap dependencies, task references, and current archive metadata, with
  stable ordering and conservative inclusion when relevance is ambiguous.
- `plan` only renders the Task Planner handoff in this task; CON-007 remains
  responsible for establishing active task artifacts and roadmap status.
- `review` renders the next-review handoff and may report the expected next
  review filename, but CON-008 remains responsible for collision-safe review
  allocation and persistence.

## Risks and open questions

- The current `workflow.Report` is optimized for status display and may not
  expose all validated metadata required for rendering. The implementation
  must extend or add a read-only context boundary without duplicating parsing
  rules or weakening validation.
- Repository prompts and embedded template prompts can drift if rendering
  introduces a second independently maintained body of role guidance.
- Deterministic relevance selection for archive summaries needs focused tests;
  filesystem iteration order and incidental timestamps must not affect output.
- State-specific `code` prompts must distinguish initial development,
  implementation continuation, changes-requested remediation, and a validated
  blocked-review route to code. `review` must distinguish initial review,
  post-remediation review, and a validated blocked-review route to review.
- File-output failure behavior must avoid partial or misleading output. The
  developer should choose a safe atomic or refusal strategy consistent with
  established CLI error handling and document it.
- No unresolved product question blocks implementation. Return to the Product
  Owner only if implementation would change which roles are selected, authorize
  workflow mutations, launch agents, or omit roadmap-required prompt content.

## Implementation phases

### Phase 1 — Confirm command and rendering context boundaries

Status: `complete`

- Trace each command's valid starting states, persona, inputs, ownership rules,
  completion evidence, and next transition through the normative documents and
  manual prompts.
- Define one read-only rendering context derived from validated repository
  artifacts, including deterministic archive/review ordering and selected
  roadmap-item validation.
- Confirm how repository and embedded prompt assets remain aligned without
  creating duplicate durable role rules.

### Phase 2 — Add deterministic prompt composition

Status: `complete`

- Compose complete Product Owner, Task Planner, Developer, and Reviewer prompts
  from the validated context and accepted prompt corpus.
- Include exact inputs, writable artifacts, current state/mode, expected role
  outcome, required validation, completion evidence, and next transition.
- Make remediation and blocked-review recovery modes explicit while preserving
  review immutability and role ownership.
- Ensure all discovered collections and rendered sections have stable ordering
  and reproducible whitespace.

### Phase 3 — Expose the four CLI commands and output modes

Status: `complete`

- Extend CLI usage, dispatch, argument validation, project discovery, state
  eligibility checks, and actionable errors for `roadmap`, `plan`, `code`, and
  `review`.
- Print rendered output to stdout by default and support exactly one
  `--output <path>` destination form without changing workflow state.
- Preserve existing commands and source-checkout wrapper behavior.

### Phase 4 — Add automated behavioral and golden verification

Status: `complete`

- Add focused unit tests for context selection, stable ordering, role and mode
  selection, exact ownership language, and invalid transitions.
- Add golden tests for every supported command and the materially distinct
  development, remediation, and review modes.
- Add CLI/integration tests showing stdout/file equivalence, installed or
  external-directory operation, non-mutation, repeatability, and safe failures.

### Phase 5 — Document and prepare independent review

Status: `complete`

- Update CLI usage and user documentation for the four prompt-only commands,
  `--output`, state preservation, and the boundary with future role-execution
  commands.
- Run the complete verification suite, inspect the final diff, record durable
  results and decisions, update phase statuses honestly, and add the reviewer
  handoff required by the Developer persona.

## Acceptance criteria

- `concoct roadmap`, `concoct plan <roadmap-id>`, `concoct code`, and
  `concoct review` render the correct role prompt from every state allowed by
  the normative command contract and reject disallowed or invalid states with
  actionable diagnostics.
- Every rendered prompt includes the selected persona, exact files to read,
  exact files or areas the role may update, detected workflow state and mode,
  expected outcome, validation and completion requirements, and recommended
  next transition.
- `plan` validates the requested identifier and renders only for an existing,
  eligible roadmap item with satisfied or explicitly handled dependencies and
  no conflicting active task.
- `code` explicitly distinguishes initial/continued implementation from
  changes-requested remediation and validated blocked-review recovery; it never
  authorizes modification of completed review files.
- `review` identifies prior review context and the expected next sequential
  review path without modifying or overwriting any review artifact.
- Repeating the same command against unchanged repository bytes produces
  byte-identical prompt output, with deterministic ordering independent of
  filesystem enumeration.
- Default stdout and `--output <path>` contain identical prompt bytes, and
  invalid arguments, invalid state, or output failure leave workflow artifacts
  unchanged and do not leave a partial generated prompt.
- Rendering alone does not change detected workflow state, roadmap status,
  current task artifacts, reviews, capabilities, archives, or Git status.
- Existing `init` and `status` behavior and the source-checkout wrapper remain
  compatible, and the CLI continues to build and run on Linux.
- Golden tests cover all four roles plus materially distinct code/review modes,
  and documentation explains prompt-only behavior and generated-output policy.

## Verification

- Run `gofmt` on changed Go files.
- Run `go test ./...` including golden, state/mode matrix, CLI argument,
  non-mutation, ordering, and stdout/file equivalence tests.
- Run `go vet ./...`.
- Build the Linux CLI to a temporary path and exercise help plus all four
  rendering commands from project-root and nested-directory contexts.
- In isolated temporary repositories, compare repeated output byte-for-byte;
  compare stdout with `--output`; snapshot workflow artifacts and Git status
  before and after successful and failed rendering.
- Run `bash -n cmd/concoct/concoct.sh` and confirm executable permissions if the
  wrapper changes.
- Run `git diff --check`.
- Search changed source, tests, docs, and prompt assets for stale persona names,
  obsolete command syntax, agent-specific invocation assumptions, and generated
  prompt files that should not be committed.

## Review remediation

- Review 01 Finding 1 is fixed. Nine checked-in golden fixtures now compare the
  complete rendered bytes for all four commands and each materially distinct
  code and review mode.
- Phase 4 remains `complete` with the plan's explicit golden-test outcome now
  evidenced directly rather than by repeat-render and substring checks alone.
- The full test suite, vet, shell syntax and executable checks, and
  `git diff --check` pass after remediation.

## Capability impact

Expected impact is `add`, proposing `CAP-006` for executable deterministic,
role-aware prompt rendering. CAP-002 remains the current manual prompt
capability and CAP-005 remains the CLI initialization/status capability; the
Archivist should decide whether their limitations or related-capability links
also need reconciliation after accepted delivery.

## Developer handoff expectations

Before implementation, revalidate the command/state matrix and prompt-asset
alignment against current repository reality. During implementation, keep
technical decisions and meaningful failed attempts in `notes.md`, update phase
statuses accurately, and do not edit roadmap, capabilities, archive, or review
artifacts. Before requesting review, provide the implemented behavior, key
decisions, files changed, checks run, known risks, skipped work, capability
impact, and suggested review focus, with particular attention to deterministic
output, state/mode selection, non-mutation, prompt drift, and ownership safety.
