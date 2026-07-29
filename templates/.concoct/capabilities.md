---
version: 1
project: <project-name>
updated: YYYY-MM-DD
---

# Capabilities

## Purpose

This file is the canonical human-readable record of what `<project-name>` can do now.

It describes current, accepted product behavior.

It is intended to give humans and agents a reliable product baseline without requiring them to infer the complete product surface from source code alone.

This file is distinct from:

- `.concoct/roadmap.md`, which records intended future work;
- `.concoct/current/task-plan.md`, which records the active implementation task;
- `.concoct/archive/`, which preserves completed task history;
- `CHANGELOG.md`, which records changes over time;
- implementation documentation, which explains how the system works internally.

The Archivist owns updates to this file after approved work is archived.

## Capability principles

### Describe current truth

Only record behavior that currently exists and has been accepted.

Do not record:

- planned work;
- partially implemented behavior;
- unapproved changes;
- ideas;
- aspirations;
- historical task details;
- speculative future support.

### Describe observable behavior

Capabilities should explain what the product enables, not how it is implemented.

Good:

```text
The application can export reports in CSV format.
```

Too implementation-specific:

```text
The report service calls `encoding/csv` from the export handler.
```

Include implementation details only when they materially constrain how the capability is used.

### Preserve stable identifiers

Each capability has a stable identifier.

Recommended format:

```text
CAP-NNN
```

Examples:

```text
CAP-001
CAP-002
CAP-003
```

Do not renumber or reuse capability identifiers.

### Keep capability scope coherent

One capability should represent one independently understandable product behavior or closely related capability group.

Split a capability when:

- different parts can exist independently;
- different audiences use them differently;
- limitations or lifecycle differ materially;
- future roadmap work is likely to affect only one part.

Do not split capabilities into trivial implementation fragments.

### Record meaningful limitations

A capability description should include limitations that affect users, operators, integrators, or future planning.

Do not clutter entries with incidental implementation constraints.

### Preserve traceability

Where practical, reference:

- the roadmap item that delivered or changed the capability;
- the archive summary that records acceptance;
- relevant user or API documentation.

Traceability should support archaeology without turning this file into a changelog.

## Capability status

Capabilities in this file are normally current and available.

When the schema needs an explicit status, use:

- `active` — currently available;
- `deprecated` — still available but expected to be removed or replaced;
- `limited` — available with a material limitation;
- `removed` — no longer available, retained only for traceability.

Prefer removing obsolete entries from the current capability sections only when project conventions preserve removal history elsewhere.

Do not silently erase capability history when other artifacts reference its identifier.

---

## CAP-001 — <Capability title>

- Status: `active`
- Audience: `<user | operator | administrator | developer | integrator | multiple>`
- Added by: `<ROADMAP-ID>`
- Archive: `.concoct/archive/YYYY-MM-DD-roadmap-id-short-task-name/`
- Documentation: `<path or none>`

### Capability

Describe what the product can currently do.

Use present tense.

Example:

```text
The product can initialize a new local project with the standard repository instructions, workflow directories, and agent-specific adapter files.
```

### User value

Explain why this capability matters and who benefits.

### Inputs

Describe the meaningful inputs required to use the capability.

Use:

```text
None.
```

when there are no user-supplied inputs.

### Outputs and effects

Describe observable outputs, state changes, or side effects.

### Limitations

- Material limitation
- Unsupported case
- Important operational boundary

Use:

```text
None currently documented.
```

when no meaningful limitations are known.

### Verification evidence

Reference evidence that supports the current capability claim.

Examples:

- automated tests;
- end-to-end tests;
- command examples;
- approved review;
- documentation;
- archive summary.

### Related capabilities

- `CAP-XXX` — relationship
- `CAP-YYY` — relationship

Use:

```text
None.
```

when the capability is independent.

---

## CAP-002 — <Capability title>

