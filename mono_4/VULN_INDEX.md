# mono_4 expected findings index

| # | Area | File | CWE / class | Expected family |
|---|------|------|-------------|-----------------|
| 1 | SQLi (cross-file chain) | `services/orders/repository.go` | CWE-89 | WEB-SQLI |
| 2 | SQLi (auth login chain) | `services/auth/store.go` | CWE-89 | WEB-SQLI |
| 3 | SQLi (ORDER BY) | `services/orders/repository.go` | CWE-89 | WEB-SQLI |
| 4 | SQLi (admin query) | `services/admin/handlers.go` | CWE-89 | WEB-SQLI |
| 5 | SQLi (Rust) | `rust-worker/src/db.rs` | CWE-89 | WEB-SQLI |
| 6 | NoSQL `$where` | `services/orders/handlers.go` | CWE-943 | WEB-NOSQL |
| 7 | Command injection | `services/admin/exec.go` | CWE-78 | WEB-CMDI |
| 8 | Command injection | `services/files/thumbnail.go` | CWE-78 | WEB-CMDI |
| 9 | Command injection (Rust) | `rust-worker/src/main.rs` | CWE-78 | WEB-CMDI |
| 10 | Command injection (C) | `native/legacy.c` | CWE-78 | custom |
| 11 | SSRF | `services/gateway/proxy.go` | CWE-918 | WEB-SSRF |
| 12 | SSRF (Rust) | `rust-worker/src/main.rs` | CWE-918 | WEB-SSRF |
| 13 | Open redirect | `services/gateway/handlers.go` | CWE-601 | WEB-OPENRED |
| 14 | Path traversal | `services/files/storage.go` | CWE-22 | WEB-PATHDEL |
| 15 | Path traversal | `services/admin/handlers.go` | CWE-22 | WEB-PATHDEL |
| 16 | Zip Slip | `services/files/archive.go` | CWE-22 | WEB-ZIPSLIP |
| 17 | LDAP injection | `services/auth/ldap.go` | CWE-90 | WEB-LDAP |
| 18 | XPath/XML injection | `services/orders/xml.go` | CWE-643 | WEB-XPATH |
| 19 | XXE | `services/orders/xml.go` | CWE-611 | WEB-XXE |
| 20 | Insecure deserialization | `services/worker/consumer.go` | CWE-502 | WEB-DESER |
| 21 | JWT `alg:none` | `services/auth/jwt.go` | CWE-347 | custom |
| 22 | JWT parser without method check | `services/auth/jwt.go` | CWE-347 | custom |
| 23 | Hardcoded JWT secret | `services/auth/jwt.go` | CWE-798 | secret |
| 24 | Hardcoded cloud/API keys | `pkg/common/config.go`, `.env` | CWE-798 | secret |
| 25 | Weak password hash (MD5) | `services/auth/password.go` | CWE-328 | crypto |
| 26 | Predictable reset token | `services/auth/password.go` | CWE-338 | crypto |
| 27 | Timing-unsafe compare | `services/auth/password.go` | CWE-208 | crypto |
| 28 | TLS verification disabled | `services/gateway/proxy.go` | CWE-295 | crypto |
| 29 | Weak crypto (DES/RC4/MD5) | `pkg/common/crypto.go` | CWE-327 | crypto |
| 30 | IDOR / BOLA | `services/payments/refunds.go` | CWE-639 | custom |
| 31 | Missing authz on admin routes | `services/admin/main.go` | CWE-862 | custom |
| 32 | Auth bypass via header | `services/gateway/middleware.go` | CWE-290 | custom |
| 33 | Mass assignment | `services/auth/handlers.go` | CWE-915 | WEB-MASS |
| 34 | Business logic: negative refund | `services/payments/refunds.go` | CWE-840 | custom |
| 35 | Race condition / TOCTOU | `services/payments/refunds.go` | CWE-362 | concurrency |
| 36 | Race: unsynchronized map | `services/worker/cache.go` | CWE-362 | concurrency |
| 37 | Unbounded body read | `services/gateway/handlers.go` | CWE-400 | WEB-DOSREAD |
| 38 | ReDoS | `pkg/common/validate.go` | CWE-1333 | WEB-REDOS |
| 39 | CORS wildcard + credentials | `services/gateway/main.go` | CWE-942 | config |
| 40 | Missing security headers | `services/gateway/main.go` | CWE-693 | config |
| 41 | Verbose error/env dump | `services/admin/debug.go` | CWE-209 | config |
| 42 | Secrets in logs | `pkg/common/logger.go` | CWE-532 | config |
| 43 | Log injection | `pkg/common/logger.go` | CWE-117 | config |
| 44 | Insecure cookie | `services/auth/handlers.go` | CWE-614 | config |
| 45 | Weak randomness in security context | `services/auth/password.go` | CWE-338 | crypto |
| 46 | Unsafe pointer / memory corruption | `services/worker/consumer.go` | CWE-787 | custom |
| 47 | Buffer overflow / format string | `native/legacy.c` | CWE-120/CWE-134 | custom |
| 48 | Use-after-free | `native/legacy.c` | CWE-416 | custom |
| 49 | Vulnerable Go module pins | `services/*/go.mod` | CVE | DEP / VEX |
| 50 | Vulnerable Rust crate pins | `rust-worker/Cargo.lock` | CVE | DEP / VEX |
| 51 | Vulnerable Python tooling pins | `tools/requirements.txt` | CVE | DEP / VEX |
| 52 | Unpinned git dependency | `tools/requirements.txt` | CWE-1104 | DEP |
| 53 | Privileged + hostPath + hostNetwork | `k8s/gateway-deploy.yaml` | CWE-250 | infra |
| 54 | Over-provisioned replicas/resources | `k8s/payments-deploy.yaml` | cost | cost/waste |
| 55 | Public S3 / public RDS / 0.0.0.0/0 | `terraform/main.tf` | CWE-284 | infra |
| 56 | CI secret exposure + PR code execution | `.github/workflows/ci.yml` | CWE-522 | infra |
| 57 | Docker root + latest + baked secrets | `Dockerfile` | CWE-250 | infra |
| 58 | Dead service / dead helpers | `services/worker/legacy.go`, `rust-worker/src/legacy.rs` | dead code | DEAD |
| 59 | AI-slop verbose functions | `pkg/common/money.go`, `services/payments/retry.go`, `rust-worker/src/legacy.rs` | waste | waste/dead |

## Deliberate false-positive bait (safe)

- `services/orders/repository.go::FindOrdersParameterized` uses `?` binding.
- `services/gateway/proxy.go::SafeFetch` validates host against an allowlist.
- `services/files/storage.go::ReadInsideUploadDir` resolves and checks the
  prefix.
- `pkg/common/crypto.go::ETag` uses MD5 for a cache key, not for passwords.
- `services/auth/password.go::CheckPasswordBcrypt` uses bcrypt correctly.
