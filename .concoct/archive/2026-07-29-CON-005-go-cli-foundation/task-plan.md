---
id: CON-005
title: Implement the Concoct CLI foundation in Go
roadmap-id: CON-005
status: implementation-complete
created: 2026-07-29
updated: 2026-07-29
capability-impact:
  type: add
  ids:
    - CAP-005
  rationale: Adds an executable Go CLI for project initialization and repository-backed workflow status; delivery should also reconcile CAP-003's initializer limitation.
---

# Task Plan

## Goal

Implement a small Go CLI providing `concoct init <project>` and `concoct status`, with shared project discovery, artifact parsing, state validation, diagnostics, and template-copying infrastructure suitable for later commands.

## Context

CON-003 delivered the normative command contract and artifact-backed state machine that govern this implementation. The repository currently has only a legacy shell initializer at `cmd/concoct/concoct`; it is not a command-oriented CLI and cannot initialize from its checked-in location because it resolves `templates/` and persona sources relative to `cmd/concoct/`.

This task is the executable foundation for CON-006 and later workflow commands. It must preserve the useful bootstrap behavior while making status derivation conform to `doc/command-reference.md` and `doc/state-machine.md`.

## Why this matters

Concoct's workflow is currently a documented and manually operated contract. A working CLI foundation restores reliable project bootstrap and makes repository state observable without asking an agent to infer it, while establishing reusable infrastructure for the rest of the command surface.

## Current state

- No Go module, Go source, or automated test suite exists.
- `cmd/concoct/concoct` is an executable Bash initializer with a legacy positional interface.
- The shell initializer copies templates including dotfiles, writes bootstrap guidance, initializes Git, stages generated files, and creates no commit.
- The initializer expects `templates/` and `personas/` beside itself, but reusable templates live at repository-root `templates/` and personas are already included beneath `templates/.concoct/personas/`; end-to-end bootstrap therefore fails before creating a target.
- `doc/command-reference.md` defines the intended `init` and read-only `status` contracts.
- `doc/state-machine.md` defines canonical artifacts, metadata, state precedence, invalid evidence, remediation fields, and next actions.
- `.concoct/current/` contains only `.gitkeep`; no active task conflicts with this plan.

## Target state

- A Linux-buildable Go CLI exposes `init` and `status` through a small, explicit command registry and shared infrastructure.
- `init` resolves its distributed templates independently of the caller's working directory, copies the complete template including dotfiles and nested content, writes any required bootstrap guidance, initializes Git, stages generated files, creates no commit, validates the result, and reports partial-target recovery information on failure.
- `status` discovers a Concoct-enabled project from the project root or a nested directory, parses the canonical metadata it needs, derives every state required by CON-003, and prints relevant identity, task, review, capability-impact, diagnostics, and next-command fields without mutation.
- Automated tests cover discovery, metadata parsing, state validation, diagnostics, template copying, and end-to-end initialization.
- The legacy initializer is either a thin compatibility wrapper for the Go CLI or is removed with a documented migration path.

## Design constraints

- Treat `AGENTS.md`, `doc/command-reference.md`, and `doc/state-machine.md` as normative; do not simplify state behavior in ways that contradict them.
- Keep the implementation agent-neutral and free of direct agent execution or tool-specific workflow logic.
- Prefer the Go standard library and small explicit internal boundaries; avoid a framework-heavy command architecture.
- Resolve project state from durable repository artifacts, never conversation history, timestamps, or narrative claims.
- Parse required YAML front matter predictably and validate known state-bearing fields, including review sequencing, `remediates-review`, and `blocked-review-resolution` where applicable to state detection.
- Empty tracked current-task placeholders do not constitute an active task; partially populated or contradictory artifacts are invalid.
- `status` is strictly read-only, including for invalid or partially archived states.
- Initialization must copy root files, nested files, and dotfiles; preserve the generated `.concoct/` versus conventional-root-file layout.
- Preserve existing bootstrap semantics by staging generated files and creating no commit. The command must report that staging occurred.
- Preserve executable entry-point usability and document the chosen compatibility or migration behavior.
- Do not initialize generated test projects inside this repository.

## Non-goals

