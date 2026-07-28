---
name: concoct
description: Work on substantial Concoct repository tasks using its internal file-based planning artifacts.
user-invocable: true
allowed-tools: "Read Write Edit Bash Glob Grep"
---

# Concoct Project Planning

Use this skill for substantial work on Concoct itself.

## Canonical context

Read and follow:

- `AGENTS.md`
- `current/task-plan.md`
- `current/notes.md`

Read the persona under `personas/` selected for the current role.

## Workflow

1. Inspect the repository before editing.
2. Keep `current/task-plan.md` aligned with the intended outcome.
3. Record durable decisions and meaningful failures in `current/notes.md`.
4. Preserve the distinction between Concoct's root-level source layout and the `.concoct/` directory installed into client projects from `templates/`.
5. Run the checks documented in `AGENTS.md`.
6. Report what changed, what passed, and what remains.
