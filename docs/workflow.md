# Concoct Workflow

`Concoct` turns ideas into implementation-ready plans that any capable coding agent can execute.

## The loop

```text
Idea
  ↓
Product Owner adds the idea to the product roadmap
  ↓
Planner turns the product roadmap into task artifacts
  ↓
Developer implements tasks
  ↓
Reviewer checks the result
  ↓
Archivist records the outcome
  ↓
Writers document the project
  ↓
Eatin' big time
```

In short: `idea → concoct → eatin’ big time`.

Concoct's source assets live in this repository's root-level directories. Projects initialized by Concoct keep their task state and personas under `.concoct/`, with `AGENTS.md` and tool-required adapters at the project root.

The roles may be handled by different tools or by the same tool in different modes.

Each role has reusable guidance under `.concoct/personas/`. A task prompt selects the appropriate persona; the persona supplements the canonical project instructions and active task context.

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
.concoct/current/task-plan.md
.concoct/current/notes.md
```

These describe the current task and durable task context.

### Archive files

```text
.concoct/archive/YYYY-MM-DD-roadmap-id-short-task-name/
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
.codex/skills/concoct/SKILL.md
```

## Personas

Use one role persona at a time for the primary work being performed:

- `.concoct/personas/planner.md` for creating or materially revising a plan
- `.concoct/personas/code-developer.md` for implementation
- `.concoct/personas/code-reviewer.md` for independent review
- `.concoct/personas/doc-technical-writer-user.md` for end-user documentation
- `.concoct/personas/doc-technical-writer-api.md` for API documentation
- `.concoct/personas/doc-technical-writer-code.md` for code-oriented developer documentation

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
