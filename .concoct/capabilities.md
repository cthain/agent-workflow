---
version: 1
project: concoct
updated: 2026-07-30
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
- Updated by: `.concoct/archive/2026-07-29-CON-003-command-contract-state-machine/`
- Documentation: `.codex/skills/concoct/SKILL.md`, `doc/workflow.md`

### Capability

Concoct provides a Markdown-based workflow contract for moving substantial software work through product ownership, task planning, implementation, independent review, and archival. The contract defines canonical roadmap, capability, active-task, review, persona, prompt, and archive artifacts, their ownership, and the expected handoffs between roles. It also provides a normative reference for the initial command surface and an artifact-backed state machine covering valid transitions, remediation, blocked-review recovery, invalid evidence, and transactional archival.

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
- `doc/command-reference.md` defines the complete normative contract for the seven initial commands.
- `doc/state-machine.md` defines workflow state from observable artifacts and specifies transitions, review recovery, invalid states, and archive atomicity.
- `.concoct/archive/2026-07-29-CON-003-command-contract-state-machine/review-02.md` records approval of the command and state-machine contract.

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

- Manual prompt use still requires a human or agent to select and supply the
  appropriate asset; CAP-006 provides executable selection and rendering for
  the initial roadmap, planning, development, and review roles.

### Verification evidence

- `.concoct/prompts/roadmap/human-roadmap-input.md`
- `.concoct/prompts/handoffs/`

### Related capabilities

- `CAP-001` defines the artifact and role contract coordinated by these prompts.
- `CAP-006` renders selected manual prompt assets with validated workflow context.

## CAP-003 — Reusable project workflow template

- Status: `limited`
- Audience: `project maintainers`
- Added by: `baseline inventory`
- Archive: `.concoct/archive/2026-07-29-legacy-hitl-restructuring/` — baseline evidence; no approving review
- Updated by: `.concoct/archive/2026-07-29-CON-005-go-cli-foundation/`
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

- Several template references use older persona names that do not match the persona files currently shipped.
- The API, code, and user writer persona files are empty.
- Installed templates require project-specific customization before they constitute finished project guidance.

### Verification evidence

- `templates/` contains the root adapters, `.concoct/` artifact hierarchy, persona files, prompts, and placeholder current-task files.
- `internal/project/project_test.go` verifies complete embedded template copying and real-Git initialization behavior.
- `.concoct/archive/2026-07-29-CON-005-go-cli-foundation/review-03.md` records approval of caller-directory-independent initialization with root files, dotfiles, nested content, personas, prompts, staged files, and no generated commit.

### Related capabilities

- `CAP-001` is the workflow contract represented by the template.
- `CAP-004` describes the tool adapters included in the template.
- `CAP-005` initializes projects from the embedded template and reports their workflow state.

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

## CAP-005 — Executable CLI initialization and workflow status

- Status: `active`
- Audience: `project maintainers, developers, and coding agents`
- Added by: `.concoct/archive/2026-07-29-CON-005-go-cli-foundation/`
- Documentation: `README.md`, `doc/command-reference.md`, `doc/state-machine.md`

### Capability

Concoct provides a Go CLI with `init` and read-only `status` commands. It can create a Concoct-enabled Git repository from the complete embedded project template and derive deterministic workflow state from canonical repository artifacts.

### User value

Project maintainers can bootstrap the workflow reliably from an installed binary, and humans or agents can inspect the current task, review outcome, capability impact, diagnostics, and recommended next action without manually interpreting artifact combinations.

### Inputs

- `concoct init <project>` accepts one new project target whose parent already exists.
- `concoct status` runs from a Concoct project root or any nested directory.

### Outputs and effects

- `init` copies root files, dotfiles, nested templates, personas, and prompts; writes bootstrap guidance; initializes Git; stages generated files; and creates no commit.
- `status` discovers the project, validates roadmap, capability, task, notes, review, remediation, and blocked-review evidence, then reports the applicable state and next action without modifying the repository.
- Malformed, incomplete, contradictory, or representative interrupted-archive evidence is reported as `invalid` with actionable diagnostics.

