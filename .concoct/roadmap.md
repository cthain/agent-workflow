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

Concoct is strict about the integrity, provenance, and interpretation of
durable evidence. It is configurable about how work enters the system, which
activities are required, who performs them, and how accepted work reaches the
product. Automation should reduce context-creation and handoff effort without
hiding the applicable prompts, personas, evidence, decisions, or transitions.

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

- Status: `delivered`
- Archive: `.concoct/archive/2026-07-30-CON-007-active-task-planning/` — accepted and archived on the task branch; delivery pending `concoct integrate`
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
- Depends on: CON-007, CON-008, CON-009, CON-018
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
- Depends on: CON-017
- Capability prerequisites: CAP-003, CAP-004, CAP-006
- Capability impact: adds a supported customization layer for client-specific workflow guidance

### Outcome

Allow a Concoct-enabled project to explicitly opt into project-specific overlays that extend or refine Concoct's reusable instructions, skills, prompts, and personas without turning those customizations into changes to the agent-neutral base templates.

Overlays extend the project-guidance and workflow-policy layers established by
CON-017. They are not the boundary between Concoct protocol and project-owned
truth, and they cannot weaken protocol invariants.

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

## CON-016 — Adopt an existing repository

- Status: `candidate`
- Priority: `high`
- Depends on: CON-017, CON-018
- Capability prerequisites: CAP-001, CAP-003, CAP-005
- Capability impact: adds safe Concoct onboarding for brownfield repositories

### Outcome

Inspect an existing repository and produce an auditable, versionable adoption
proposal before Concoct creates or changes files. Approved application must
preserve repository-owned truth and validate that the inspected state has not
materially drifted.

### Requirements

- Follow an explicit `inspect → report → configure → approve → apply → validate` lifecycle.
- Discover repository boundaries, instructions, product and architecture documentation, verification commands, compatibility constraints, delivery practices, and existing planning sources.
- Report proposed artifacts, preserved files, conflicts, uncertainty, and questions requiring human resolution.
- Treat existing `AGENTS.md` and equivalent guidance as project-owned.
- Never manufacture historical archives or silently promote discovered claims into accepted capabilities.
- Keep inspection non-mutating, make repeated inspection safe, and detect drift before apply.

### Acceptance criteria

- A mature repository can receive a non-mutating adoption report without first becoming Concoct-shaped.
- Every proposed create, modify, reference, and preserve action is inspectable before approval.
- Apply requires explicit approval, refuses stale proposals safely, and leaves project-owned truth intact.
- Validation explains the resulting effective workflow and any unresolved uncertainty.

---

## CON-017 — Separate protocol, policy, and project guidance

- Status: `candidate`
- Priority: `high`
- Depends on: None
- Capability prerequisites: CAP-001, CAP-003, CAP-004, CAP-006
- Capability impact: separates Concoct invariants from configurable workflow policy and repository-owned conventions

### Outcome

Define explicit ownership and composition boundaries for Concoct protocol,
project-selected workflow policy, and repository-owned project guidance while
retaining `AGENTS.md` as a usable human- and agent-facing entry point.

### Requirements

- Make evidence integrity, immutable completed reviews, and invalid-state handling Concoct-owned protocol.
- Make required phases, controls, and Git strategy project-selected policy.
- Keep naming, architecture, coding standards, and verification commands project-owned.
- Permit Concoct-owned material to be upgraded without silently replacing project guidance.
- Detect conflicts deterministically and render the effective instruction set with source attribution.
- Allow project guidance to strengthen policy without weakening protocol invariants.

### Acceptance criteria

- Every effective instruction has an identifiable ownership layer and source.
- Project-owned guidance survives initialization, composition, and supported upgrades unchanged unless explicitly reconciled.
- Conflicting or invariant-weakening instructions fail with actionable diagnostics.
- Existing default workflow behavior remains expressible through the layered model.

---

## CON-018 — Configure workflow policy

