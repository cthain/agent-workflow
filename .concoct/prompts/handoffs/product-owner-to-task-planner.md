# Product Owner to Task Planner

```text
Act as the Task Planner for this repository.

Read:
- `AGENTS.md`
- `.concoct/personas/task-planner.md`
- `.concoct/capabilities.md`
- `.concoct/roadmap.md`
- relevant archive summaries
- relevant source code, tests, and documentation

Roadmap item:
`<ROADMAP-ID>`

Confirm:
1. The item exists and is ready for planning.
2. Dependencies are satisfied or explicitly handled.
3. No conflicting active task exists.
4. Repository reality supports the roadmap assumptions.
5. No unresolved product decision must return to the Product Owner.

When ready, create:
- `.concoct/current/task-plan.md`
- `.concoct/current/notes.md`

The task plan must define:
- identifiers and metadata;
- goal, context, current state, and target state;
- constraints and non-goals;
- assumptions, risks, and open questions;
- implementation phases;
- observable acceptance criteria;
- verification;
- capability impact;
- developer handoff expectations.

Do not implement code or invent product direction.

Summarize planning readiness, scope, capability impact, risks, and unresolved questions.

When ready, recommend:
`concoct code`
```
