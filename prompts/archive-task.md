# Prompt: Archive completed task

Use this when a task is complete and the current planning artifacts should be moved into the archive.

```text
The active task is complete.

Read `AGENTS.md` and the active planning artifacts before archiving. This is an archival operation, so do not adopt an implementation or review persona.

Archive the task artifacts from:

- `.concoct/current/task-plan.md`
- `.concoct/current/notes.md`

Move them to:

- `.concoct/archive/YYYY-MM-DD-short-task-name/`

Create:

- `.concoct/archive/YYYY-MM-DD-short-task-name/summary.md`

If `.concoct/current/review.md` exists, archive it too.

The summary should include:

- task
- outcome
- key decisions
- files changed
- test results
- skipped work
- follow-up tasks

Use hyphens in the archive directory name.

After archiving, leave `.concoct/current/` ready for the next task.
```