- Status: `candidate`
- Priority: `high`
- Depends on: CON-017
- Capability prerequisites: CAP-001, CAP-005, CAP-006, CAP-007
- Capability impact: makes lifecycle requirements explicit and configurable

### Outcome

Allow a repository to select a small typed policy for required, conditional,
externally satisfied, unsupported, or inapplicable workflow activities without
turning Concoct into an arbitrary workflow-graph engine.

### Requirements

- Resolve each governed activity to visible evidence such as `completed`, `not-required`, `not-applicable`, `externally-satisfied`, or `blocked`.
- Require a reason when policy permits a phase or control to be skipped.
- Keep contradictory evidence invalid, completed reviews immutable, capability impact resolved before acceptance, and archival factual.
- Generate handoffs and transition recommendations from effective policy and actual repository state.
- Preserve the current happy path as the supported default profile.

### Acceptance criteria

- Two repositories can select different supported lifecycle policies without changing Concoct protocol.
- Status and rendered prompts expose the resolved requirement and disposition of every governed activity.
- Invalid or invariant-weakening policy is rejected deterministically.
- Absence of an artifact is never silently interpreted as an authorized skip.

---

## CON-019 — Support multiple task origins

- Status: `candidate`
- Priority: `high`
- Depends on: CON-018
- Capability prerequisites: CAP-001, CAP-005
- Capability impact: allows repository work that does not originate in the product roadmap

### Outcome

Give every task explicit provenance while treating roadmap work as one origin
alongside issues, incidents, maintenance, security, dependency changes,
investigations, experiments, review findings, and external changes.

### Requirements

- Record a typed origin and an optional external reference in durable task metadata.
- Preserve `concoct plan <roadmap-id>` as the roadmap-origin path.
- Support creation of non-roadmap tasks without manufacturing strategy entries.
- Classify capability impact as add, change, remove, none, or unknown during planning.
- Require unknown capability impact to be resolved before acceptance.

### Acceptance criteria

- Roadmap and non-roadmap tasks enter the same evidence model with inspectable provenance.
- Origin-specific validation fails clearly without imposing irrelevant roadmap requirements.
- Status, prompts, archives, and reports retain the original provenance.
- Accepted non-roadmap work reconciles capability truth when affected.

---

## CON-020 — Make Git lifecycle strategy-selectable

- Status: `candidate`
- Priority: `high`
- Depends on: CON-018
- Capability prerequisites: CAP-007
- Capability impact: generalizes Git integration while preserving task-branch behavior as the default managed strategy

### Outcome

Expose the accepted task-branch lifecycle through a stable Git strategy boundary
and support current-branch, externally managed, and non-Git lifecycles with
strategy-appropriate evidence and recovery.

### Requirements

- Define where work is isolated, which repository state is authoritative, who integrates, and what proves completion for every strategy.
- Retain CAP-007 task-branch isolation and squash integration as the default managed strategy.
- Let external strategies record trustworthy integration evidence without making Concoct perform the integration.
- Define interruption, resume, drift, and reconciliation behavior for each strategy.
- Never claim integration or completion from ambient branch state alone.

### Acceptance criteria

- Each supported strategy has deterministic starting, archival, integration, recovery, and completion states.
- Existing task-branch projects retain their accepted behavior by default.
- External integration can be proven without granting Concoct control of the provider operation.
- Strategy changes cannot reinterpret existing task evidence silently.

---

## CON-021 — Represent provisional product knowledge

- Status: `candidate`
- Priority: `medium`
- Depends on: CON-016
- Capability prerequisites: CAP-001
- Capability impact: distinguishes accepted product truth from repository discoveries under evaluation

### Outcome

Record adoption and investigation findings with status, confidence, and cited
evidence without prematurely promoting them into canonical capabilities.

### Requirements

- Keep `.concoct/capabilities.md` as simple accepted current truth.
- Store proposed, disputed, obsolete, and unknown claims in a distinct discovery artifact.
- Distinguish verified, documented, and inferred confidence and retain evidence references.
- Make promotion to accepted capability truth explicit and reviewable.

