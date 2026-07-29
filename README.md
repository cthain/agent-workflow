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

Concoct's source files use ordinary root-level directories. The `.concoct/`
directory contains this repository's own workflow artifacts and has the same
role in generated client projects.

```text
.
├── AGENTS.md
├── README.md
├── LICENSE
├── cmd/concoct/   # Go command and source-checkout wrapper
├── .concoct/      # Concoct's own workflow artifacts
├── doc/           # workflow and command documentation
└── templates/     # files installed into generated projects
```

In a generated project, Concoct-owned task state and personas live under `.concoct/`. Conventional instruction and tool integration files remain at the project root. Removing `.concoct/` therefore removes Concoct's working material without removing the project's conventional files or tool configuration.

## Bootstrap a project

Build or install the Go CLI, then initialize a project from any working directory:

```bash
go build -o ./bin/concoct ./cmd/concoct
./bin/concoct init ../my-new-project
```

From a source checkout, `./cmd/concoct/concoct init ../my-new-project` is a thin
compatibility wrapper around the same Go implementation.

`concoct init` creates the project, copies its embedded templates and personas
(including dotfiles), initializes Git, stages the generated files, and writes:

```text
.concoct/current/bootstrap-prompt.md
```

It does not create a commit. Review the staged files before committing them.

Inspect workflow state from the project root or any nested directory:

```bash
concoct status
```

Next:

1. Open the generated project.
2. Follow `.concoct/current/bootstrap-prompt.md` to start Product Owner intake.
3. Use the handoff prompts under `.concoct/prompts/` for manual role transitions.

## Workflow

Use planning files when losing context would be costly: architecture changes, multi-file refactors, features, public API changes, repository setup, or multi-session work. Skip them for tiny edits and one-shot answers.

The reusable prompts are in `.concoct/prompts/`. See
[workflow.md](doc/workflow.md) for the full loop and
[multi-agent-workflow.md](doc/multi-agent-workflow.md) for role coordination.

## Development

```bash
go test ./...
go vet ./...
go build ./cmd/concoct
bash -n cmd/concoct/concoct
```

## Naming conventions

- Use hyphens, not underscores, in file and directory names.
- Keep conventional, long-lived project artifacts such as `AGENTS.md`, `README.md`, `CHANGELOG.md`, and `CONTRIBUTING.md` at the root.
- Use uppercase Markdown filenames for long-lived project-level artifacts.
- Use lowercase hyphenated Markdown filenames for task and workflow artifacts.

## Manual repository rename

The local content is branded for the intended repository name `concoct`. A repository owner must still rename the GitHub repository, update its description and topics, confirm clone URL redirects, and update local remotes where needed.
