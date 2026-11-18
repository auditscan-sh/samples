# Northwind Commerce — intentionally vulnerable scanner target

**DO NOT DEPLOY. DO NOT COPY. DO NOT USE ANY OF THIS CODE.**

This repository is a synthetic, deliberately insecure TypeScript/Node.js
monorepo used as a local target for static-analysis accuracy testing. Every
credential in this tree is fake, every vulnerability is planted on purpose, and
several files contain intentional "AI slop" (over-engineered trivial functions,
hallucinated APIs, duplicated logic, dead code) to exercise waste/dead-code
detectors.

The application is fictional: a small commerce platform with a gateway, auth,
payments, catalog, uploads, admin, webhooks and a React storefront.

## Why it exists

- Measure true-positive, false-positive and evidence quality for a scanner.
- Exercise web injection detection across TypeScript route -> service ->
  repository call chains (not only single-function regex hits).
- Exercise secret scanning, dependency/CVE/VEX, IaC, CI/CD, container and
  supply-chain checks.
- Provide deliberate false-positive bait (safe parameterized queries, safe
  crypto used for cache keys, allowlisted redirects) next to real defects.

## Layout

```
packages/gateway      Express edge: proxy, static, auth glue
packages/auth         login, JWT, password reset, user store
packages/payments     orders, refunds, checkout, tax
packages/catalog      search, NoSQL, GraphQL, suggestions
packages/uploads      file upload, archive extract, preview fetch
packages/admin        internal admin console, shell, debug
packages/webhooks     inbound webhook verification/dispatch
packages/web          React storefront
packages/shared       money/format/logger/config helpers
packages/ghost-service abandoned service (dead code)
packages/legacy-tools  Python release tooling with vulnerable pins
k8s/ terraform/ .github/  intentionally misconfigured delivery assets
```

## Scan

```bash
auditscan scan ./mono_3 --no-runtime
```

Supply-chain emulation can be enabled with `AUDITSCAN_OSV=1` if you want live
OSV enrichment; the npm pins are also in the bundled offline advisory table.

## Expected signal

See `VULN_INDEX.md`. It lists the planted issue, the CWE, and the file that
should carry evidence. Severity is intentionally mixed: critical issues sit next
to low-value dead code and noisy AI-generated helpers.

## Safety

- Secrets are structurally valid-looking but randomly generated and unusable.
- The private key fixture is not a valid key.
- No code here is executed by the scanner; running the app is unsupported.
- If you fork this, delete the credentials and the vulnerable code first.
