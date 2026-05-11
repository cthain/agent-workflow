# AGENTS.md

## Project intent

Describe what this project is, who it is for, and what it should be good at.

Keep this section durable. It should not describe one temporary task.

## Canonical instruction file

`AGENTS.md` is the canonical project instruction file.

Agent-specific files may point back to `AGENTS.md`, but should not redefine project rules unless the tool requires different mechanics.

If another instruction file conflicts with this one, prefer this file unless a higher-priority system or user instruction says otherwise.

## Design principles

- Keep the core model small and understandable.
- Prefer explicit code over hidden magic.
- Prefer boring, maintainable solutions over clever abstractions.
- Preserve public API compatibility where practical.
- Do not keep bad abstractions only to avoid changing a prototype.
- Tests should demonstrate behavior, not implementation details.
- Make failure modes easy to understand.
- Keep changes focused on the active task.

## Package boundaries

Core should own:

- _Add project-specific core responsibilities here._

Core should not own:

- _Add project-specific non-core responsibilities here._

## Coding style

- Use the project's standard formatter.
- Keep names short but clear.
- Avoid unnecessary global state.
- Avoid hidden initialization behavior.
- Prefer small interfaces defined near the consumer.
- Return useful errors with enough context to diagnose failures.
- Keep examples realistic and minimal.

## Naming conventions

Use hyphens, not underscores, in file and directory names.

Reason: underscores are easily obscured in browser links and rendered text.

Use uppercase Markdown filenames for long-lived project-level artifacts:

- `AGENTS.md`
- `README.md`
- `CHANGELOG.md`
- `CONTRIBUTING.md`
- `CLAUDE.md`
- `CONVENTIONS.md`

Use lowercase hyphenated Markdown filenames for task-specific or working artifacts:

- `task-plan.md`
- `notes.md`
- `summary.md`
- `design-sketch.md`
- `review.md`

## Planning artifacts

Use file-based planning for substantial multi-step work.

Do not create planning artifacts for trivial edits, typo fixes, or one-shot changes.

Active planning files live under:

```text
.planning/current/
  task-plan.md
  notes.md
```

Optional active review file:

```text
.planning/current/
  review.md
```

Only create `review.md` when there is a meaningful separate review phase.

Completed planning artifacts should be archived under:

```text
.planning/archive/YYYY-MM-DD-short-task-name/
  task-plan.md
  notes.md
  summary.md
```

If a review file was created, archive it too:

```text
.planning/archive/YYYY-MM-DD-short-task-name/
  review.md
```

The archive directory name should use the date the task was completed or archived.

## Planning file purposes

### `.planning/current/task-plan.md`

The active task plan.

It should describe:

- the current goal
- why the task matters
- design constraints
- non-goals
- working assumptions
- implementation phases or checklist
- completion criteria

This file is per task. Treat it like a user story, work item, or implementation ticket.

### `.planning/current/notes.md`

Durable working context for the active task.

Use it for:

- decisions
- findings
- failed attempts
- gotchas
- handoffs
- follow-up ideas

This file is per `task-plan.md`.

Keep it useful. Do not turn it into a transcript.

### `.planning/current/review.md`

Optional separate review artifact.

Use it only when a distinct reviewer agent or review pass is part of the workflow.

### Archived `summary.md`

When a task is complete, create a short `summary.md` in the archive directory.

It should describe:

- what changed
- key decisions
- test results
- skipped work
- remaining follow-ups

The summary should be clean enough to read months later without needing to reconstruct the whole task.

## Planning workflow

For substantial work:

1. Read this file.
2. Read `.planning/current/task-plan.md`.
3. Read `.planning/current/notes.md` if it exists.
4. Inspect the relevant code before making changes.
5. Update the plan if inspection reveals the existing plan is wrong or incomplete.
6. Make changes in small, understandable steps.
7. Record important findings, failed attempts, or handoff notes in `notes.md`.
8. Before finishing, update task status and notes.
9. Run the relevant project checks.
10. Archive the planning artifacts when the task is complete.

Planning files are working memory, not bureaucracy. Prefer compact, useful updates over exhaustive journaling.

## Agent coordination

All agents should treat these as shared task context:

- `.planning/current/task-plan.md`
- `.planning/current/notes.md`

Do not overwrite another agent's work without first reading the current files.

When handing off between agents, update `notes.md` with:

- what changed
- what remains
- known risks
- commands already run
- commands that still need to be run
- suggested next step

Suggested handoff format:

```md
## Handoff

### Current state

### Completed

### Remaining

### Known risks

### Commands run

### Suggested next step
```

## Error handling during agent work

Do not blindly repeat failed actions.

When something fails:

1. Read the error carefully.
2. Identify the likely root cause.
3. Try a targeted fix.
4. If the same approach fails again, change approach.
5. Record the failed attempt in `notes.md` if it affects future work.

Use this format when recording meaningful failed attempts:

```md
### Attempt: short title

- Tried:
- Error/result:
- Why it failed:
- Next approach:
```

## Commands

Replace these with the correct commands for this project.

```bash
# format
# test
# lint / vet / typecheck
```

If the repository documents additional commands, prefer the repository's documented workflow.

If a command fails, record the failure and either fix it or clearly explain why it remains unresolved.

## Completion expectations

Before finishing a substantial task:

- Update `.planning/current/task-plan.md`.
- Update `.planning/current/notes.md`.
- Run formatting.
- Run tests.
- Run linting/vetting/typechecking where applicable.
- Summarize what changed.
- Summarize what passed.
- Summarize what remains.

If the task is complete, archive the planning files under `.planning/archive/YYYY-MM-DD-short-task-name/` and add a `summary.md`.

## Git hygiene

Keep changes focused on the active task.

Do not mix unrelated cleanup with the requested work unless it is necessary to complete the task.

When reviewing changes, separate:

- functional changes
- test changes
- documentation changes
- planning artifact updates
