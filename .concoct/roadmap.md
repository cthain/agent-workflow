---
version: 1
project: concoct
updated: 2026-07-28
---

# Roadmap

## Purpose

This roadmap defines the planned evolution of Concoct.

Concoct is a lightweight, agent-neutral workflow coordinator that turns rough ideas into implementation-ready work and guides transitions between planning, implementation, review, and archival roles.

Core loop:

```text
idea → Chappie concocts the plan → agents cook → eatin' big time
```

This file records intended future work.

It is distinct from:

- `.concoct/capabilities.md`, which records what Concoct can do now;
- `.concoct/current/task-plan.md`, which records the currently active implementation task;
- `.concoct/archive/`, which preserves completed task history.

## Roadmap conventions

Each roadmap item has a stable identifier.

Statuses:

- `candidate` — accepted as a possible direction but not yet ordered for implementation;
- `planned` — ready to be turned into an active task plan;
- `active` — currently represented by `.concoct/current/task-plan.md`;
- `blocked` — cannot proceed until a dependency or decision is resolved;
- `delivered` — implemented, reviewed, archived, and reflected in `capabilities.md`;
- `deferred` — intentionally postponed;
- `cancelled` — no longer intended.

Priorities:

- `critical`
- `high`
- `medium`
- `low`

A roadmap item should describe a coherent outcome. Detailed implementation steps belong in the task plan created by:

```text
concoct plan <roadmap-id>
```

---

## CON-001 — Establish the canonical Concoct project structure

- Status: `planned`
- Priority: `critical`
- Depends on: none
- Capability impact: updates existing project structure and documentation

### Outcome

Restructure the repository so Concoct-owned artifacts live under `.concoct/`, while tool-discovery files remain in their conventional root-level locations.

### Target structure

```text
.
├── AGENTS.md
├── README.md
├── LICENSE
├── init
├── .codex/
│   └── skills/
│       └── concoct/
│           └── SKILL.md
└── .concoct/
    ├── capabilities.md
    ├── roadmap.md
    ├── current/
    ├── archive/
    ├── personas/
    ├── prompts/
    ├── docs/
    └── templates/
```

### Requirements

- Keep `AGENTS.md` at the repository root.
- Keep `.codex/` at the repository root.
- Move Concoct-owned documentation, prompts, personas, templates, active artifacts, and archives under `.concoct/`.
- Preserve root-level conventional project files.
- Preserve executable permissions on scripts.
- Update all internal links and path references.
- Remove stale references to the previous project identity and layout.
- Clearly distinguish Concoct's own workflow artifacts from templates installed into generated projects.

### Acceptance criteria

- The repository has one authoritative structure.
- All documentation and scripts use current paths.
- No meaningful stale references remain.
- `bash -n init` succeeds.
- Existing bootstrap behavior still works after path changes.

---

## CON-002 — Define the Concoct artifact model

- Status: `planned`
- Priority: `critical`
- Depends on: CON-001
- Capability impact: adds a formal artifact contract

### Outcome

Define the durable files Concoct reads and writes, their responsibilities, and their machine-readable metadata.

### Artifacts

```text
.concoct/capabilities.md
.concoct/roadmap.md
.concoct/current/task-plan.md
.concoct/current/notes.md
.concoct/current/review-NN.md
.concoct/archive/YYYY-MM-DD-short-task-name/
```

### Requirements

Create an artifact schema that defines:

- required front matter;
- stable identifiers;
- valid statuses;
- ownership by role;
- append-only versus mutable artifacts;
- archive naming;
- capability-impact declarations;
- review outcomes;
- handoff information.

Suggested metadata for task plans:

```yaml
---
id: CON-XXX
title: Short task title
roadmap-id: CON-XXX
status: planned
created: YYYY-MM-DD
updated: YYYY-MM-DD
capability-impact:
  type: add | update | remove | none
  ids: []
  rationale:
---
```

Suggested review metadata:

```yaml
---
task-id: CON-XXX
review: 1
status: approved | changes-requested | blocked
created: YYYY-MM-DD
persona: code-reviewer
---
```

### Acceptance criteria

- `.concoct/docs/artifact-schema.md` defines every durable artifact.
- Human-readable Markdown remains the source of truth.
- Metadata is sufficient for deterministic CLI validation.
- The design does not require a separate opaque state database.

---

## CON-003 — Define the command contract and workflow state machine

