---
version: 1
project: concoct
updated: 2026-07-30
---

# Roadmap

## Purpose

This roadmap defines the planned evolution of Concoct.

Concoct is a lightweight, agent-neutral workflow coordinator that turns rough ideas into implementation-ready work and guides transitions between planning, implementation, review, and archival roles.

Core loop:

```text
human idea → concoct(**product-owner** → roadmap → **task-planner** → task-plan → **developer** → source code → **reviewer** → review → **archivist** → capabilities) → product
```

This file records intended future work.

It is distinct from:

- `.concoct/capabilities.md`, which records what Concoct can do now;
- `.concoct/current/task-plan.md`, which records the currently active implementation task;
- `.concoct/archive/`, which preserves completed task history.

## Roadmap conventions

Each roadmap item has a stable identifier.

Statuses for outstanding work:

- `candidate` — accepted as a possible direction but not yet ordered for implementation;
- `planned` — ready to be turned into an active task plan;
- `active` — currently represented by `.concoct/current/task-plan.md`;
- `blocked` — cannot proceed until a dependency or decision is resolved;
- `deferred` — intentionally postponed;
- `cancelled` — no longer intended and awaiting removal after its rationale and
  identifier reservation are preserved.

Priorities:

- `critical`
- `high`
- `medium`
- `low`

`Depends on` records only unresolved delivery dependencies on other outstanding
roadmap items. `Capability prerequisites` records enduring reliance on accepted
current behavior in `.concoct/capabilities.md`. Satisfied sequencing constraints
and historical provenance do not belong in either field.

Delivered and cancelled items leave the active roadmap after their relationships
are reconciled. Their identifiers remain reserved and must not be reused.

Reserved historical identifiers: `CON-003`, `CON-004`, `CON-005`, `CON-006`,
`CON-015`. Accepted delivery evidence is preserved by the corresponding
capability records and archives; CON-004 was cancelled as redundant and has no
delivery archive.

A roadmap item should describe a coherent outcome. Detailed implementation steps belong in the task plan created by:

```text
concoct plan <roadmap-id>
```

---

## CON-007 — Implement active task planning

- Status: `planned`
- Priority: `high`
- Depends on: None
- Capability prerequisites: CAP-001, CAP-005, CAP-006, CAP-007
- Capability impact: adds roadmap-to-task transition

### Outcome

Implement:

```text
concoct plan <roadmap-id>
```

to turn one roadmap item into an active task-planning session.

### Requirements

- Validate that the roadmap item exists and is eligible for planning.
- Validate that every declared capability prerequisite resolves to accepted
  capability truth and that any documented limitation is compatible with the
  planned outcome.
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
- Depends on: CON-007
- Capability prerequisites: CAP-001, CAP-005, CAP-006, CAP-007
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
- Depends on: CON-008
- Capability prerequisites: CAP-001, CAP-005, CAP-007
- Capability impact: automates accepted-task archival and product-truth reconciliation across Git-backed and non-Git lifecycles

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
- move current task artifacts into a dated archive directory;
- create `summary.md`;
- preserve pending delivery and active task evidence for Git-backed tasks until
  `concoct integrate` completes the accepted integration;
- update the roadmap item to `delivered` and clear `.concoct/current/` only at
  the lifecycle's accepted delivery boundary;
- leave non-Git projects ready for the next task after successful archival.

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
- Git-backed archival ends in `archived`, preserves current task evidence, and
  recommends `concoct integrate` without claiming delivery.
- Non-Git archival marks delivery, clears current artifacts, and returns the
  repository to `ready` after all durable writes succeed.
- Successful Git integration performs final delivery bookkeeping and clears
  current artifacts according to the accepted Git lifecycle in CAP-007.
- Roadmap, capabilities, and archive remain cross-referenced.
- Interrupted archival or integration remains recoverable without premature
  capability or roadmap claims.

---

## CON-010 — Add direct agent execution adapters

- Status: `candidate`
- Priority: `medium`
- Depends on: CON-007, CON-008, CON-009
- Capability prerequisites: CAP-005, CAP-006
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
- Depends on: CON-009
- Capability prerequisites: CAP-001, CAP-005, CAP-007
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
- Capability prerequisites: CAP-001
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

## CON-013 — Add opt-in client project upgrades

- Status: `candidate`
- Priority: `medium`
- Depends on: CON-014
- Capability prerequisites: CAP-003, CAP-005
- Capability impact: adds safe lifecycle upgrades for Concoct-enabled projects

### Outcome

Allow a user to run:

```text
concoct upgrade
```

