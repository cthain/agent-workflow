# GitHub Copilot Instructions

Read and follow the repository's canonical agent instructions in `AGENTS.md`.

Task prompts select a role from `.concoct/personas/`. Read and adopt the selected persona for that task; `AGENTS.md` remains authoritative if instructions conflict.

For substantial implementation tasks, use the active planning files:

```text
.concoct/current/task-plan.md
.concoct/current/notes.md
```

Before making changes:

1. Read `AGENTS.md`.
2. Read the active planning files.
3. Inspect relevant code.

Before finishing:

1. Update the active planning files.
2. Run the documented project checks.
3. Summarize what changed, what passed, and what remains.

Keep changes focused on the active task.