### Acceptance criteria

- Repository inspection can preserve useful uncertain findings without changing capability truth.
- Users can trace each discovery claim to evidence and see its acceptance status.
- Disputed or obsolete findings cannot be consumed as accepted capabilities.
- Promotion preserves provenance and requires an explicit acceptance boundary.

---

## CON-022 — Plan from repository evidence

- Status: `candidate`
- Priority: `high`
- Depends on: CON-016, CON-021
- Capability prerequisites: CAP-001, CAP-006
- Capability impact: makes brownfield task planning evidence-aware

### Outcome

Assemble a bounded, task-relevant repository evidence package before planning
work in an existing project, showing what was examined, why it matters, and
what remains uncertain.

### Requirements

- Select evidence according to the requested change and repository domain rather than dumping the repository.
- Include relevant structure, interfaces, compatibility guarantees, tests, CI, migrations, consumers, and documentation.
- Distinguish current inspection from adoption-baseline knowledge that may have gone stale.
- Preserve an inventory of included, excluded, and unresolved evidence for planner and reviewer use.

### Acceptance criteria

- A planner can trace material assumptions to bounded repository evidence.
- The evidence package is inspectable, deterministic for unchanged inputs, and proportionate to the task.
- Stale baseline claims and unresolved conflicts remain visible.
- Historical Concoct reporting remains separately scoped to CON-012.

---

## CON-023 — Support task profiles

- Status: `candidate`
- Priority: `medium`
- Depends on: CON-018, CON-019
- Capability prerequisites: CAP-001
- Capability impact: adds reusable, inspectable policy presets for common kinds of work

### Outcome

Allow projects to name policy selections for feature, maintenance, incident,
experiment, and other recurring work without creating separate hard-coded
workflow engines.

### Requirements

- Resolve a profile into ordinary visible workflow policy for the selected task.
- Keep profile source, selected values, and overrides inspectable.
- Validate profiles against protocol invariants and the project's supported policy schema.
- Allow explicit per-task choices without requiring a profile.

### Acceptance criteria

- Selecting a profile produces the same observable policy as selecting its values directly.
- Prompts, status, and archives identify the selected profile and resolved rules.
- Missing or incompatible profiles fail without partial task creation.
- Profiles cannot hide skipped activities or weaken protocol invariants.

---

## CON-024 — Support concurrent and interrupting work

- Status: `candidate`
- Priority: `medium`
- Depends on: CON-019, CON-020
- Capability prerequisites: CAP-001, CAP-005, CAP-007
- Capability impact: removes the single-active-task repository constraint

### Outcome

Allow multiple durable tasks to coexist and be selected unambiguously across
branches, worktrees, contributors, interruptions, and external reviews.

### Requirements

- Give each task a stable identity and isolated artifact location.
- Resolve the applicable task from explicit selection or trustworthy repository context.
- Prevent branches, worktrees, reviews, and integration evidence from being attributed to the wrong task.
- Support pausing blocked work while maintenance or an incident proceeds.
- Preserve the current single-task model as a compatible default or migration source.

### Acceptance criteria

- Two tasks can coexist without ambiguous ownership of plans, notes, reviews, or Git evidence.
- Task selection and valid transitions are deterministic in supported contexts.
- Interrupting work does not rewrite or invalidate the paused task's evidence.
- Status reports ambiguity instead of guessing when no unique task can be resolved.

---

## CON-025 — Explain effective workflow and state

- Status: `candidate`
- Priority: `medium`
- Depends on: CON-018, CON-020, CON-023
- Capability prerequisites: CAP-005, CAP-006
- Capability impact: makes configured workflow behavior and state interpretation directly inspectable

### Outcome

Explain the effective workflow, task rules, state evidence, valid transitions,
and blocked transitions together with the configuration source of each rule.

### Requirements

- Report active profile, resolved phase requirements, Git strategy, and task origin.
- Cite the evidence used to determine current state.
- Explain valid and blocked transitions and their reasons.
- Attribute effective values to defaults, project policy, profiles, or task-specific choices.
- Keep explanation read-only and script-friendly.

