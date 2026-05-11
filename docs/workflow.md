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
