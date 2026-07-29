---
task-id: none
roadmap-id: none
status: historical-override
archived: 2026-07-29
review: none
capability-impact:
  type: none
  ids: []
---

# Legacy HITL restructuring summary

## Archive context

This archive preserves task artifacts created during human-in-the-loop restructuring and prototyping before Concoct began using its current workflow on its own repository. The human owner explicitly authorized archival despite the absence of a roadmap identifier, artifact metadata, and an approved review.

The archived plan and notes are preserved verbatim. Their claims describe the repository state and checks at the time; later restructuring superseded parts of them. This archive does not retroactively approve those claims or mark a roadmap item delivered.

## Historical outcome

The work established the distinction between reusable Concoct source assets and the workflow material installed in client repositories. It also explored keeping client-owned workflow state under `.concoct/` while leaving conventional tool adapters at the generated project root.

Subsequent prototyping changed Concoct's own layout again: this repository now keeps its living workflow artifacts under `.concoct/`, the executable under `cmd/concoct/`, reusable templates under `templates/`, and documentation under `doc/`.

## Initial capability inventory findings

Repository inspection on 2026-07-29 supported four current capability records:

- `CAP-001` — a durable, Markdown-based workflow contract;
- `CAP-002` — manual prompts for role transitions;
- `CAP-003` — a reusable but currently limited project workflow template;
- `CAP-004` — thin adapters for multiple coding-agent tools.

The inspection also found that the checked-in initializer passes shell syntax validation but cannot complete initialization from its current location because it searches for `templates/` and `personas/` under `cmd/concoct/`. Stale persona references, empty writer persona files, and the lack of automated tests were recorded as capability gaps rather than delivered behavior.

## Verification and evidence

- `git log` traces the archived task plan to commit `3e38e72` dated 2026-07-28.
- `bash -n cmd/concoct/concoct` passed during the capability inventory.
- An end-to-end initializer invocation failed with `templates directory not found at: .../cmd/concoct/templates`.
- `git diff --check` passed before archival.

## Capability impact

Archiving the stale task changes no observable product behavior. The initial capability inventory was created independently from current repository evidence; this archive provides historical traceability for that inventory without acting as an approving review.

## Follow-up work

- Define the command and state-machine contract in `CON-003`.
- Restore and test initialization as part of the CLI foundation in `CON-005`.
- Implement deterministic prompt rendering in `CON-006`.
- Reconcile stale documentation and persona filenames, and decide whether empty writer personas should be implemented or removed from the shipped template.

## References

- `.concoct/capabilities.md`
- `.concoct/roadmap.md`
- `AGENTS.md`
- `README.md`
- `cmd/concoct/concoct`
