---
id: CON-015
title: Isolate and integrate tasks with Git branches
roadmap-id: CON-015
status: implementation-complete
created: 2026-07-30
updated: 2026-07-30
remediates-review: review-02.md
capability-impact:
  type: add
  ids:
    - CAP-007
  rationale: Adds an optional Git-backed task branch, archival, and integration lifecycle while preserving the non-Git workflow.
---

# Task Plan

## Goal

Implement the complete `CON-015` Git branch lifecycle so a Git-backed Concoct task records its starting trunk, runs on a deterministic task branch, archives before integration, squash-integrates into that trunk, and returns to `ready` only after successful final bookkeeping and cleanup. Preserve the existing unbranched workflow for non-Git projects.

## Context

The Go CLI currently initializes Git and provides read-only state detection plus deterministic prompt rendering for `roadmap`, `plan`, `code`, and `review`. Workflow state is inferred from Markdown artifacts; it has no durable Git identity, archive/integration states, archive or integration commands, or checkout/worktree validation.

`CON-015` supplies the branch lifecycle required by later active-planning work (`CON-007`) and changes archival from an immediate return to `ready` into `archived → integrating/integrated → ready`. The implementation must extend the current explicit packages and installed artifact contracts without requiring a provider or remote.

## Why this matters

Task work is currently mixed with the user's checked-out branch. A recorded task branch and source trunk make the implementation/review boundary inspectable, keep accepted changes isolated until archival, and make integration and recovery explicit.

## Current state

- `internal/cli` recognizes `init`, `status`, and four prompt-rendering commands; no archive or integration command exists.
- `internal/project` shells out to Git only during initialization and has no reusable repository abstraction.
- `internal/workflow` recognizes `ready` through review outcomes from Markdown evidence. Task metadata has no trunk, branch, base, archival commit, or integration evidence.
- `internal/prompt` validates artifact state and renders role guidance but does not validate checkout identity or describe transition commits.
- The command/state documentation defines the pre-branch lifecycle and transactional archival returning directly to `ready`.
- The embedded `templates/` tree is the installed product contract, so applicable schema and guidance changes must be synchronized there.

## Target state

- Git-backed planning establishes a deterministic task branch from the current branch and records the integration trunk and immutable task base before implementation.
- Task-scoped transitions validate clean input/output boundaries and reject checkout drift or unsafe Git operations without overwriting work.
- Role-produced changes are committed at defined transition boundaries, with Archivist ownership of the archival commit.
- Approved archival validates archive/capability/pending-roadmap evidence on the task branch and ends in `archived`; it does not mark the item delivered.
- `concoct integrate`, `concoct integrate --continue`, and `concoct integrate --abort` implement squash integration and recovery using local `.git/concoct/integrations/` evidence only while incomplete.
- Successful integration records final evidence, marks the item delivered, clears active state, removes the accepted task branch, leaves the recorded trunk checked out, and reports `ready`.
- A non-Git project retains the existing local, unbranched lifecycle.

## Design constraints

- Treat recorded branch/trunk/base identity as durable evidence; never infer task identity solely from the checkout.
- Derive a portable deterministic branch name from roadmap ID and short name, with documented normalization and collision behavior.
- Use a small explicit Git boundary tested with real temporary repositories; avoid a framework-heavy CLI or provider API.
- Define clean as empty `git status --short` at workflow input/output boundaries. Unpushed commits are permitted.
- Never discard a dirty worktree, rewrite reviews, overwrite a collision, resolve semantic conflicts, or delete recovery evidence before validation.
- Preserve task branch, archival commit, pre-integration trunk head, and recovery metadata whenever integration is incomplete.
- Keep committed archive evidence authoritative; `.git/concoct/` is local recovery state only.
- Squash-merge into the exact recorded trunk, not an inferred conventional branch.
- Prompt before pushing a matching upstream unless configuration enables automatic pushes; do not push without a matching upstream.
- Keep source and template contracts synchronized, preserve the shell wrapper's executable mode, and retain agent neutrality.
- Mark the roadmap item delivered only after integration succeeds.

## Non-goals

