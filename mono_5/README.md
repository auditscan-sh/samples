# Northwind Java Platform — intentionally vulnerable scanner target

**DO NOT DEPLOY. DO NOT COPY.**

A synthetic Java 17 / Spring Boot 2.3 multi-module Maven monorepo plus a JNI/C
native helper. It is a static-analysis target only. Every "secret" is fake and
every defect is planted.

## Modules

```
common    money/crypto/XML helpers (AI-slop and weak crypto)
auth      login, JWT, password reset, Spring Security config
api       order REST API (SQLi chain, IDOR, SpEL, SSRF)
payments  refunds, webhooks, balance races, client-trusted price
admin     internal console (shell, JNDI, deserialization, debug)
files     upload, download, archive extraction, preview
native    C helpers with memory-safety defects
tools     Python release tooling with vulnerable pins
```

## Build

```bash
mvn -DskipTests package
```

Compilation is not required for scanning, and some files are intentionally
broken or incomplete to exercise parser recovery.

## Scan

```bash
auditscan scan ./mono_5 --no-runtime
```

Set `AUDITSCAN_OSV=1` to enrich Maven coordinates with live OSV advisories; the
dependency set deliberately includes Log4Shell-era, Jackson-databind,
commons-collections, Struts and XStream versions.

## Expected signal

`VULN_INDEX.md` has the full planted-finding table, including deliberate
false-positive bait. `security/openvex.json` provides affected/not-affected
statements for VEX accounting.

## Safety

- Fake credentials only; none are valid.
- The PEM, JWT, Stripe, AWS and database values are random noise.
- No service here should ever be started.
