---
name: project-planning
description: Lightweight file-based planning for coding-agent tasks that touch multiple files, require architectural judgment, or may span multiple sessions.
user-invocable: true
allowed-tools: "Read Write Edit Bash Glob Grep"
---

# Project Planning

Use this skill for substantial coding work.

Do not use it for trivial edits, typo fixes, or one-shot answers.

## Canonical instructions

Read and follow:

```text
AGENTS.md
```

`AGENTS.md` is the canonical project instruction file.

## Files

Use these files in the project root:

- `AGENTS.md` — standing repository instructions
- `.planning/current/task-plan.md` — the active plan
- `.planning/current/notes.md` — durable decisions, findings, failed attempts, handoffs, and follow-up ideas

Do not create extra planning files unless the task genuinely needs them.

## Workflow

### 1. Re-orient

Before editing code:

1. Read `AGENTS.md`.
2. Read `.planning/current/task-plan.md`.
3. Read `.planning/current/notes.md` if it exists.

### 2. Plan only enough

Keep the plan useful, not ceremonial.

A good plan includes:

- goal
- context
- constraints
- non-goals
- phases or checklist
- completion criteria

### 3. Write down durable findings

Update `.planning/current/notes.md` when you discover something that affects the implementation, design, tests, or future work.

Do not log obvious facts or noisy command output.

### 4. Avoid dumb retry loops

If an action fails, do not repeat the exact same action blindly.

Record the failure in `.planning/current/notes.md` when it matters.

### 5. Finish cleanly

Before finishing:

1. Update `.planning/current/task-plan.md` statuses.
2. Update `.planning/current/notes.md` with important decisions or remaining follow-up work.
3. Run the relevant project checks.
4. Report what changed, what passed, and what remains.
5. If the task is complete, archive the planning files and create `summary.md`.

## Good judgment

Planning files are working memory, not bureaucracy.

Prefer useful, compact updates over exhaustive journaling.
