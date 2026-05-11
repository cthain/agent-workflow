# Prompt: Archive completed task

Use this when a task is complete and the current planning artifacts should be moved into the archive.

```text
The active task is complete.

Archive the task artifacts from:

- `.planning/current/task-plan.md`
- `.planning/current/notes.md`

Move them to:

- `.planning/archive/YYYY-MM-DD-short-task-name/`

Create:

- `.planning/archive/YYYY-MM-DD-short-task-name/summary.md`

If `.planning/current/review.md` exists, archive it too.

The summary should include:

- task
- outcome
- key decisions
- files changed
- test results
- skipped work
- follow-up tasks

Use hyphens in the archive directory name.

After archiving, leave `.planning/current/` ready for the next task.
```