### Acceptance criteria

- A user can determine why Concoct selected a state or refused a transition without reading implementation code.
- Explanations agree with status and prompt selection for the same repository state.
- Defaulted and explicitly configured behavior are distinguishable.
- Invalid evidence is reported rather than normalized away.

---

## CON-026 — Reconcile externally performed work

- Status: `candidate`
- Priority: `medium`
- Depends on: CON-019, CON-020, CON-022
- Capability prerequisites: CAP-001, CAP-005
- Capability impact: establishes trustworthy Concoct state for work begun or completed outside its managed lifecycle

### Outcome

Inspect existing branches, commits, pull requests, emergency fixes, automation,
or externally merged contributions and propose the evidence and remaining work
needed to bring them into Concoct's durable model.

### Requirements

- Inspect and propose before mutating workflow truth.
- Support retrospective task records, proposed capability changes, missing review work, integration evidence, and unresolved provenance questions.
- Refuse acceptance when trustworthy state cannot be established.
- Distinguish reconciliation from adoption of the repository itself.
- Preserve external identifiers and evidence sources.

### Acceptance criteria

- Existing work can enter Concoct without being falsely represented as having followed the managed lifecycle.
- The proposal identifies verified evidence, gaps, required decisions, and the path to acceptance.
- No external work is marked accepted solely because a branch or commit exists.
- Applied reconciliation remains traceable through task, review, archive, capability, and integration records as applicable.

---

## CON-027 — Track and resolve product bugs

- Status: `candidate`
- Priority: `high`
- Depends on: CON-009, CON-019
- Capability prerequisites: CAP-001, CAP-005, CAP-006
- Capability impact: adds a durable bug register and bug-origin task lifecycle without treating defects as product-roadmap work

### Outcome

Allow projects to record, triage, resolve, verify, and close observed
contradictions in accepted or intended product behavior while keeping defects
distinct from roadmap direction, capability truth, and active-task review
findings.

A defect discovered before task acceptance remains an obligation of the active
task. A defect discovered after acceptance becomes a separately tracked bug.

### Rationale

Roadmap items describe intended product evolution, so using them as a defect
register obscures both strategy and current product truth. A dedicated bug
lifecycle preserves durable evidence of observed failures and their disposition
while allowing remediation to use the ordinary reviewed task lifecycle with
explicit bug provenance.

### Requirements

- Introduce `.concoct/bugs.md` as the authoritative human-readable register for
  independently tracked product defects.
- Give each bug a stable unique identifier, lifecycle status, severity,
  expected and observed behavior, evidence, affected capabilities or intended
  behavior, and durable resolution references.
- Support the primary lifecycle `reported → confirmed → planned → in-progress
  → resolved → verified → closed`, plus reasoned alternate dispositions for
  duplicate, not-a-bug, cannot-reproduce, deferred, and superseded reports.
- Keep pre-acceptance defects and review findings within their active task;
  create a separate bug only after acceptance or when the defect is otherwise
  independent of active-task obligations.
- Allow a confirmed bug to originate or become associated with an ordinary
  CON-019 task without manufacturing a roadmap item, and maintain deterministic
  bidirectional references between the bug and repair task.
- Carry repair work through the configured planning, implementation, review,
  archive, and delivery lifecycle while retaining links among the bug, task,
  reviews, archive, capability reconciliation, and delivery evidence.
- Classify each report during triage as a failure to conform to an accepted
  capability, an incorrect or incomplete capability contract, intended but not
  yet accepted behavior, or not a product defect.
- Require repair tasks to classify capability impact as none, clarification,
  change, or unknown; permit unknown during triage and planning but resolve it
  before acceptance.
- Do not change capability truth merely because a bug is reported. Restore
  conformance without manufacturing a new capability when the accepted
  contract is already correct; route genuine product evolution through an
  explicit roadmap decision and accepted capability reconciliation.
