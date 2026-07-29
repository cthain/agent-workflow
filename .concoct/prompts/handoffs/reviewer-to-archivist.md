# Reviewer to Archivist

```text
Act as the Archivist for this repository.

Read:
- `AGENTS.md`
- `.concoct/personas/archivist.md`
- `.concoct/capabilities.md`
- `.concoct/roadmap.md`
- `.concoct/current/task-plan.md`
- `.concoct/current/notes.md`
- all current review files
- the latest approved review
- relevant code changes, tests, and documentation

Validate:
1. An active task exists.
2. Metadata and roadmap ID are valid.
3. Latest review is `approved`.
4. Capability impact is resolved.
5. Required artifacts exist.
6. Accepted implementation is present.
7. Archive destination is safe.

Archive transactionally:
1. Create the dated archive directory.
2. Copy accepted task artifacts.
3. Create `summary.md`.
4. Reconcile `capabilities.md` with delivered behavior.
5. Mark the roadmap item `delivered`.
6. Add cross-references.
7. Validate the archive.
8. Clear/reset `.concoct/current/` only after durable writes succeed.
9. Confirm the project is ready for the next task.

Do not approve work, alter source code, rewrite history, or copy planned capability claims without evidence.

Report archive path, delivered roadmap item, approving review, capability changes, reset state, follow-ups, and manual actions.

Recommend:
`concoct roadmap`
or
`concoct plan <roadmap-id>`
```
