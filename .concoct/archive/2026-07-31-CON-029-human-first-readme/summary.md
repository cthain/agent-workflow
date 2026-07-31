---
task-id: CON-029
roadmap-id: CON-029
status: archived
archived: 2026-07-31
review: review-01.md
delivery: pending-integration
capability-impact:
  type: none
  ids: []
---

# Summary

## Task

Introduce Concoct through a human-first root README that helps a developer
understand the product, assess its fit, and follow a representative Git-backed
workflow before encountering repository internals.

## Delivered outcome

The README now opens with Concoct's product purpose, intended audience, durable
workflow value, role boundaries, and present automation limits. Its quick start
begins with `concoct init hello-world` and walks through initial commit,
status, next-action selection, roadmap intake, planning, implementation,
independent review and remediation, archival, and local Git integration.

Each transition distinguishes bounded CLI effects from the Product Owner,
Planner, Developer, Reviewer, and Archivist work performed by a human or coding
agent. Source-build, installed-layout, contributor, naming, and repository
maintenance material remains available after the user journey, with focused
links to the normative workflow documentation.

## Key decisions

- Put product promise, fit, limitations, and the representative journey before
  installed structure and contributor details.
- Use concise actor/effect explanations so prompt rendering is not mistaken for
  autonomous role execution.
- Describe the current Git-backed happy path while keeping uncommon invalid
  states and recovery detail in normative documentation.
- Preserve source-build instructions without inventing a package or release
  channel.

## Files and areas changed

- Rewrote the root `README.md`.
- Updated active task artifacts with planning, implementation, verification,
  review, and archival evidence.

## Verification

- `go run ./cmd/concoct help` passed and confirmed every displayed command.
- The quick start was traced against `doc/command-reference.md`,
  `doc/state-machine.md`, CAP-001, CAP-005, CAP-006, and CAP-007.
- Every README Markdown link resolved to an existing repository path.
- Searches found no unintended autonomous-agent or semantic-correctness claim.
- `git diff --check` passed.
- `go test ./...` and `go vet ./...` were not run because the accepted change
  is Markdown-only, as permitted by the approved verification plan.

## Review outcome

`review-01.md` approved the implementation with no material findings. It
confirmed the onboarding narrative, command ordering, actor/effect boundaries,
Git lifecycle claims, agent-neutral language, scope discipline, and
documentation accuracy.

## Capability changes

None. CON-029 improves presentation and adoption of CAP-001, CAP-005, CAP-006,
and CAP-007 without changing observable behavior or accepted capability truth.

## Follow-up work

Delivery remains pending `concoct integrate`. CON-029 stays active and current
task evidence remains intact until integration succeeds.

## References

- Roadmap item: `CON-029` in `.concoct/roadmap.md`
- Approving review: `review-01.md`
- Product documentation: `README.md`
- Normative contracts: `doc/workflow.md`, `doc/command-reference.md`,
  `doc/state-machine.md`