- Status: `active`
- Audience: `<audience>`
- Added by: `<ROADMAP-ID>`
- Archive: `.concoct/archive/YYYY-MM-DD-roadmap-id-short-task-name/`
- Documentation: `<path or none>`

### Capability

Describe the current accepted behavior.

### User value

Explain the value.

### Inputs

Describe required inputs.

### Outputs and effects

Describe observable results.

### Limitations

- Limitation

### Verification evidence

- Evidence

### Related capabilities

- Related capability

---

## Deprecated capabilities

Use this section for capabilities that still exist but should no longer be used for new work.

### CAP-XXX — <Deprecated capability title>

- Status: `deprecated`
- Deprecated by: `<ROADMAP-ID>`
- Replacement: `CAP-YYY | none`
- Planned removal: `<roadmap-id | unknown>`

#### Current behavior

Describe what still works.

#### Deprecation reason

Explain why the capability is deprecated.

#### Migration guidance

Explain what users should use instead.

#### Limitations

Describe any additional restrictions during deprecation.

## Removed capabilities

Use this section only when the project keeps removed capabilities here for traceability.

Otherwise, preserve removal history in the archive and changelog.

### CAP-XXX — <Removed capability title>

- Status: `removed`
- Removed by: `<ROADMAP-ID>`
- Archive: `.concoct/archive/YYYY-MM-DD-roadmap-id-short-task-name/`
- Replacement: `CAP-YYY | none`

#### Former behavior

Describe what the capability previously enabled.

#### Removal reason

Explain why it was removed.

#### Migration or replacement

Describe the supported alternative.

## Capability relationships

Use this section only when relationships are important enough to summarize globally.

Examples:

```text
CAP-001 provides the project initialization required before CAP-004 can manage active workflow transitions.

CAP-006 extends CAP-003 by supporting review-remediation mode.
```

Avoid duplicating relationships already clear in individual capability entries.

## Known capability gaps

Use this section sparingly.

A capability gap describes an important absence in the current product baseline.

It is not automatically a roadmap commitment.

### <Gap title>

- Current limitation:
- Affected users:
- Related capabilities:
- Related roadmap item:
- Notes:

If the gap is accepted as future work, reference the relevant roadmap item.

Do not use this section as an unrestricted backlog.

## Archivist update workflow

When approved work is archived:

1. Read `AGENTS.md`.
2. Read `.concoct/personas/archivist.md`.
3. Read the completed task plan, notes, and reviews.
4. Inspect the delivered implementation and verification evidence.
5. Determine whether capabilities were:
   - added;
   - updated;
   - removed;
   - unaffected.
6. Update only the capability truth supported by the accepted outcome.
7. Preserve stable capability identifiers.
8. Add archive and roadmap traceability.
9. Confirm future or incomplete work is not presented as current capability.
10. Update the `updated` date.

## Capability impact examples

### Add

```yaml
capability-impact:
  type: add
  ids:
    - CAP-004
```

Create a new capability entry based on delivered behavior.

### Update

```yaml
capability-impact:
  type: update
  ids:
    - CAP-002
```

Revise the existing capability to reflect the new accepted behavior.

Preserve the identifier.

### Remove

```yaml
capability-impact:
  type: remove
  ids:
    - CAP-003
```

Remove the behavior from current capability truth while preserving required traceability.

### None

```yaml
capability-impact:
  type: none
  rationale: Internal refactor with no observable product behavior change.
```

Do not make a capability edit merely to show that archival occurred.

## Capability review checklist

Before completing an update, confirm:

- every listed capability exists now;
- capability descriptions use present tense;
- planned behavior is excluded;
- claims are observable and evidence-based;
- stable identifiers are preserved;
- limitations are meaningful and current;
- capability impact matches the approved implementation;
- roadmap and archive references are valid;
- deprecated and removed capabilities are handled deliberately;
- current truth is not mixed with task history;
- the `updated` date is current.