- Implementing `roadmap`, `plan`, `code`, `review`, `archive`, `handoff`, `abandon`, or `doctor`.
- Rendering role prompts beyond bootstrap guidance required by `init`.
- Launching or integrating with coding agents.
- Mutating malformed workflow artifacts or automatically recovering workflow state.
- Implementing archive transactions, capability reconciliation, upgrades, overlays, task branches, or remote execution.
- Broadly rewriting templates, personas, prompts, or documentation unrelated to making `init` and `status` correct.
- Creating an initial Git commit in generated projects.

## Working assumptions

- The Go module and executable layout may be selected by the Developer after inspecting repository conventions, provided users retain a clear `concoct` entry point and later commands can reuse the internal boundaries.
- A packaged or otherwise deterministically located copy of repository-root `templates/` is the initialization source; tests must prove operation is independent of the caller's working directory.
- Existing staging behavior is part of bootstrap compatibility because the shell implementation, README, and archived history all describe it. No Product Owner decision is needed unless implementation evidence reveals a material user-facing conflict.
- `CAP-005` is the proposed identifier for the new executable CLI capability. The Archivist must validate that identifier and update CAP-003's initializer limitation only after accepted delivery.
- Human-readable output need not be byte-for-byte identical to the roadmap example, but all applicable fields and actionable diagnostics must be present and deterministic.

## Risks and open questions

- Template discovery differs between development, installed binaries, and test binaries. The implementation must choose and document a distribution strategy that works outside the repository checkout without embedding repository-local absolute paths.
- YAML parsing may tempt either a large dependency or an incomplete ad hoc parser. The chosen approach must support the checked-in schemas and reject malformed state-bearing metadata clearly.
- The normative state machine includes remediation, blocked-review recovery, and partial archival contradictions even though commands for producing those artifacts arrive later; `status` must not defer those detection rules.
- Transactional initialization can leave partial targets after filesystem or Git failures. Cleanup is allowed only when ownership and safety are provable; otherwise diagnostics must preserve and identify the target.
- Legacy entry-point migration could break documented invocation or executable permissions if not tested.
- No unresolved product decision blocks implementation. If repository inspection forces a choice that changes observable bootstrap or status behavior beyond the roadmap and CON-003 contract, return that choice to the Product Owner rather than deciding it in code.

## Implementation phases

### Phase 1 — Establish the Go CLI and distribution boundary

Status: `complete`

- Add the minimal Go module, executable entry point, explicit command dispatch, usage, and exit-code behavior.
- Define a template-distribution strategy and shared filesystem abstractions that support installed use and isolated tests.
- Decide the legacy script's wrapper or documented-removal path without maintaining two independent implementations.

### Phase 2 — Implement project discovery and artifact models

Status: `complete`

- Locate a Concoct project reliably from its root and nested directories, with explicit outside-project diagnostics.
- Model and parse roadmap, capability, task-plan, notes, review, and archive-reference evidence needed by `status`.
- Validate identifiers, known status and impact values, required metadata, placeholder versus populated artifacts, review numbering, and cross-artifact consistency.

### Phase 3 — Implement deterministic state detection and status reporting

Status: `complete`

- Encode the CON-003 detection order and precedence for `ready`, `planned`, `implementation-in-progress`, `implementation-complete`, all review outcomes, and `invalid`.
- Cover remediation and blocked-review-resolution evidence plus representative partial-archive contradictions.
- Render applicable project, roadmap, task, review, capability-impact, diagnostic, and next-action fields without modifying the repository.

### Phase 4 — Implement initialization and compatibility behavior

Status: `complete`

- Validate target arguments and safety, create a new project from the complete template, and ensure expected planning and archive directories exist.
- Produce bootstrap guidance consistent with the installed personas and current workflow transition.
- Initialize Git, stage generated files, create no commit, and validate the resulting project as `ready`.
- Provide precise partial-target and recovery reporting for failures, and complete the legacy wrapper or migration documentation.

### Phase 5 — Verify, document, and prepare review

Status: `complete`

- Add focused unit tests for parsing and state logic plus integration tests using temporary directories and real Git where required.
- Exercise the built CLI from outside the repository root and from nested directories.
- Update user and developer documentation only where the executable location, invocation, staging behavior, build/test commands, or migration changed.
- Update this plan and notes with implementation status, checks, durable decisions, risks, and reviewer handoff evidence.

### Phase 6 — Remediate Review 01

Status: `complete`

- Make retained remediation and blocked-review-resolution metadata historical
  when a later valid review exists, without accepting references that do not
  identify an earlier review of the required outcome.
