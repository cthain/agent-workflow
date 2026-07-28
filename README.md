# Concoct

Turn raw ideas into work agents can execute.

Concoct is a lightweight, agent-neutral workflow for turning ideas into durable task plans that coding agents can inspect, implement, review, and archive.

```text
idea → concoct helps agents cook → eatin’ big time
```

It works with Codex, Claude Code, GitHub Copilot, Aider, and other capable coding agents. Durable project and task context lives in ordinary repository files rather than in one tool's conversation history.

## How the installed workflow works

Generated projects use this stable, agent-neutral contract:

```text
AGENTS.md                         # canonical project instructions
.concoct/personas/                # role-specific working guidance
.concoct/current/task-plan.md      # active task
.concoct/current/notes.md          # durable task context
.concoct/archive/                  # completed task history
```

Optional adapters point back to those canonical files:

```text
CLAUDE.md
CONVENTIONS.md
.aider.conf.yml
.github/copilot-instructions.md
.github/prompts/
.codex/skills/concoct/SKILL.md
```

## Repository layout

Concoct's source files use ordinary root-level directories. The `.concoct/` directory is reserved for generated client installs.

```text
.
├── AGENTS.md
├── README.md
├── LICENSE
├── concoct
├── current/       # Concoct's own active planning files
├── archive/       # Concoct's completed task history
├── docs/          # workflow documentation
├── prompts/       # reusable role prompts
├── personas/      # planner, developer, reviewer, and writer personas
├── skills/        # skills for working on Concoct itself
└── templates/     # files installed into generated projects
```

In a generated project, Concoct-owned task state and personas live under `.concoct/`. Conventional instruction and tool integration files remain at the project root. Removing `.concoct/` therefore removes Concoct's working material without removing the project's conventional files or tool configuration.

## Bootstrap a project

```bash
git clone <repository-url>
cd concoct
./concoct ../my-new-project
```

`concoct` creates the project, copies templates and personas (including dotfiles), initializes Git, stages the generated files, and writes:

```text
.concoct/current/bootstrap-prompt.md
```

It does not create a commit. Review the staged files before committing them.

Next:

1. Open the generated project.
2. Give `.concoct/current/bootstrap-prompt.md` and your idea to Chappie or ChatGPT.
3. Review the generated `README.md`, `AGENTS.md`, task plan, and notes.
4. Ask a coding agent to adopt `.concoct/personas/code-developer.md` and implement the plan.
5. Use the reviewer prompt before archiving completed work.

## Workflow

Use planning files when losing context would be costly: architecture changes, multi-file refactors, features, public API changes, repository setup, or multi-session work. Skip them for tiny edits and one-shot answers.

The reusable prompts are in `prompts/`:

- `create-task-plan.md`
- `continue-task.md`
- `review-task.md`
- `document-task.md`
- `archive-task.md`
- `multi-agent-handoff.md`

See [workflow.md](docs/workflow.md) for the full loop and [multi-agent-workflow.md](docs/multi-agent-workflow.md) for role coordination.

## Naming conventions

- Use hyphens, not underscores, in file and directory names.
- Keep conventional, long-lived project artifacts such as `AGENTS.md`, `README.md`, `CHANGELOG.md`, and `CONTRIBUTING.md` at the root.
- Use uppercase Markdown filenames for long-lived project-level artifacts.
- Use lowercase hyphenated Markdown filenames for task and workflow artifacts.

## Manual repository rename

The local content is branded for the intended repository name `concoct`. A repository owner must still rename the GitHub repository, update its description and topics, confirm clone URL redirects, and update local remotes where needed.