- Status: `planned`
- Priority: `critical`
- Depends on: CON-002
- Capability impact: defines the intended CLI behavior

### Outcome

Define each Concoct command, the files it reads and writes, the role it selects, valid starting states, resulting states, and failure behavior.

### Initial command surface

```text
concoct init <project>
concoct status
concoct roadmap
concoct plan <roadmap-id>
concoct code
concoct review
concoct archive
```

Optional later commands:

```text
concoct handoff
concoct abandon
concoct doctor
```

### Workflow states

```text
uninitialized
    ↓ init
ready
    ↓ roadmap
roadmapped
    ↓ plan <id>
planned
    ↓ code
implemented
    ↓ review
reviewed
    ├── changes requested → code
    └── approved → archive
                     ↓
                   ready
```

### Requirements

For each command, document:

- purpose;
- required inputs;
- valid starting states;
- files read;
- persona selected;
- files created or updated;
- prompt produced;
- resulting state;
- failure conditions;
- recommended next commands.

Create:

```text
.concoct/docs/command-reference.md
.concoct/docs/state-machine.md
```

### Acceptance criteria

- Every initial command has a complete contract.
- Invalid transitions have clear, actionable errors.
- `handoff` is not required in the normal happy path.
- Each role command performs its own incoming and outgoing handoff.
- The model supports implementation-remediation loops after review.

---

## CON-004 — Establish `capabilities.md` as current accepted product truth

- Status: `planned`
- Priority: `high`
- Depends on: CON-002
- Capability impact: adds the capability ledger

### Outcome

Create `.concoct/capabilities.md` as the canonical human-readable record of what Concoct can do now.

### Definition

`capabilities.md` is:

- a record of current accepted capabilities;
- updated only after completed work is reviewed and archived;
- readable by humans and agents;
- distinct from the roadmap, backlog, changelog, and archive.

It is not:

- a list of ideas;
- a future roadmap;
- a task history;
- a design proposal.

### Requirements

- Give capabilities stable identifiers such as `CAP-001`.
- Describe observable behavior rather than implementation history.
- Allow roadmap items and archive summaries to reference capability IDs.
- Require every archived task to declare whether it adds, updates, removes, or does not affect capabilities.
- Update capabilities from delivered outcomes and actual code, not promises in the task plan.

### Acceptance criteria

- `.concoct/capabilities.md` exists with an initial capability inventory.
- The archive workflow defines how capability changes are reconciled.
- Internal-only refactors can explicitly declare no capability impact.
- Roadmap planning reads `capabilities.md` before proposing future work.

---

## CON-005 — Implement the Concoct CLI foundation in Go

- Status: `planned`
- Priority: `critical`
- Depends on: CON-001, CON-002, CON-003
- Capability impact: adds the executable CLI foundation

### Outcome

Replace the growing shell-based workflow coordinator with a small Go CLI while preserving the existing bootstrap behavior.

### Initial scope

Implement:

```text
concoct init
concoct status
```

Provide command registration and shared infrastructure for later commands.

### Requirements

- Resolve the project root reliably.
- Detect whether a repository is Concoct-enabled.
- Read and validate artifact metadata.
- Produce clear errors and suggested next actions.
- Preserve agent-neutral behavior.
- Keep CLI internals small and explicit.
- Avoid a framework-heavy command architecture.
- Continue to support template copying, including dotfiles.
- Do not create an initial Git commit automatically.
- Decide deliberately whether generated files should be staged.

### `concoct status`

Report:

- project name;
- active roadmap item;
- workflow phase;
- task status;
- latest review;
- review outcome;
- capability impact;
- recommended next command.

Example:

```text
Project: Portage
Roadmap item: PORT-003 — Implement CLT transfer retries
Phase: review
Task status: implementation-complete
Latest review: review-02.md
Review outcome: changes-requested

Next:
  concoct code
```

### Acceptance criteria

- The CLI builds and runs on Linux.
- `concoct init` passes an end-to-end temporary-directory test.
- `concoct status` correctly handles ready, planned, implemented, changes-requested, approved, and invalid states.
- Tests cover path resolution, metadata parsing, transition validation, and template copying.
- The existing `init` script is either retained as a compatibility wrapper or removed with documented migration.

---

## CON-006 — Implement deterministic prompt rendering

- Status: `planned`
- Priority: `high`
- Depends on: CON-003, CON-005
- Capability impact: adds role-aware prompt generation