- Validate the canonical roadmap status vocabulary and required task/review
  metadata with file-and-field-specific diagnostics.
- Add focused regression tests for later-review supersession and missing or
  unknown state-bearing metadata.
- Re-run the complete verification suite and prepare a fresh reviewer handoff.

### Phase 7 — Remediate Review 02

Status: `complete`

- Let historical recovery metadata fall through to evaluation of recovery
  evidence that applies to the latest review.
- Preserve strict validation for missing, wrong-outcome, and otherwise invalid
  historical recovery references.
- Add regression coverage that composes historical changes-requested
  remediation with resolution of a later blocked review.
- Re-run the complete verification suite and prepare a fresh reviewer handoff.

## Acceptance criteria

- CON-005 builds as a Go CLI on Linux and exposes working `concoct init <project>` and `concoct status` commands with clear usage and non-zero operational-error exits.
- `init` succeeds end to end in a temporary parent directory when invoked independently of the caller's working directory.
- The generated project contains all root templates, nested content, dotfiles, personas, prompts, planning directories, and bootstrap guidance; it is a Git repository whose generated files are staged and whose history has no automatically created commit.
- `init` refuses unsafe, empty, extra-argument, or existing targets without overwriting them, and any creation-time failure identifies whether a partial target remains and gives a safe next action.
- `status` discovers the same project from its root and a nested directory and does not alter files, Git state, or workflow state.
- `status` correctly reports `ready`, `planned`, `implementation-in-progress`, `implementation-complete`, `review-changes-requested`, `review-approved`, and `review-blocked`, with the correct recommended next action.
- `status` treats malformed metadata, mismatched roadmap/task identities, incomplete active artifacts, invalid review sequences or outcomes, stale remediation evidence, invalid blocked-review resolution, and representative interrupted-archive contradictions as actionable `invalid` diagnostics rather than guessing a state.
- Status output includes every applicable roadmap-required field: project name, active roadmap item, phase, task status, latest review, review outcome, capability impact, and recommended next command.
- Tests cover project-root discovery, front-matter parsing, transition/state validation, template copying, Git behavior, status read-only behavior, and the end-to-end initialization path.
- The legacy initializer is retained only as a tested compatibility wrapper or removed with clear migration documentation; executable permissions remain correct for any retained script.
- Documentation accurately describes the Go CLI invocation, template distribution, staging/no-commit behavior, and verification commands.
- `go test ./...`, relevant build/static checks, `git diff --check`, and the repository's template/initialization checks pass.

## Verification

- Run `go test ./...`.
- Run `go vet ./...` unless the selected module/tooling provides an equivalent stricter documented check.
- Build the CLI for Linux using the documented command and run its help plus both subcommands.
- Run table-driven state tests covering every normative state, required invalid cases, review precedence, remediation, and blocked-review resolution.
- Run an end-to-end `init` test under a temporary parent outside this repository; compare required template paths including dotfiles, verify bootstrap guidance and personas, run `git status --short`, and confirm `git log` has no initial commit.
- Run `status` from the generated project root and a nested directory; snapshot repository and Git state before and after to prove read-only behavior.
- Run `bash -n` on any retained shell wrapper and confirm it delegates rather than duplicating initialization logic.
- Run `git diff --check` and search changed files for stale executable paths, legacy invocation, persona names, and branding relevant to this task.

## Capability impact

Expected impact is `add`: proposed `CAP-005` will describe the executable Go CLI's initialization and workflow-status behavior. Accepted delivery is also expected to update `CAP-003` by removing or revising its broken-initializer limitation and replacing its verification evidence. Capability truth must not change until review approval and archival reconciliation.

## Handoff expectations

The Developer should begin by re-reading CON-003's command and state-machine contracts, then validate this plan against the selected Go/template distribution approach before implementation. Keep changes limited to the CLI foundation, its tests, necessary template/bootstrap corrections, and directly affected documentation.

Before review, the Developer must:

- update phase and task status honestly;
- record architecture and distribution decisions in `notes.md`;
- record all commands run and their outcomes, including any skipped check;
- inspect the full diff for unrelated template or documentation churn;
- identify files changed, compatibility behavior, known risks, and capability impact;
- provide a reviewer handoff emphasizing state-machine fidelity, invalid-state diagnostics, initialization safety, status non-mutation, installed template discovery, and legacy migration.

When implementation is complete, recommend `concoct review`.