### Limitations

- Role-completion mutations and archival remain later roadmap work; CAP-006
  adds state-preserving prompt rendering to the executable surface.
- Remediation disposition validation is textual because the notes schema does not define structured review-finding identifiers.
- Metadata parsing intentionally targets the checked-in Concoct schemas rather than arbitrary Markdown documents.

### Verification evidence

- `internal/project/project_test.go` covers discovery, template copying, Git behavior, initialization safety, and read-only inspection.
- `internal/workflow/workflow_test.go` covers normative states, metadata validation, sequential reviews, remediation, blocked-review recovery, composed recovery precedence, and invalid evidence.
- `.concoct/archive/2026-07-29-CON-005-go-cli-foundation/review-03.md` records the approving independent review and end-to-end verification.

### Related capabilities

- `CAP-001` defines the workflow and state contract implemented by status detection.
- `CAP-003` supplies the embedded project template used by initialization.
- `CAP-004` supplies the agent adapters installed with that template.
- `CAP-006` adds deterministic role-prompt rendering to the CLI.

## CAP-006 — Deterministic role-aware prompt rendering

- Status: `active`
- Audience: `developers and coding agents`
- Added by: `.concoct/archive/2026-07-30-CON-006-deterministic-prompt-rendering/`
- Documentation: `README.md`, `doc/command-reference.md`, `doc/state-machine.md`

### Capability

Concoct can render complete, deterministic prompts for Product Owner roadmap
intake, task planning, development, and independent review from validated
repository state. It selects the applicable persona and workflow mode,
including implementation continuation, changes-requested remediation, and
blocked-review recovery routes.

### User value

Humans and coding agents can obtain the correct inspectable role handoff
without manually interpreting workflow state or duplicating durable role
rules.

### Inputs

- `concoct roadmap` renders Product Owner roadmap guidance from the ready state.
- `concoct plan <roadmap-id>` validates an eligible item and satisfied
  dependencies before rendering Task Planner guidance.
- `concoct code` renders the applicable initial, continuing, remediation, or
  blocked-recovery Developer guidance.
- `concoct review` renders Reviewer guidance with prior reviews and the next
  sequential review path.
- Each command accepts optional create-only `--output <path>` output.

### Outputs and effects

- Commands write deterministic prompt bytes to stdout by default or create a
  new output file containing identical bytes.
- Rendered prompts identify exact inputs, authorized updates, detected state
  and mode, expected outcome, validation requirements, and next transition.
- Rendering validates command eligibility and workflow evidence but does not
  mutate workflow state, persist role outcomes, or launch an agent.

### Limitations

- Prompt commands provide guidance only; later roadmap work owns task/review
  mutations, archival automation, and direct agent execution.
- Output files are create-only and existing destinations are never overwritten.
- Archive-summary relevance is selected conservatively from identifiers in
  validated task and command context because no archive index exists.

### Verification evidence

- `internal/prompt/render_test.go` and its nine golden fixtures cover all four
  commands and the materially distinct development and review modes.
- `internal/cli/cli_test.go` covers stdout/file byte equality, nested project
  discovery, workflow non-mutation, argument errors, and collision refusal.
- `.concoct/archive/2026-07-30-CON-006-deterministic-prompt-rendering/review-02.md`
  records approval after full-output golden coverage was added.

### Related capabilities

- `CAP-001` defines the workflow and state contract used for eligibility.
- `CAP-002` supplies the canonical manual prompt assets appended to rendering.
- `CAP-005` supplies CLI project discovery and validated state detection.

## Known capability gaps

- The `roadmap`, `plan`, `code`, and `review` commands render prompts but do not
  perform role-completion state mutations; `archive` is not implemented as a
  CLI command.
- Direct agent execution, workflow diagnostics, recovery, history reporting,
  upgrades, and overlays remain roadmap intent rather than current capabilities.
- Repository documentation and parts of the template retain stale paths and persona names from earlier layouts.