### Outcome

Generate complete, inspectable prompts for each workflow role without directly launching an agent.

### Commands

```text
concoct roadmap
concoct plan <roadmap-id>
concoct code
concoct review
```

### Requirements

Each generated prompt must include:

- the selected persona;
- the exact files to read;
- the exact files the role may update;
- the current workflow state;
- the expected outcome;
- validation and completion requirements;
- the recommended next transition.

Prompt generation should be deterministic from repository state.

Default behavior:

```text
concoct code
```

prints the prompt to stdout.

Optional output:

```text
concoct code --output <path>
```

Generated prompts are reproducible output and should not be committed by default.

### Role inputs

#### Roadmap

Reads:

- `AGENTS.md`
- `.concoct/personas/planner.md`
- `.concoct/capabilities.md`
- `.concoct/roadmap.md`
- relevant archive summaries

Updates:

- `.concoct/roadmap.md`

#### Plan

Reads:

- `AGENTS.md`
- planner persona
- `.concoct/capabilities.md`
- `.concoct/roadmap.md`
- relevant archive artifacts

Creates or updates:

- `.concoct/current/task-plan.md`
- `.concoct/current/notes.md`

#### Code

Reads:

- `AGENTS.md`
- developer persona
- current task plan
- current notes
- latest review, when present
- relevant capability context

Updates:

- code and tests
- current task plan
- current notes

#### Review

Reads:

- `AGENTS.md`
- reviewer persona
- current task plan
- current notes
- repository diff
- prior reviews

Creates:

- next `.concoct/current/review-NN.md`

### Acceptance criteria

- Prompts select the correct role and context.
- Review-remediation mode is handled explicitly.
- Developer prompts never authorize modification of completed review files.
- Prompt output can be tested with golden files.
- No agent-specific invocation is required.

---

## CON-007 — Implement active task planning

- Status: `planned`
- Priority: `high`
- Depends on: CON-004, CON-005, CON-006
- Capability impact: adds roadmap-to-task transition

### Outcome

Implement:

```text
concoct plan <roadmap-id>
```

to turn one roadmap item into an active task-planning session.

### Requirements

- Validate that the roadmap item exists and is eligible for planning.
- Refuse to overwrite an active task.
- Read current capabilities and relevant archive history.
- Generate the planner prompt.
- Create task placeholders only when appropriate.
- Update roadmap item status after a valid plan is established.
- Preserve the roadmap identifier in task metadata.

### Acceptance criteria

- A roadmap item can be selected deterministically.
- Active-task conflicts produce clear errors.
- The task plan and notes use the artifact schema.
- The generated planner prompt contains all required context.
- `concoct status` reflects the planned task.

---

## CON-008 — Implement code and review transitions

- Status: `planned`
- Priority: `high`
- Depends on: CON-005, CON-006, CON-007
- Capability impact: adds developer and reviewer coordination

### Outcome

Implement role transitions for:

```text
concoct code
concoct review
```

### Code requirements

- Read the active task plan and notes.
- Read the latest review when one exists.
- Detect normal implementation versus review remediation mode.
- Generate a developer prompt with narrow file ownership.
- Require disposition of unresolved review findings.
- Do not modify completed review artifacts.

### Review requirements

- Allocate the next review sequence safely.
- Read prior reviews for context without blindly inheriting their conclusions.
- Produce one of:
  - `approved`;
  - `changes-requested`;
  - `blocked`.
- Preserve review artifacts as append-only after completion.
- Update status without rewriting developer-owned history.

### Acceptance criteria

- The loop `code → review → code → review` works.
- Review numbers are deterministic and collision-safe.
- Status output recommends the correct next command.
- Invalid role transitions fail safely.
- Review artifacts are preserved for archive.

---

## CON-009 — Implement archive and capability reconciliation

- Status: `planned`
- Priority: `high`
- Depends on: CON-004, CON-005, CON-008
- Capability impact: adds accepted-task archival

### Outcome

Implement:

```text
concoct archive
```

as the acceptance boundary between completed work and current product truth.

### Requirements

Archive must:

- require an approved review by default;
- support a deliberate override only when explicitly requested;
- inspect the completed task, notes, reviews, and repository changes;
- determine or validate capability impact;
- update `.concoct/capabilities.md`;
- update the roadmap item to `delivered`;
- move current task artifacts into a dated archive directory;
- create `summary.md`;
- leave `.concoct/current/` ready for the next task.

