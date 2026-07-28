# Agent Planning Workflow

A lightweight, agent-neutral workflow for turning ideas into implementation-ready plans.

The core loop is:

1. Good idea.
1. Turn it into a plan.
1. Agent(s) ship(s) it.
1. Eatin’ big time.

This project is designed for working with coding agents such as Codex, Claude Code, GitHub Copilot coding agent, Aider, or future agent tools.

The durable project/task context lives in ordinary repository files. Agent-specific files are thin adapters that point back to the canonical instructions.

## Core idea

Keep the workflow agent-neutral:

```text
AGENTS.md                         # canonical project instructions
.planning/current/task-plan.md    # active task
.planning/current/notes.md        # durable task context
.planning/archive/...             # completed task history
personas/...                      # reusable role-specific working guidance
```

Then add thin adapters for specific tools:

```text
CLAUDE.md
CONVENTIONS.md
.aider.conf.yml
.github/copilot-instructions.md
.codex/skills/project-planning/SKILL.md
```

## What this is

This is a small starter kit for Manus-style file-based planning without the heavy ceremony.

It helps agents avoid common failure modes:

- forgetting context
- drifting from the task
- repeating failed actions
- losing important decisions after a context reset
- mixing unrelated cleanup into the active task

## What this is not

This is not a framework.

This is not a project management system.

This is not meant to create planning bureaucracy for every tiny change.

Use it when the task has enough complexity that losing context would hurt.

## Repository structure

```text
.
├── README.md
├── docs/
│   ├── workflow.md
│   └── multi-agent-workflow.md
├── prompts/
│   ├── create-task-plan.md
│   ├── continue-task.md
│   ├── review-task.md
│   ├── archive-task.md
│   ├── multi-agent-handoff.md
│   └── document-task.md
├── personas/
│   ├── planner.md
│   ├── code-developer.md
│   ├── code-reviewer.md
│   ├── doc-technical-writer-user.md
│   ├── doc-technical-writer-api.md
│   └── doc-technical-writer-code.md
└── templates/
    ├── AGENTS.md
    ├── CLAUDE.md
    ├── CONVENTIONS.md
    ├── .aider.conf.yml
    ├── .planning/
    │   └── current/
    │       ├── task-plan.md
    │       └── notes.md
    ├── .codex/
    │   └── skills/
    │       └── project-planning/
    │           └── SKILL.md
    └── .github/
        ├── copilot-instructions.md
        └── prompts/
            ├── create-task-plan.prompt.md
            ├── implement-task.prompt.md
            ├── review-task.prompt.md
            └── document-task.prompt.md
```

## How to use this in a project

Run `./init <new-project-name> [destination-parent]` to create a repository with the templates and personas installed. To install the workflow manually, copy the contents of `templates/` and the `personas/` directory into the target repository root.

At minimum, copy:

```text
AGENTS.md
.planning/current/task-plan.md
.planning/current/notes.md
personas/
```

Then customize `AGENTS.md` for the project.

Tool-specific adapters are optional.

## Recommended minimum setup

```text
AGENTS.md
.planning/current/task-plan.md
.planning/current/notes.md
personas/
```

## Recommended multi-agent setup

```text
AGENTS.md
CLAUDE.md
CONVENTIONS.md
.aider.conf.yml
.github/copilot-instructions.md
.codex/skills/project-planning/SKILL.md

.planning/current/task-plan.md
.planning/current/notes.md
personas/
```

## Naming conventions

Use hyphens, not underscores, in file and directory names.

Use uppercase Markdown filenames for long-lived project-level artifacts:

```text
AGENTS.md
README.md
CHANGELOG.md
CONTRIBUTING.md
CLAUDE.md
CONVENTIONS.md
```

Use lowercase hyphenated Markdown filenames for task-specific artifacts:

```text
task-plan.md
notes.md
summary.md
design-sketch.md
review.md
```

## Suggested first prompt

```text
Use the agent planning workflow.

Read and adopt personas/planner.md for this task.

Turn this idea into Codex/agent-ready planning artifacts for:

- .planning/current/task-plan.md
- .planning/current/notes.md

Keep the plan useful, not ceremonial.

Here is the idea:

[PASTE IDEA HERE]
```

## Suggested implementer prompt

```text
Read AGENTS.md, .planning/current/task-plan.md, and .planning/current/notes.md.

Read and adopt personas/code-developer.md for this task.

Inspect the repository before making code changes.

Update the plan if inspection shows it is wrong or incomplete.

Then implement the task in small, understandable steps.

Before finishing, run the documented checks, update the planning files, and summarize what changed, what passed, and what remains.
```