- No hosting-provider, pull-request, remote-branch, or review-service integration.
- No autonomous semantic conflict resolution.
- No assumption of `main`, `master`, or a remote default branch.
- No worktree concurrency, stacked branches, or multiple active tasks.
- No broad implementation of `CON-007`, `CON-008`, or `CON-009` beyond behavior strictly required for this Git lifecycle.
- No CLI framework migration, generalized Git library, or unrelated cleanup.
- No automatic push without explicit project opt-in.

## Working assumptions

- `CAP-007` is the next stable capability ID; the Archivist may refine its title while preserving planned behavior.
- Full local branch names are recorded. Detached HEAD, unborn/ambiguous source history, active Git operations, dirty worktrees, and non-matching branch collisions are unsafe planning inputs.
- The Developer may choose the precise committed metadata schema if it supports deterministic validation and recovery and is synchronized with templates/docs.
- Transition commits may be finalized by an applicable CLI transition or explicitly required role handoff, but ownership, validation order, contents, and retries must be defined.
- Prompt output remains deterministic for fixed durable state; ambient Git affects it only through validation.

## Risks and open questions

- `CON-015` intersects `CON-007` through `CON-009`; implement every Git-specific hook needed here without absorbing their unrelated automation.
- CLI invocation and role work are separate moments. The implementation must define a two-phase boundary so interruption cannot make a commit look like completed role evidence.
- Final bookkeeping follows the squash commit while the task branch is due for deletion; ordering must keep integration evidence and delivered status recoverable on trunk.
- Conflict recovery changes the trunk index/worktree and must distinguish Concoct's operation from unrelated Git operations.
- Push prompting must remain usable non-interactively and a declined/unavailable push must not invalidate local delivery.
- Non-Git tests must construct a project independently because `concoct init` always creates Git today.

## Implementation phases

### Phase 1 — Specify durable Git and state contracts

Status: `complete`

- [x] Choose canonical metadata for task branch, integration trunk, task base, archival commit, pre-integration trunk head, integration status/commit, and pending roadmap reconciliation.
- [x] Define branch normalization, collision checks, clean boundaries, non-Git fallback, transition-commit ownership, push policy, and interruption/idempotency invariants.
- [x] Update plan/notes for in-scope technical refinements; escalate product contradictions instead of guessing.

### Phase 2 — Add Git and state foundations

Status: `complete`

- [x] Introduce a small Git boundary for detection, branch/head/upstream and worktree inspection, operation checks, branch operations, commits, squash integration, and recovery primitives.
- [x] Extend artifact parsing and detection for Git metadata plus `archived`, `integrating`, and `integrated`, including invalid partial evidence and actionable next commands.
- [x] Preserve the documented non-Git workflow without fabricated Git metadata.
- [x] Test normalization/collisions, checkout drift, dirty/detached/unsafe starts, metadata, state precedence, and interruptions.

### Phase 3 — Enforce task transitions and archival

Status: `complete`

- [x] Extend planning behavior/guidance to record trunk/base, create and checkout the task branch, validate both active artifacts, and activate only the selected item.
- [x] Make code, remediation, and review entry points validate the task branch, clean boundaries, and transition-commit evidence while preserving reviews.
- [x] Add/extend Archivist behavior so approved work is reconciled and committed on the task branch, records pending delivery, and ends in `archived` with current evidence intact.
- [x] Define retry behavior around each transition commit without duplicates or false state advancement.

### Phase 4 — Implement integration and recovery

Status: `complete`

- [x] Add CLI help/parsing/dispatch for `concoct integrate`, `--continue`, and `--abort`, plus archive entry changes strictly required by this lifecycle.
- [x] Validate `archived` and squash-integrate the recorded archival commit into the recorded trunk, preserving pre-integration evidence and entering `integrating` on conflict.
- [x] Implement human-attested continue validation and exact abort restoration, rejecting unrelated operations and unresolved/unstaged evidence.
- [x] Complete `integrated → ready` idempotently: record integration commit, reconcile delivery, clear active state after validation, delete task branch/recovery metadata, and leave trunk checked out.
- [x] Implement upstream prompting/auto-push policy without making a remote necessary.

### Phase 5 — Synchronize contracts and verify

Status: `complete`

