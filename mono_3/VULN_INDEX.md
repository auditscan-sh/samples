# mono_3 expected findings index

This file is the ground-truth cheat sheet for the TypeScript/Node target. Line
numbers are approximate; scan evidence is authoritative. "Chain" means the
tainted value crosses a function/file boundary.

| # | Area | File | CWE / class | Expected family |
|---|------|------|-------------|-----------------|
| 1 | SQLi (chain) | `packages/payments/src/repository.ts` | CWE-89 | WEB-SQLI |
| 2 | SQLi (direct) | `packages/auth/src/index.ts` | CWE-89 | WEB-SQLI |
| 3 | Second-order SQLi | `packages/payments/src/repository.ts` | CWE-89 | WEB-SQLI |
| 4 | NoSQL `$where` injection | `packages/catalog/src/search.ts` | CWE-943 | WEB-NOSQL |
| 5 | Command injection | `packages/admin/src/index.ts` | CWE-78 | WEB-CMDI |
| 6 | Command injection via preview | `packages/uploads/src/preview.ts` | CWE-78 | WEB-CMDI |
| 7 | Code injection / eval | `packages/admin/src/index.ts` | CWE-94 | WEB / custom |
| 8 | Reflected XSS | `packages/catalog/src/index.ts` | CWE-79 | WEB-XSS |
| 9 | DOM XSS | `packages/web/src/UserProfile.tsx` | CWE-79 | WEB-XSS |
| 10 | SSTI | `packages/catalog/src/index.ts` | CWE-94 | WEB-SSTI |
| 11 | SSRF | `packages/gateway/src/proxy.ts` | CWE-918 | WEB-SSRF |
| 12 | SSRF to cloud metadata | `packages/uploads/src/preview.ts` | CWE-918 | WEB-SSRF |
| 13 | Open redirect | `packages/gateway/src/routes.ts` | CWE-601 | WEB-OPENRED |
| 14 | Path traversal (read/write/delete) | `packages/uploads/src/storage.ts` | CWE-22 | WEB-PATHDEL |
| 15 | Zip Slip | `packages/uploads/src/archive.ts` | CWE-22 | WEB-ZIPSLIP |
| 16 | Insecure deserialization | `packages/admin/src/index.ts` | CWE-502 | WEB-DESER |
| 17 | Prototype pollution | `packages/webhooks/src/handler.ts` | CWE-1321 | custom |
| 18 | JWT `alg:none` acceptance | `packages/auth/src/jwt.ts` | CWE-347 | custom |
| 19 | JWT alg confusion | `packages/auth/src/jwt.ts` | CWE-347 | custom |
| 20 | Hardcoded JWT secret | `packages/auth/src/jwt.ts` | CWE-798 | secret |
| 21 | Hardcoded cloud/API keys | `packages/shared/src/config.ts`, `.env` | CWE-798 | secret |
| 22 | Fake private key committed | `packages/auth/keys/jwt_private.pem` | CWE-321 | secret |
| 23 | Weak password hash (MD5) | `packages/auth/src/password.ts` | CWE-328 | crypto |
| 24 | Insecure RNG reset token | `packages/auth/src/password.ts` | CWE-338 | crypto |
| 25 | Timing-unsafe compare | `packages/webhooks/src/verify.ts` | CWE-208 | crypto |
| 26 | TLS verification disabled | `packages/gateway/src/proxy.ts` | CWE-295 | crypto |
| 27 | Broken access control / IDOR | `packages/payments/src/refunds.ts` | CWE-639 | custom |
| 28 | Broken access control / IDOR | `packages/auth/src/index.ts` | CWE-639 | custom |
| 29 | Missing function-level authz | `packages/admin/src/index.ts` | CWE-862 | custom |
| 30 | Mass assignment | `packages/auth/src/index.ts`, `packages/webhooks/src/handler.ts` | CWE-915 | WEB-MASS |
| 31 | Business logic: client price | `packages/payments/src/checkout.ts` | CWE-840 | custom |
| 32 | Business logic: negative refund | `packages/payments/src/refunds.ts` | CWE-840 | custom |
| 33 | Race condition: coupon redemption | `packages/payments/src/checkout.ts` | CWE-362 | concurrency |
| 34 | CSRF on state-changing routes | `packages/gateway/src/routes.ts` | CWE-352 | WEB-CSRF |
| 35 | Permissive CORS + credentials | `packages/gateway/src/server.ts` | CWE-942 | config |
| 36 | Missing security headers | `packages/gateway/src/server.ts` | CWE-693 | config |
| 37 | Cookie without Secure/HttpOnly | `packages/auth/src/index.ts` | CWE-614 | config |
| 38 | Verbose error / stack leakage | `packages/gateway/src/server.ts` | CWE-209 | config |
| 39 | Sensitive data in logs | `packages/shared/src/logger.ts` | CWE-532 | config |
| 40 | Log injection | `packages/shared/src/logger.ts` | CWE-117 | config |
| 41 | ReDoS | `packages/shared/src/validation.ts` | CWE-1333 | WEB-REDOS |
| 42 | Unbounded request body | `packages/gateway/src/server.ts` | CWE-400 | WEB-DOSREAD |
| 43 | Host-header injection | `packages/gateway/src/routes.ts` | CWE-644 | WEB-HOSTHDR |
| 44 | Web cache poisoning | `packages/gateway/src/routes.ts` | CWE-524 | custom |
| 45 | HTTP request smuggling surface | `packages/gateway/src/routes.ts` | CWE-444 | custom |
| 46 | Unrestricted file upload | `packages/uploads/src/index.ts` | CWE-434 | custom |
| 47 | GraphQL introspection + no depth limit | `packages/catalog/src/graphql.ts` | CWE-200 | custom |
| 48 | Vulnerable npm pins | `packages/*/package.json` | CVE | DEP / VEX |
| 49 | Vulnerable Python tooling pins | `packages/legacy-tools/requirements.txt` | CVE | DEP / VEX |
| 50 | Unpinned git dependency | `packages/legacy-tools/requirements.txt` | CWE-1104 | DEP |
| 51 | Privileged/hostPath container | `k8s/gateway-deploy.yaml` | CWE-250 | infra |
| 52 | Over-provisioned replicas/resources | `k8s/payments-deploy.yaml` | cost | cost/waste |
| 53 | Public S3 + 0.0.0.0/0 + public RDS | `terraform/main.tf` | CWE-284 | infra |
| 54 | CI/CD secret exposure | `.github/workflows/deploy.yml` | CWE-522 | infra |
| 55 | Docker root + latest + baked secrets | `Dockerfile` | CWE-250 | infra |
| 56 | Dead service / dead helpers | `packages/ghost-service`, `src/legacy.ts` | dead code | DEAD |
| 57 | AI-slop verbose functions | `packages/shared/src/money.ts`, `packages/payments/src/tax.ts`, `packages/web/src/UserProfile.tsx` | waste | waste/dead |
| 58 | Hallucinated APIs / junk comments | `packages/shared/src/index.ts`, `packages/catalog/src/suggestions.ts` | waste | waste |

## Deliberate false-positive bait

These are intentionally **safe** and should not be reported as the matching
vulnerability:

- `packages/payments/src/repository.ts::getOrderParameterized` uses bound `?`
  placeholders.
- `packages/gateway/src/routes.ts::safeRedirect` allows only `/dashboard` and
  `/orders`.
- `packages/auth/src/tokens.ts::checksumEtag` uses MD5 for an ETag, not for
  passwords or signatures.
- `packages/uploads/src/storage.ts::readInsideUploadDir` resolves and verifies
  the path remains under the upload root.
- `packages/catalog/src/graphql.ts::resolvePublicProduct` has no user input.
