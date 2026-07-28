# Workflow

This workflow turns ideas into implementation-ready plans that any capable coding agent can execute.

## The loop

```text
Idea
  ↓
Planner turns the idea into task artifacts
  ↓
Implementer reads canonical instructions and task artifacts
  ↓
Implementer changes the code
  ↓
Reviewer checks the result
  ↓
Archivist records the outcome
  ↓
Everybody's eatin' big time
```

The roles may be handled by different tools or by the same tool in different modes.

Each role has reusable guidance under `personas/`. A task prompt selects the appropriate persona; the persona supplements the canonical project instructions and active task context.

## Durable files

The workflow depends on durable files in the repository.

### Canonical project instruction

```text
AGENTS.md
```

This is the source of truth for project-level agent instructions.

It should describe:

- project intent
- design principles
- package boundaries
- coding style
- naming conventions
- verification commands
- planning workflow expectations

### Active task files

```text
.planning/current/task-plan.md
.planning/current/notes.md
```

These describe the current task and durable task context.

### Archive files

```text
.planning/archive/YYYY-MM-DD-short-task-name/
  task-plan.md
  notes.md
  summary.md
```

These preserve completed work for future reference.

## Agent-specific adapters

Different tools may read different instruction files.

Keep those files thin.

They should point back to `AGENTS.md` instead of duplicating project rules.

Examples:

```text
CLAUDE.md
CONVENTIONS.md
.github/copilot-instructions.md
.codex/skills/project-planning/SKILL.md
```

## Personas

Use one role persona at a time for the primary work being performed:

- `personas/planner.md` for creating or materially revising a plan
- `personas/code-developer.md` for implementation
- `personas/code-reviewer.md` for independent review
- `personas/doc-technical-writer-user.md` for end-user documentation
- `personas/doc-technical-writer-api.md` for API documentation
- `personas/doc-technical-writer-code.md` for code-oriented developer documentation

Read the selected persona before starting the role. If a task spans implementation and documentation, use the developer persona for implementation and then explicitly switch to the appropriate writer persona for the documentation pass.

## When to use planning files

Use planning files for:

- architecture changes
- refactors touching multiple files
- public API changes
- new features
- repository setup
- multi-session work
- tasks where losing context would be costly

Skip planning files for:

- typo fixes
- tiny one-file edits
- quick explanations
- throwaway experiments

## Good operating principle

Planning files are working memory, not bureaucracy.

Write down what helps future work.

Do not log noise.