- [x] Update command/state/workflow/user docs, personas, prompts, status output, templates, and goldens for branch and non-Git behavior.
- [x] Add real-repository end-to-end tests for local success, collision/drift/dirty refusal, conflict/continue, abort/retry, interruption recovery, branch deletion, checkout, delivery, and no-remote operation.
- [x] Run all documented checks and record exact results in notes.
- [x] Leave a reviewer handoff focused on transactional safety, state precedence, destructive Git operations, non-Git compatibility, and source/template parity.

### Review 01 remediation

Status: `complete`

- [x] Guard abort and later-phase continuation with exact recorded HEAD validation and phase-appropriate operation/worktree checks.
- [x] Make `prepared` recovery safe to continue or abort while the task branch remains checked out.
- [x] Add real-repository coverage for dirty abort, unrelated trunk commits, unexpected later-phase HEAD, and prepared recovery before trunk checkout.
- [x] Rerun the complete verification suite and record both finding dispositions in notes.

### Review 02 remediation

Status: `complete`

- [x] Restrict pre-commit abort to index/worktree paths owned by the recorded archive integration.
- [x] Refuse and preserve staged additions or tracked-file modifications outside the task's merge-base diff.
- [x] Add real-repository coverage for staged new and staged unrelated tracked files during conflict recovery.
- [x] Rerun the complete verification suite and record the finding disposition in notes.

## Acceptance criteria

- [x] Safe Git planning deterministically creates/checks out the task branch and records exact trunk/base before implementation.
- [x] Branch normalization is documented/tested; collisions fail without changing branches, worktree, active artifacts, or roadmap status.
- [x] Code, remediation, review, and archive reject checkout drift, dirty boundaries, contradictory metadata, or unrelated Git operations with recovery guidance.
- [x] Every role-produced change has a validated retry-safe commit boundary; the Archivist creates the archival commit on the task branch.
- [x] Approved archival validates archive, capability, and pending roadmap reconciliation and reports `archived` without delivery or active cleanup.
- [x] `concoct integrate` locally squash-integrates into the recorded trunk, completes bookkeeping, marks delivery, clears active state, deletes task branch, and leaves a clean `ready` trunk.
- [x] Conflicts report `integrating` and preserve trunk checkout, task branch, archival commit, pre-integration head, conflict state, and local recovery evidence.
- [x] `--continue` accepts only resolved/staged human-attested results and safely resumes through `ready` without duplicate integration.
- [x] `--abort` restores the exact pre-integration trunk, preserves/checks out the approved task branch, returns to `archived`, and permits a fresh attempt.
- [x] Upstream behavior prompts by default, honors explicit auto-push, and treats no matching upstream as local success.
- [x] Valid non-Git projects continue through the documented unbranched workflow.
- [x] Status, prompts, docs, personas, templates, and tests agree on states, ownership, commands, and recovery.
- [x] Unsafe/interrupted scenarios never delete unintegrated history, overwrite work, falsely report delivery, or return to `ready` early.

## Verification

- `gofmt` changed Go files; `go test ./...`; `go vet ./...`; `go build ./cmd/concoct`.
- `bash -n cmd/concoct/concoct` and executable-mode inspection.
- Initialize under a temporary parent and confirm copied root/dot/nested files, personas/planning directories, bootstrap prompt, Git, staging, no commit, and initial status.
- Exercise happy-path local squash integration, no remote, collision, dirty worktree, checkout drift, conflict/continue, abort/retry, and interruptions around integration/bookkeeping.
- Exercise a valid non-Git fixture.
- `git diff --check` and searches for stale state lists, archive-to-ready claims, old command usage, missing template counterparts, branding, and paths.

## Capability impact

Add `CAP-007` for optional Git-backed task isolation/integration: deterministic branches, durable trunk identity, transition validation, archival/integration states, conflict recovery, local-only operation, and non-Git fallback. Capability truth changes only after acceptance and archival.

## Handoff expectations

Set `implementation-complete` only after the full lifecycle and proportional recovery coverage exist. Notes must record schema/transaction decisions, exact checks, risks/skips, changed files, and a fresh Reviewer handoff. Review should independently exercise destructive-operation guards and interrupted integration, inspect source/template pairs, and verify archival cannot claim delivery before integration.
