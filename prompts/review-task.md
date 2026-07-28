# Prompt: Review implementation

Use this after an agent has implemented a task.

```text
Adopt the code reviewer persona in `.concoct/personas/code-reviewer.md` for this task. Read it before beginning.

Review the completed changes against:

- `AGENTS.md`
- `.concoct/current/task-plan.md`
- `.concoct/current/notes.md`

Give me a concise review with:

1. Whether the implementation satisfies the task.
2. Any design concerns.
3. Any testing gaps.
4. Any documentation gaps.
5. Any unrelated changes.
6. Whether the task is ready to archive.
7. Suggested contents for `summary.md` if it is ready.
```
