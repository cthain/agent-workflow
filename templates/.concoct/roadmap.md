---
version: 1
project: <project-name>
updated: YYYY-MM-DD
---

# Roadmap

## Purpose

This roadmap defines the intended future evolution of `<project-name>`.

It records planned product outcomes, their priority, dependencies, readiness, and delivery status.

This file is distinct from:

- `.concoct/capabilities.md`, which records what the product can do now;
- `.concoct/current/task-plan.md`, which records the active implementation task;
- `.concoct/archive/`, which preserves completed task history.

The Product Owner owns this file.

## Product direction

Briefly describe the product’s intended direction.

Include:

- who the product is for;
- what problem it solves;
- what it should become good at;
- important product principles;
- major constraints.

Example:

```text
<project-name> helps <target users> accomplish <primary outcome> by <core approach>.

The product should prioritize:
- <principle>
- <principle>
- <principle>
```

## Roadmap conventions

Each roadmap item has a stable identifier.

Recommended identifier format:

```text
<PROJECT-PREFIX>-NNN
```

Example:

```text
APP-001
APP-002
APP-003
```

Do not renumber or reuse identifiers.

### Statuses

- `candidate` — plausible future work that is not yet ready or committed for planning;
- `planned` — sufficiently defined and eligible for task planning;
- `active` — represented by the current active task plan;
- `blocked` — cannot proceed until a dependency or decision is resolved;
- `delivered` — transitional marker pending Product Owner reconciliation and
  removal from the active roadmap;
- `deferred` — still valid but intentionally postponed;
- `cancelled` — no longer intended and awaiting reference/rationale
  reconciliation before removal.

### Priorities

- `critical`
- `high`
- `medium`
- `low`

Priority describes importance.

Dependencies and readiness determine implementation order.

`Depends on` contains only unresolved delivery dependencies on outstanding
roadmap items. `Capability prerequisites` contains stable references to
accepted current behavior in `.concoct/capabilities.md`. Satisfied sequencing
constraints and delivery provenance belong in neither field.

After relationships are reconciled, remove delivered and cancelled items from
the active roadmap. Keep their identifiers reserved; capabilities and archives
preserve accepted delivery provenance.

### Capability impact

Each roadmap item should state whether it is expected to:

- add a capability;
- update an existing capability;
- remove a capability;
- leave observable capabilities unchanged.

Expected capability impact is planning intent. Delivered capability impact is validated during review and archival.

## Planning readiness

A roadmap item is ready to move from `candidate` to `planned` when:

- the desired product outcome is clear;
- the current capability baseline is understood;
- scope boundaries are clear;
- important dependencies are identified;
- major product decisions are resolved;
- acceptance criteria are observable;
- the item is coherent enough for one implementation and review cycle.

When ready, use:

```text
concoct plan <roadmap-id>
```

---

## <PROJECT-PREFIX>-001 — <Outcome-oriented title>

- Status: `candidate`
- Priority: `high`
- Depends on: `none`
- Capability prerequisites: `none`
- Capability impact: `add | update | remove | none`

### Outcome

Describe the product outcome.

Focus on what will become possible or meaningfully different.

Avoid implementation steps.

### Rationale

Explain:

- what problem this solves;
- who benefits;
- why it matters now;
- why this belongs on the roadmap.

### Current state

Describe the relevant current capability or limitation.

Reference `.concoct/capabilities.md` where applicable.

### Target state

Describe the expected product behavior after delivery.

### Requirements

- Requirement
- Requirement
- Requirement

### Non-goals

- Explicitly excluded scope
- Related work that belongs elsewhere
- Tempting implementation or product expansion that is not part of this item

### Dependencies

Describe real dependencies, decisions, or external constraints.

Use `none` when the item is independently planable.

### Acceptance criteria

- Observable product outcome
- Observable product outcome
- Observable product outcome

### Product decisions

Record unresolved product decisions that must be resolved before planning.

Use:

```text
None.
```

when no decisions remain.

### Delivery

- Active task:
- Archive:
- Delivered capabilities:

---

## <PROJECT-PREFIX>-002 — <Outcome-oriented title>

- Status: `candidate`
- Priority: `medium`
- Depends on: `<PROJECT-PREFIX>-001`
- Capability prerequisites: `CAP-NNN | none`
- Capability impact: `add | update | remove | none`

### Outcome

Describe the intended product outcome.

### Rationale

Explain why the outcome matters.

### Current state

Describe the relevant current state.

### Target state

Describe the desired accepted behavior.

### Requirements

- Requirement
- Requirement

### Non-goals

- Non-goal
- Non-goal

### Dependencies

Explain each unresolved delivery dependency and why it is necessary. Explain
separately how each capability prerequisite is used.

### Acceptance criteria

- Observable completion condition
- Observable completion condition

### Product decisions

- Decision still required

### Delivery

- Active task:
- Archive:
- Delivered capabilities:

---

## Candidate ideas

Use this section sparingly for ideas that are directionally relevant but not yet understood well enough to become full roadmap items.

Do not use it as an unrestricted backlog.

### <Short idea title>

- Product need:
- Why it may matter:
- Missing information:
- Decision needed:
- Related roadmap items or capabilities:

## Deferred work

Preserve deliberately postponed work here or in its original roadmap item.

Record:

- why it was deferred;
- what may cause reconsideration;
- relevant dependencies or constraints.

## Cancelled work

Preserve cancellation rationale while references are reconciled, then remove
the item from the active roadmap. Keep its identifier in a compact reserved-ID
record and never reuse it.

## Recommended next work

List the roadmap items currently most appropriate for task planning.

Example:

```text
1. APP-001 — <title>
2. APP-004 — <title>
```

Explain briefly why each item is ready and appropriately sequenced.

## Product Owner review checklist

Before completing a roadmap update, confirm:

- `capabilities.md` was read;
- delivered work is not presented as future work;
- roadmap items describe outcomes rather than implementation steps;
- stable identifiers are preserved;
- priorities and dependencies are explicit;
- planned items are actually ready for task planning;
- blocked items name their blockers;
- deferred items preserve rationale and removed cancelled identifiers remain
  reserved;
- product ambiguity is visible rather than pushed into implementation;
- material changes and recommended next steps are summarized.
