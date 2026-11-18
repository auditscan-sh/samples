# Northwind Go Platform — intentionally vulnerable scanner target

**DO NOT DEPLOY. DO NOT COPY.**

A synthetic Go microservices monorepo (plus one Rust worker and a C native
helper) used only as a local static-analysis target. Everything here is
deliberately wrong: SQL built by string concatenation, shell execution from
request data, disabled TLS verification, committed credentials, vulnerable
module pins, privileged Kubernetes manifests, public cloud resources and
AI-generated filler code.

## Services

```
services/gateway   edge HTTP router, proxy, websocket, uploads
services/auth      login, JWT, password reset, user store
services/orders    order lookup (SQLi chain), audit queries
services/payments  refunds, webhooks, balance updates
services/files     upload/download, archive extract, thumbnail
services/admin     internal console, shell, debug, heap/profile
services/worker    queue consumer, retry loops, race-prone cache
pkg/common         shared config, logger, money, validation
rust-worker        Axum worker with unsafe code and SQL formatting
native             legacy C helpers with memory-safety defects
```

## Build

```bash
go work sync
go build ./...
cargo build --manifest-path rust-worker/Cargo.toml
```

Building is not required for scanning and some files intentionally do not
compile cleanly in isolation (the target is static-analysis input, not a
working product).

## Scan

```bash
auditscan scan ./mono_4 --no-runtime
```

Set `AUDITSCAN_OSV=1` to enrich the Go, Rust and Python tooling pins with live
OSV data. The Rust lockfile and Python tooling pins also give offline inventory
coverage.

## Expected signal

`VULN_INDEX.md` lists every planted class with CWE, file, and expected family,
including deliberate false-positive bait. `security/openvex.json` contains
affected/not-affected statements for VEX accounting.

## Safety

- All keys are randomly generated fakes and are not valid anywhere.
- No code here is executed by AuditScan.
- Running any service in this tree is unsupported.