### Archive structure

```text
.concoct/archive/YYYY-MM-DD-roadmap-id-short-task-name/
  task-plan.md
  notes.md
  review-01.md
  review-02.md
  summary.md
```

### Summary requirements

Include:

- task and roadmap identifier;
- delivered outcome;
- key decisions;
- files changed;
- checks run;
- review result;
- capability changes;
- skipped work;
- follow-up work.

### Acceptance criteria

- Archive refuses unapproved work by default.
- Capability truth is reconciled transactionally.
- Current artifacts are cleared only after successful archive completion.
- Roadmap, capabilities, and archive remain cross-referenced.
- The repository returns to the `ready` state.

---

## CON-010 — Add direct agent execution adapters

- Status: `candidate`
- Priority: `medium`
- Depends on: CON-006, CON-007, CON-008, CON-009
- Capability impact: adds optional agent invocation

### Outcome

Allow Concoct to launch a configured agent with the same deterministic prompt that can already be printed.

### Proposed usage

```text
concoct code --run codex
concoct review --run codex
concoct plan CON-007 --run claude
```

### Requirements

- Keep prompt generation independent from execution.
- Make agent invocation optional.
- Keep the complete rendered prompt inspectable.
- Support per-project configuration in `.concoct/config.toml`.
- Avoid embedding one agent's assumptions into the artifact model.
- Capture exit status and useful execution metadata.
- Do not silently grant permissions or bypass agent safety controls.

### Possible configuration

```toml
default-agent = "codex"

[agents.codex]
command = ["codex", "exec"]

[agents.claude]
command = ["claude"]
```

### Acceptance criteria

- The prompt-only workflow remains fully supported.
- At least one adapter works end to end.
- Agent command failures leave repository state recoverable.
- Configuration errors produce clear diagnostics.

---

## CON-011 — Add workflow diagnostics and recovery

- Status: `candidate`
- Priority: `medium`
- Depends on: CON-005, CON-009
- Capability impact: adds maintenance and recovery support

### Outcome

Add commands that detect drift, incomplete transitions, and malformed artifacts.

### Proposed commands

```text
concoct doctor
concoct abandon
concoct recover
```

### Requirements

`doctor` should detect:

- missing canonical files;
- malformed front matter;
- invalid status transitions;
- orphaned review files;
- archive-reference drift;
- capability references to missing IDs;
- roadmap references to missing archives;
- stale generated prompts;
- inconsistent current-task state.

`abandon` should:

- require explicit confirmation;
- preserve the abandoned task in an archive-like record;
- explain whether roadmap status changes to `planned`, `deferred`, or `cancelled`.

`recover` should:

- reconstruct state from durable artifacts;
- never discard code or planning files silently;
- explain every proposed repair.

### Acceptance criteria

- Diagnostics are read-only by default.
- Repairs require explicit action.
- Recovery preserves project archaeology.
- Errors contain actionable next steps.

---

## CON-012 — Improve project archaeology and reporting

- Status: `candidate`
- Priority: `low`
- Depends on: CON-009
- Capability impact: adds historical reporting

### Outcome

Make the archive useful for understanding how the project evolved without turning Concoct into a project-management platform.

### Possible capabilities

```text
concoct history
concoct show <roadmap-id>
concoct capability <capability-id>
```

### Requirements

- Read from Markdown artifacts rather than a separate database.
- Trace roadmap items to tasks, reviews, archives, and capabilities.
- Keep reports human-readable and script-friendly.
- Avoid introducing dashboards or remote services.

### Acceptance criteria

- A user can trace why a capability exists.
- A user can inspect the delivery history of a roadmap item.
- Historical reporting does not mutate project state.

---

## Recommended implementation order

Implement in this order:

```text
CON-001  Canonical project structure
CON-002  Artifact model
CON-003  Command contract and state machine
CON-004  Capability ledger
CON-005  Go CLI foundation and status
CON-006  Deterministic prompt rendering
CON-007  Active task planning
CON-008  Code and review transitions
CON-009  Archive and capability reconciliation
```

Treat these as later work:

```text
CON-010  Direct agent execution
CON-011  Diagnostics and recovery
CON-012  Archaeology and reporting
```

## First task

The recommended first active task is:

```text
concoct plan CON-001
```

CON-001 should establish the authoritative repository structure before implementation work depends on unstable paths.
