---
task-id: CON-015
roadmap-id: CON-015
status: delivered
archived: 2026-07-30
review: review-03.md
capability-impact:
  type: add
  ids:
    - CAP-007
---

# Summary

## Task

Add an optional Git-backed lifecycle that isolates substantial Concoct tasks on
dedicated branches and integrates approved, archived work into the exact trunk
from which each task began.

## Delivered outcome

Concoct now establishes deterministic task branches during Git-backed planning,
records durable trunk, branch, and base identity, and validates that development,
review, and archival occur at safe repository boundaries. Approved work is
archived on its task branch before local squash integration into the recorded
trunk. Integration supports conflict continuation, exact abort, interruption
recovery, final delivery bookkeeping, accepted branch cleanup, optional guarded
upstream pushes, and local repositories without remotes.

Non-Git projects retain the existing unbranched archival workflow. CON-015
itself was planned under that pre-existing workflow, so its accepted changes
were archived directly on its original branch as the documented transitional
exception.

## Key decisions

- Record task branch, integration trunk, immutable base, archival commit, and
  integration evidence instead of inferring identity from the ambient checkout.
- Define clean workflow boundaries as an empty `git status --short`.
- Keep incomplete integration recovery under `.git/concoct/integrations/`
  while committed archive evidence remains authoritative.
- Use local squash integration and require human-attested semantic conflict
  resolution; Concoct validates repository structure and transaction bounds.
- Bound destructive abort behavior to exact recorded commits and paths changed
  by the archived task.
- Prompt before pushing a matching upstream by default and require explicit
  configuration for automatic pushes.

## Files and areas changed

- Added reusable Git repository operations and transactional integration logic
  under `internal/gitrepo/` and `internal/integration/`.
- Extended CLI, workflow detection, prompt rendering, tests, and golden evidence
  for planning, archival, integration, recovery, and Git-aware states.
- Updated README, command, state-machine, workflow, persona, handoff, skill, and
  embedded-template contracts.
- Added real-repository coverage for success, conflict, continue, abort,
  interruption, collision, drift, unsafe worktrees, branch cleanup, upstream
  behavior, and non-Git fallback.

## Verification

- `go test -count=1 ./...` — passed.
- `go vet ./...` — passed.
- `go build ./cmd/concoct` — passed.
- `bash -n cmd/concoct/concoct.sh` — passed.
- Source and generated wrapper executable-mode checks — passed.
- Source/template counterpart comparison — passed.
- Temporary initialization confirmed Git creation, staged root/dot/nested
  templates, bootstrap prompt, personas, planning directories, no generated
  commit, executable wrapper, and `ready` status.
- `git diff --check` — passed before archival.

## Review outcome

`review-03.md` approved the implementation after two remediation cycles.
The final review confirms that prepared-state recovery, exact phase and HEAD
validation, and transaction-path-bounded abort behavior resolve all destructive
recovery findings from `review-01.md` and `review-02.md`.

## Capability changes

Added `CAP-007` for optional Git-backed task isolation, archival, squash
integration, guarded recovery, local-only operation, and non-Git fallback.

## Skipped and follow-up work

Provider and pull-request integration, autonomous semantic conflict resolution,
worktrees, concurrent or stacked tasks, and automatic push without explicit
opt-in remain out of scope.

## References

- Roadmap item: `CON-015` in `.concoct/roadmap.md`
- Capability: `CAP-007` in `.concoct/capabilities.md`
- Approving review: `review-03.md`
- Prior reviews: `review-01.md`, `review-02.md`
- User documentation: `README.md`
- Normative contracts: `doc/command-reference.md`, `doc/state-machine.md`,
  `doc/workflow.md`