- Provide supported operations to report, inspect, list and filter, triage,
  reprioritize, associate repair work, record resolution and independent
  verification, close, reopen, and find bugs by affected capability.
- Make status and role prompts expose applicable bug state, invalid evidence,
  and valid next actions without silently synchronizing divergent bug and task
  states or treating implementation, review, archive, or integration as closure.
- Preserve unresolved, deferred, duplicate, rejected, superseded, closed, and
  reopened bug history rather than deleting or rewriting prior evidence.

### State integrity

- Confirmed bugs require durable supporting evidence.
- Planned and in-progress bugs reference a valid repair task, and a bug-origin
  repair task references a valid bug.
- Resolved means implementation claims the defect is corrected; verified
  requires independent evidence against the originally reported behavior.
- Verification requires accepted repair work, and closure requires complete
  resolution, verification, provenance, and capability-impact evidence.
- Duplicate reports identify their canonical bug, and reopened bugs preserve
  earlier resolution and verification history.
- Contradictory or incomplete evidence produces an invalid state rather than an
  inferred or arbitrarily selected result.
- Bug state remains derivable from durable artifacts without conversation
  history.

### Initial validation

- Seed the first accepted post-CON-015 output-path defect as a confirmed bug
  record: ignored and untracked in-repository prompt output should be allowed,
  while tracked or unignored output remains rejected and external output
  remains supported.
- Validate the artifact and state model during CON-027 with representative
  evidence for bidirectional task references, capability-impact resolution,
  invalid states, closure, and reopening; do not fabricate those later states
  on the seeded real bug.
- Do not claim that CON-027 itself completed a bug-origin repair lifecycle. Its
  accepted delivery makes that lifecycle available; the seeded bug's repair
  must then proceed as a separately evidenced bug-origin task.

### Documentation

- Explain the boundaries among bugs, roadmap items, capabilities, tasks, and
  active-task review findings.
- Document lifecycle states and dispositions, triage and capability
  responsibilities, repair-task creation and resumption, the distinction
  between resolution and verification, and preservation of historical evidence.

### Acceptance criteria

- A post-acceptance defect can be recorded and triaged without adding product
  evolution to the roadmap or changing accepted capability truth.
- A pre-acceptance defect remains visibly owned by its active task and cannot be
  closed through the separate bug lifecycle to bypass task review.
- A confirmed bug can originate or associate with an ordinary repair task, and
  both artifacts retain deterministic provenance from report through review,
  archive, delivery, capability reconciliation, and verified closure.
- Bug closure requires observable verification against the reported behavior;
  implementation completion, review approval, or integration alone is not
  treated as verification unless the required evidence is present.
- Capability-conformance repairs do not create spurious capabilities, while
  changed expected behavior cannot be closed as a bug without an explicit
  product decision and accepted capability reconciliation.
- Listing and filtering can identify open bugs and all bugs affecting a
  selected capability without requiring archive or conversation reconstruction.
- Reopening and alternate dispositions preserve reasons, links, and prior
  evidence with clear, valid next actions.
- Malformed, contradictory, missing, or stale bug and task references fail
  clearly without partially mutating bug, task, roadmap, or capability state.

---

## CON-028 — Recommend the next project action

- Status: `delivered`
- Delivery: pending integration; archived evidence at `.concoct/archive/2026-07-30-CON-028-recommend-next-project-action/`
- Priority: `high`
- Depends on: None
- Capability prerequisites: CAP-001, CAP-005, CAP-006, CAP-008
- Capability impact: adds an explicit Product Owner decision step between ready state and selection of the next workflow action

### Outcome

Provide `concoct next` as the single recommended command when a project is
ready. It renders an evidence-backed Product Owner prompt that recommends the
next valid action without selecting work, creating a task, or changing project
lifecycle state.

### Rationale

Ready state currently presents an unexplained choice between roadmap intake and
planning a known item. After initialization or accepted delivery, users need a
clear decision step that considers strategy, eligibility, blockers, and other
supported work origins before directing them into the appropriate existing
workflow.

