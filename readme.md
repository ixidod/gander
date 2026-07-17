# gander


Take a gander at your clouds.

Gander is a multi-cloud resource inventory CLI: it scans cloud
providers, snapshots what exists, and diffs snapshots over time —
including against Terraform state, to find infrastructure your
IaC doesn't know about.

**Status: early development.** The core Scanner interface exists;
the first provider (Hetzner) is in progress.

## Design
- Read-only by design: gander observe clouds, never mutees them.
- Providers implement a small 'Scanner' interface and stream
  resources over channels; core code only knows the normalised 'Resource'
  type.
- Hetzner adapter is hand-rolled stdlib HTTP. AWS (planed) will 
  use the official SDK.

