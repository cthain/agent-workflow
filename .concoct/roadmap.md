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

## CON-003 — Define the command contract and workflow state machine

- Status: `delivered`
- Priority: `critical`
- Depends on: None
- Capability impact: defines the intended CLI behavior
- Archive: `.concoct/archive/2026-07-29-CON-003-command-contract-state-machine/`

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
doc/command-reference.md
doc/state-machine.md
```

### Acceptance criteria

- Every initial command has a complete contract.
- Invalid transitions have clear, actionable errors.
- `handoff` is not required in the normal happy path.
- Each role command performs its own incoming and outgoing handoff.
- The model supports implementation-remediation loops after review.

---

## CON-004 — Establish `capabilities.md` as current accepted product truth

- Status: `cancelled`
- Priority: `high`
- Depends on: None
- Capability impact: none; the accepted baseline inventory already established the capability ledger

### Resolution

Cancelled as redundant on 2026-07-29. The pre-workflow baseline inventory already established `.concoct/capabilities.md`, stable `CAP-*` identifiers, capability-aware planning guidance, and archive reconciliation rules. The historical override archive records the inventory evidence without representing an approving review, while archived CON-003 demonstrates reviewed reconciliation of `CAP-001`.

This item is intentionally not marked `delivered`: no approved CON-004 task archive exists. Creating implementation work solely to manufacture that history would add no product outcome. Existing roadmap work may rely directly on the accepted capability truth rather than on delivery of CON-004.

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

- Status: `delivered`
- Priority: `critical`
- Depends on: CON-003
- Capability impact: adds the executable CLI foundation
- Delivered: `.concoct/archive/2026-07-29-CON-005-go-cli-foundation/`
- Capability: CAP-005

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

- Status: `delivered`
- Priority: `high`
- Depends on: CON-003, CON-005
- Capability impact: adds role-aware prompt generation
- Delivered: `.concoct/archive/2026-07-30-CON-006-deterministic-prompt-rendering/`
- Capability: CAP-006

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
- `.concoct/personas/product-owner.md`
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
- Depends on: CON-005, CON-006, CON-015
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
- Depends on: CON-005, CON-008, CON-015
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
  current artifacts according to the accepted CON-015 lifecycle.
- Roadmap, capabilities, and archive remain cross-referenced.
- Interrupted archival or integration remains recoverable without premature
  capability or roadmap claims.

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

## CON-013 — Add opt-in client project upgrades

- Status: `candidate`
- Priority: `medium`
- Depends on: CON-005, CON-014
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
- Depends on: CON-005, CON-006
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

## CON-015 — Isolate and integrate tasks with Git branches

- Status: `delivered`
- Priority: `high`
- Depends on: CON-003, CON-005
- Capability impact: adds a Git-backed task branch lifecycle to planning and archival
- Delivered: `.concoct/archive/2026-07-30-CON-015-isolate-integrate-git-tasks/`
- Capability: CAP-007

### Outcome

For a Concoct-enabled Git repository, keep roadmap intake and selection on the current trunk while carrying each planned task through implementation and review on a dedicated feature branch, then integrate approved and archived work back into the trunk from which the task began.

### Rationale

Concoct currently coordinates artifact ownership and role transitions without isolating task changes from the repository's current branch. A durable task-to-branch association protects the trunk as the home of the roadmap, gives implementation and review a bounded change set, and makes integration an explicit acceptance step rather than an incidental user action.

### Requirements

- Apply the branch lifecycle only when the project is inside a Git repository; define and preserve appropriate behavior for non-Git projects.
- Start planning from the user's current branch and preserve that branch as the task's integration trunk.
- Give each task branch a deterministic name derived from the roadmap identifier and the roadmap item's short name.
- Create and switch to the task branch as part of establishing the active task, before task implementation begins.
- Preserve durable evidence linking the active task, its task branch, and its integration trunk so later roles do not infer them from ambient checkout state.
- Require development, remediation, and review for the task to remain associated with the task branch.
- After approval, have the Archivist update and validate archival and capability records, plus the pending roadmap reconciliation, on the task branch before integrating the accepted task into its recorded trunk; mark the roadmap item delivered only after integration succeeds.
- Return to `ready` only when archival reconciliation and integration into the recorded trunk have both succeeded.
- Detect unsafe starting states, branch-name collisions, checkout drift, integration conflicts, and partial completion without overwriting work or losing workflow evidence.
- Keep the workflow usable with local Git repositories and avoid requiring a hosting provider, pull request, or remote branch.

### Product decisions

- Concoct creates a commit for each role transition that produces task changes. The Archivist owns the archival commit before integration.
- Accepted task branches are squash-merged into their recorded trunk and deleted after successful integration.
- The branch checked out when planning begins is the task's recorded integration trunk. Concoct does not infer or require a conventional primary branch.
- A clean worktree means that `git status --short` produces no entries at workflow input and output boundaries. Local commits that have not been pushed do not make the worktree dirty.
- When the recorded trunk has a matching upstream branch, Concoct prompts before pushing by default; a project may explicitly enable automatic pushes. Without a matching upstream, Concoct works locally and does not attempt to push.
- Any conflict reported by Git during merge or rebase requires human-owned integration recovery; Concoct does not attempt to resolve conflicted paths autonomously. The handoff must preserve the task and archival evidence, explain the conflict, and recommend a resolution.
- Non-Git projects continue through the existing unbranched local workflow; Git is not a prerequisite for using Concoct.
- Git archival introduces an `integrating` state when integration requires human intervention and an `integrated` state after integration succeeds but before the workflow returns to `ready`.

### Remaining product decisions before planning

None.

### Integration recovery

- When Git reports a conflict during squash integration, Concoct enters the `integrating` state, leaves the recorded trunk checked out, and preserves the task branch, archival commit, pre-integration trunk commit, and recovery metadata.
- The human resolves conflicted paths and stages the intended integration result but does not create the squash commit manually.
- `concoct integrate --continue` resumes the workflow. It validates the recorded trunk, task branch and archival commit, pre-integration trunk head, resolved index, staged changes, expected archival records, and absence of unrelated repository operations before creating the squash integration commit.
- Invoking `concoct integrate --continue` is the human attestation that the staged semantic conflict resolution is correct. Concoct validates repository and workflow structure but does not independently determine whether human conflict choices are semantically correct.
- After creating the squash commit, Concoct enters `integrated`, records the integration commit in durable archival metadata through a final bookkeeping commit, deletes the accepted task branch, resets active workflow artifacts, and returns to `ready`.
- `concoct integrate --continue` is idempotent and may resume either an incomplete `integrating → integrated` transition or an incomplete `integrated → ready` transition.
- `concoct integrate --abort` abandons the current integration attempt, restores the recorded trunk to its pre-integration state, preserves and checks out the approved task branch, and returns the workflow to its pre-integration archived state.
- Concoct stores local recovery metadata under `.git/concoct/integrations/<task-id>.md` while the operation is incomplete. The committed archive remains the authoritative historical record after integration succeeds.

### Archived state and integration entry point

- `concoct archive` ends in the observable `archived` state. In this state, the latest review is approved, the accepted task artifacts and archive summary exist on the task branch, capability and roadmap reconciliation has been committed, the archival commit is recorded, and integration into the recorded trunk remains pending.
- An archived task is not yet delivered and the repository is not yet `ready`.
- `concoct integrate` starts or retries squash integration of the recorded archival commit into the recorded integration trunk.
- `concoct integrate` is valid only from `archived`. It validates the recorded trunk, task branch, task base, archival commit, clean worktree, and absence of another active Git operation before beginning integration.
- A successful integration advances through `integrated` and then completes final bookkeeping, branch deletion, active-state cleanup, and transition to `ready`.
- A Git conflict leaves the workflow in `integrating`.
- `concoct integrate --continue` resumes an active integration after the human resolves and stages conflicted paths.
- `concoct integrate --abort` abandons the active integration attempt, restores the recorded trunk to its pre-integration state, preserves the task branch and archival commit, and returns the workflow to `archived`.
- After an abort, a new integration attempt is started with `concoct integrate`, not `concoct integrate --continue`.
- The roadmap item is marked `delivered` only after integration succeeds and final workflow cleanup completes.

### Acceptance criteria

- Planning an eligible roadmap item in a Git repository establishes and checks out a deterministically named task branch and records its source trunk.
- Code and review transitions reject or clearly diagnose operation from a checkout that does not match the active task branch.
- Approved archival reconciles durable product records before integrating the task branch into its recorded trunk.
- Successful integration leaves the recorded trunk checked out, containing the accepted implementation and archival records, with the roadmap item delivered and the repository in `ready` state.
- An integration conflict leaves the workflow in `integrating` with the task branch and archival evidence preserved until human resolution can be validated.
- A collision, dirty worktree, checkout mismatch, merge conflict, or interrupted integration preserves recoverable task and repository evidence and reports an actionable next step.
- The feature does not require a remote repository or a provider-specific review mechanism.
- Non-Git project behavior is explicit and does not fail merely because Git branch isolation is unavailable.

---

## Recommended implementation order

Implement in this order:

```text
CON-005  Go CLI foundation and status
CON-006  Deterministic prompt rendering
CON-015  Git-backed task isolation and integration
CON-007  Active task planning
CON-008  Code and review transitions
CON-009  Archive and capability reconciliation
```

Treat these as later work:

```text
CON-010  Direct agent execution
CON-011  Diagnostics and recovery
CON-012  Archaeology and reporting
CON-013  Client project upgrades
CON-014  Client overlays
```
