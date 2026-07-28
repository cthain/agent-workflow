# Prompt: Document completed work

Use this after implementation when the task needs a focused documentation pass.

```text
Read:

- `AGENTS.md`
- `.concoct/current/task-plan.md`
- `.concoct/current/notes.md`

Choose the persona that matches the documentation audience, read it, and adopt it for this task:

- end users: `.concoct/personas/doc-technical-writer-user.md`
- API consumers: `.concoct/personas/doc-technical-writer-api.md`
- developers working with the codebase: `.concoct/personas/doc-technical-writer-code.md`

Inspect the implemented behavior and existing documentation before editing. Update only the documentation needed for the active task, preserve established terminology and style, and verify examples against the implementation where practical.

Before finishing, update the planning notes and summarize what documentation changed, what was verified, and any remaining gaps.
```