### Requirements

- Permit `concoct next` only from structurally valid ready state and keep the
  command advisory and read-only.
- Deterministically assemble authoritative roadmap, accepted capability,
  dependency, prerequisite, relevant archive, and supported task-origin
  evidence, then render it with the Product Owner guidance needed to make the
  recommendation.
- Keep semantic Product Owner judgment distinct from CLI validation: the CLI
  must not claim to choose work merely because it assembled the prompt.
- Require the Product Owner result to recommend exactly one of: planning an
  eligible roadmap item; addressing another currently supported task origin;
  refining or reconciling the roadmap; resolving a specific blocker or
  inconsistency; or acknowledging that no actionable work is recorded.
- Explain why the recommended action is next, cite its durable evidence and
  blockers, and provide the exact follow-up command when one exists.
- Do not let priority override unresolved delivery dependencies, unsupported
  capability prerequisites, invalid evidence, or missing product decisions.
- Support ordinary deterministic prompt output to stdout or a create-only
  `--output <path>` destination with the existing output-safety contract.
- Update ready-state status, successful integration output, documentation,
  templates, help, skills, and applicable personas to recommend `concoct next`
  consistently instead of presenting an unexplained roadmap-or-plan choice.
- Preserve the boundaries among `status` for lifecycle state, `next` for work
  recommendation, `roadmap` for product-direction changes, and `plan` for an
  already selected eligible work origin.
- Allow future task-origin types to participate when their accepted contracts
  become available without making them delivery dependencies of this item.

### Safety constraints

- Invoking the command does not create or activate a task, change project
  phase, or mutate roadmap, capability, bug, current-task, or archive artifacts.
- Structurally missing or contradictory state remains invalid and is not
  normalized into a recommendation prompt.
- The prompt may recommend reconciliation but does not perform it as a side
  effect of deciding what should happen next.
- Implementation existence and passing checks are not treated as acceptance or
  eligibility evidence.

### Acceptance criteria

- Ready-state status and successful integration recommend exactly `concoct
  next`, giving a new or returning user one unambiguous next command.
- `concoct next` renders a complete Product Owner prompt from authoritative
  repository evidence without mutating workflow state or selecting work.
- Prompt output is byte-deterministic for unchanged evidence and behaves
  consistently on stdout and through `--output <path>`.
- The Product Owner guidance covers an eligible roadmap item, a supported
  non-roadmap origin, roadmap reconciliation, a specific blocker, and no
  actionable recorded work as distinct recommendation outcomes.
- Recommendations distinguish structural eligibility from product judgment,
  explain their evidence, and never present blocked work as immediately
  plannable.
- Unsupported future origins are not implied to be available, while accepted
  origin types can be incorporated without changing the command boundary.
- Documentation, templates, help, personas, and skill guidance consistently
  explain the roles of `status`, `next`, `roadmap`, and `plan`.
- Repository-defined tests cover command-state validation, evidence assembly,
  every recommendation outcome, non-mutation, output safety, and ready-state
  guidance.

---

## Recommended implementation order

Near-term delivery sequence:

```text
CON-007  Active task planning
CON-028  Next-action recommendation
CON-008  Code and review transitions
CON-009  Archive and capability reconciliation
```

After the initial CLI lifecycle is complete, prioritize the flexibility work in
three increments:

```text
Foundation:   CON-017 → CON-018 → CON-016 → CON-021 → CON-022
Everyday use: CON-019 → CON-027 → CON-020 → CON-023 → CON-025 → CON-026
Scale:        CON-024
```

CON-014 now builds on CON-017 and remains a candidate pending its recorded
product decisions; CON-013 follows CON-014. CON-010 remains behind workflow
policy so direct execution does not automate an accidentally universal
lifecycle. CON-011 and CON-012 remain useful later work, with task-scoped
repository archaeology assigned to CON-022 and historical reporting retained
in CON-012. CON-027 follows typed task origins but does not wait for Git
strategy selection because its register and provenance rules apply across
delivery strategies.