to deliberately bring an existing Concoct-enabled project onto a supported newer Concoct contract without silently replacing project-owned content or losing workflow history.

### Rationale

Concoct installs durable workflow files into client repositories. As that installed contract evolves, users need a low-friction maintenance path that keeps client projects current while respecting local changes and making every upgrade an explicit choice.

### Requirements

- Make upgrades opt-in and show the proposed target before changing the project.
- Identify the project's installed Concoct version or contract level and report when its origin cannot be determined reliably.
- Preview the affected files, migrations, conflicts, and preserved local changes.
- Distinguish Concoct-owned state from conventional project files and locally customized guidance.
- Never silently overwrite conflicting client changes; stop with actionable choices or preserve them for explicit reconciliation.
- Preserve active task state, roadmap and capability records, reviews, archives, and repository history across supported upgrades.
- Leave the project recoverable when an upgrade cannot complete, and report what changed and what remains unresolved.
- Explain when a source or target version is unsupported and recommend a safe next action.

### Product decisions before planning

- Define how installed projects record their Concoct version or contract level.
- Define the authoritative source of upgrade content and how users select or constrain a target release.
- Define ownership and merge policy for conventional files that Concoct installs but projects are expected to customize, including `AGENTS.md` and tool adapters.
- Decide whether preview is the default behavior or whether execution instead requires an explicit confirmation or apply option.

### Acceptance criteria

- An eligible client project can preview an upgrade without mutation.
- Applying an upgrade requires explicit user intent and reports the selected source and target.
- An unmodified supported installation upgrades to the expected contract while preserving project workflow data.
- Locally modified or ambiguous files are never overwritten without an explicit resolution.
- Failed or unsupported upgrades leave the project usable and provide actionable recovery guidance.
- A completed upgrade reports every changed, preserved, skipped, and conflicted artifact.

---

## CON-014 — Add explicit, optional client overlays

- Status: `candidate`
- Priority: `high`
- Depends on: None
- Capability prerequisites: CAP-003, CAP-004, CAP-006
- Capability impact: adds a supported customization layer for client-specific workflow guidance

### Outcome

Allow a Concoct-enabled project to explicitly opt into project-specific overlays that extend or refine Concoct's reusable instructions, skills, prompts, and personas without turning those customizations into changes to the agent-neutral base templates.

### Rationale

Concoct's shared workflow contract must remain reusable and portable, while client projects need durable guidance for their own domain, organization, and working practices. A first-class overlay boundary lets clients specialize the installed workflow without forking the base contract or making ordinary local edits indistinguishable from supported customization.

### Requirements

- Keep overlays optional; projects without an overlay retain the standard Concoct behavior.
- Require overlays to be explicitly selected or declared rather than inferred from incidental files.
- Support client-specific instructions, skills, and prompts, plus augmentation of base personas.
- Define deterministic composition and precedence so humans and agents can inspect the effective guidance and understand its origin.
- Preserve the agent-neutral base contract and keep overlay content distinct from Concoct-owned templates and workflow state.
- Validate overlay references and incompatible customizations with clear, actionable errors.
- Ensure generated or rendered role guidance consistently includes applicable overlays without requiring a specific agent integration.
- Make the overlay boundary available to lifecycle operations so upgrades can preserve client-owned customization without treating it as an ambiguous edit to the base installation.

### Product decisions before planning

- Decide whether a project may compose multiple overlays or selects exactly one.
- Define whether overlays are project-local only or may also come from reusable external packages, and how their identity is recorded durably.
- Define which customization types may replace base guidance and which may only augment it.
- Define whether overlays apply during initialization, to existing enabled projects, or both.

### Acceptance criteria

- A project can use Concoct with no overlay and receive the standard installed and rendered workflow contract.
- A project can explicitly enable an overlay containing each supported customization type.
- Applicable client instructions, skills, prompts, and persona augmentations appear in the effective workflow guidance with deterministic precedence.
- A user can identify which effective guidance came from the base contract and which came from an overlay.
- Missing, malformed, or incompatible overlays fail clearly without partially changing project workflow state.
- Overlay behavior remains portable across the prompt-only workflow and does not require a direct agent execution adapter.
- Upgrade planning can distinguish and preserve declared overlay content from locally modified base files.

---

## Recommended implementation order

Near-term delivery sequence:

```text
CON-007  Active task planning
CON-008  Code and review transitions
CON-009  Archive and capability reconciliation
```

CON-014 is independently planable but remains a candidate pending its recorded
product decisions. CON-013 follows CON-014. Treat CON-010 through CON-012 as
later work governed by their remaining delivery dependencies and priorities.
