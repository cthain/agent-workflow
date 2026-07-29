---
version: 1
project: concoct
updated: 2026-07-29
---

# Capabilities

## Purpose

This file is the canonical human-readable record of what Concoct can do now. It describes observable behavior evidenced by the repository, not planned behavior from the roadmap.

This initial inventory predates Concoct's normal reviewed-task archive history. Its entries cite repository evidence directly. A historical override archive preserves the stale pre-workflow task and the inventory findings, but does not constitute an approving review or roadmap delivery.

## CAP-001 — Durable file-based workflow contract

- Status: `active`
- Audience: `developers and coding agents`
- Added by: `baseline inventory`
- Archive: `.concoct/archive/2026-07-29-legacy-hitl-restructuring/` — baseline evidence; no approving review
- Documentation: `.codex/skills/concoct/SKILL.md`, `doc/workflow.md`

### Capability

Concoct provides a Markdown-based workflow contract for moving substantial software work through product ownership, task planning, implementation, independent review, and archival. The contract defines canonical roadmap, capability, active-task, review, persona, prompt, and archive artifacts, their ownership, and the expected handoffs between roles.

### User value

Humans and agents can preserve product direction, task state, decisions, review evidence, and accepted outcomes in version-controlled files instead of relying on one tool's conversation history.

### Inputs

A repository whose participants follow `AGENTS.md` and the relevant Concoct artifacts and role guidance.

### Outputs and effects

The workflow produces and maintains human-readable roadmap, task-plan, notes, sequential review, capability, and archive records. The files remain inspectable and usable without a Concoct service or database.

### Limitations

- Workflow transitions are currently carried out by humans or agents following the files; they are not enforced by a state-aware CLI.
- The repository contains no automated validation for artifact metadata or workflow state.

### Verification evidence

- `.codex/skills/concoct/SKILL.md` defines the canonical artifacts, role workflows, state discipline, review outcomes, and archive process.
- `.concoct/personas/product-owner.md`, `task-planner.md`, `developer.md`, `reviewer.md`, and `archivist.md` provide role-specific operating guidance.
- `.concoct/roadmap.md` and `.concoct/current/` demonstrate the living artifact layout in this repository.

### Related capabilities

- `CAP-002` supplies reusable prompts for the workflow transitions.
- `CAP-003` packages the contract for use in another repository.
- `CAP-004` connects several coding-agent tools to the shared contract.

## CAP-002 — Manual role-transition prompts

- Status: `active`
- Audience: `developers and coding agents`
- Added by: `baseline inventory`
- Archive: `.concoct/archive/2026-07-29-legacy-hitl-restructuring/` — baseline evidence; no approving review
- Documentation: `.concoct/prompts/README.md`

### Capability

Concoct provides reusable prompts for roadmap intake and for handoffs from product owner to task planner, task planner to developer, developer to reviewer, reviewer to developer or archivist, blocked reviewer to the responsible role, and archivist back to product owner.

### User value

Users can run the workflow manually with an agent while keeping role boundaries, required inputs, allowed mutations, completion evidence, and next actions explicit.

### Inputs

The repository's current workflow artifacts and the prompt matching the desired transition.

### Outputs and effects

Each prompt instructs an agent which persona and artifacts to read, which artifacts the selected role may update, what outcome to produce, and which transition should follow.

### Limitations

- Prompts must currently be selected and supplied to an agent manually.
- No command currently renders prompts from detected repository state.

### Verification evidence

- `.concoct/prompts/roadmap/human-roadmap-input.md`
- `.concoct/prompts/handoffs/`

### Related capabilities

- `CAP-001` defines the artifact and role contract coordinated by these prompts.

## CAP-003 — Reusable project workflow template

- Status: `limited`
- Audience: `project maintainers`
- Added by: `baseline inventory`
- Archive: `.concoct/archive/2026-07-29-legacy-hitl-restructuring/` — baseline evidence; no approving review
- Documentation: `README.md`, `templates/AGENTS.md`

### Capability

Concoct supplies a reusable filesystem template for equipping another repository with canonical project instructions, Concoct workflow state, roadmap and capability schemas, role personas, transition prompts, and coding-agent adapters.

### User value

Project maintainers can reuse a consistent, agent-neutral workflow contract rather than assembling planning, review, and archival guidance from scratch.

### Inputs

The contents of `templates/` and project-specific edits to the installed placeholders and guidance.

### Outputs and effects

The template defines conventional root files and tool adapters alongside Concoct-owned material under `.concoct/`, including `current/`, `archive/`, personas, prompts, a roadmap, and a capability ledger.

### Limitations

- The current bootstrap executable cannot install the template from its repository location because it resolves `templates/` and `personas/` relative to `cmd/concoct/`, where those directories do not exist.
- Several template references use older persona names that do not match the persona files currently shipped.
- The API, code, and user writer persona files are empty.
- Installed templates require project-specific customization before they constitute finished project guidance.

### Verification evidence

- `templates/` contains the root adapters, `.concoct/` artifact hierarchy, persona files, prompts, and placeholder current-task files.
- `bash -n cmd/concoct/concoct` passes.
- An end-to-end invocation of `cmd/concoct/concoct` fails before project creation with `templates directory not found at: .../cmd/concoct/templates`.

### Related capabilities

- `CAP-001` is the workflow contract represented by the template.
- `CAP-004` describes the tool adapters included in the template.

## CAP-004 — Agent-neutral tool adapters

- Status: `active`
- Audience: `developers using coding agents`
- Added by: `baseline inventory`
- Archive: `.concoct/archive/2026-07-29-legacy-hitl-restructuring/` — baseline evidence; no approving review
- Documentation: `doc/multi-agent-workflow.md`

### Capability

Concoct provides thin adapters that direct Codex, Claude Code, GitHub Copilot, Aider, and tools that read a generic conventions file toward the same repository-owned instructions and active task context.

### User value

Teams can use the file-based workflow with multiple coding-agent tools without maintaining a separate durable rule set for each tool.

### Inputs

The adapter appropriate to the user's tool and the canonical files installed in the project.

### Outputs and effects

The adapters point tools to `AGENTS.md`, `.concoct/current/task-plan.md`, `.concoct/current/notes.md`, and role guidance where supported.

### Limitations

- Adapters provide instructions only; they do not launch agents or enforce workflow transitions.
- Some GitHub prompt templates refer to older persona filenames and need manual correction before those particular prompts can be used as written.

### Verification evidence

- `templates/.codex/skills/concoct/SKILL.md`
- `templates/CLAUDE.md`
- `templates/.github/copilot-instructions.md` and `templates/.github/prompts/`
- `templates/.aider.conf.yml`
- `templates/CONVENTIONS.md`

### Related capabilities

- `CAP-001` provides the canonical workflow and artifact rules referenced by the adapters.
- `CAP-003` distributes the adapters with the project template.

## Known capability gaps

- The initializer does not currently complete an end-to-end project bootstrap from its checked-in location.
- The planned `init`, `status`, `roadmap`, `plan`, `code`, `review`, and `archive` command contracts and state transitions are not implemented as a CLI.
- Prompt rendering, direct agent execution, workflow diagnostics, recovery, history reporting, upgrades, and overlays remain roadmap intent rather than current capabilities.
- Repository documentation and parts of the template retain stale paths and persona names from earlier layouts.
- There are no automated tests in the repository.
