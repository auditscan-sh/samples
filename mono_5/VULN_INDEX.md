# mono_5 expected findings index

| # | Area | File | CWE / class | Expected family |
|---|------|------|-------------|-----------------|
| 1 | SQLi (controller -> service -> repo) | `api/.../OrderRepository.java` | CWE-89 | WEB-SQLI |
| 2 | SQLi (auth login chain) | `auth/.../UserRepository.java` | CWE-89 | WEB-SQLI |
| 3 | SQLi (ORDER BY) | `api/.../OrderRepository.java` | CWE-89 | WEB-SQLI |
| 4 | SQLi (admin audit) | `admin/.../AdminController.java` | CWE-89 | WEB-SQLI |
| 5 | SQLi (payments search) | `payments/.../PaymentRepository.java` | CWE-89 | WEB-SQLI |
| 6 | NoSQL `$where` | `api/.../OrderService.java` | CWE-943 | WEB-NOSQL |
| 7 | Command injection | `admin/.../CommandService.java` | CWE-78 | WEB-CMDI |
| 8 | Command injection | `files/.../StorageService.java` | CWE-78 | WEB-CMDI |
| 9 | SpEL injection | `admin/.../AdminController.java` | CWE-94 | WEB-SSTI |
| 10 | JNDI injection | `admin/.../AdminController.java` | CWE-502/CWE-917 | custom |
| 11 | Insecure deserialization | `admin/.../DeserializationService.java` | CWE-502 | WEB-DESER |
| 12 | XXE | `common/.../LegacyXml.java` | CWE-611 | WEB-XXE |
| 13 | SSRF | `api/.../OrderService.java` | CWE-918 | WEB-SSRF |
| 14 | SSRF (payment receipt) | `payments/.../PaymentController.java` | CWE-918 | WEB-SSRF |
| 15 | Open redirect | `auth/.../AuthController.java` | CWE-601 | WEB-OPENRED |
| 16 | Path traversal | `files/.../StorageService.java` | CWE-22 | WEB-PATHDEL |
| 17 | Zip Slip | `files/.../ArchiveService.java` | CWE-22 | WEB-ZIPSLIP |
| 18 | LDAP injection | `auth/.../AuthController.java` | CWE-90 | WEB-LDAP |
| 19 | XPath injection | `api/.../OrderService.java` | CWE-643 | WEB-XPATH |
| 20 | Reflected XSS | `api/.../OrderController.java` | CWE-79 | WEB-XSS |
| 21 | JWT `alg:none` + unsecured parser | `auth/.../JwtService.java` | CWE-347 | custom |
| 22 | Hardcoded JWT secret | `auth/.../JwtService.java` | CWE-798 | secret |
| 23 | Hardcoded cloud/API keys | `common/.../CryptoUtils.java`, `.env` | CWE-798 | secret |
| 24 | Weak password hash (MD5) | `auth/.../PasswordService.java` | CWE-328 | crypto |
| 25 | Predictable reset token | `auth/.../PasswordService.java` | CWE-338 | crypto |
| 26 | Timing-unsafe comparison | `auth/.../PasswordService.java` | CWE-208 | crypto |
| 27 | DES/ECB + RSA-1024 + Random | `common/.../CryptoUtils.java` | CWE-327 | crypto |
| 28 | TLS trust-all | `common/.../CryptoUtils.java` | CWE-295 | crypto |
| 29 | CSRF disabled | `auth/.../SecurityConfig.java` | CWE-352 | WEB-CSRF |
| 30 | CORS wildcard + credentials | `auth/.../SecurityConfig.java` | CWE-942 | config |
| 31 | IDOR / BOLA (refund) | `payments/.../RefundService.java` | CWE-639 | custom |
| 32 | IDOR (user) | `auth/.../AuthController.java` | CWE-639 | custom |
| 33 | Mass assignment | `auth/.../AuthController.java` | CWE-915 | WEB-MASS |
| 34 | Business logic: client price | `payments/.../PaymentController.java` | CWE-840 | custom |
| 35 | Business logic: negative refund | `payments/.../RefundService.java` | CWE-840 | custom |
| 36 | Race / double-spend | `payments/.../BalanceService.java` | CWE-362 | concurrency |
| 37 | Unbounded upload / read | `files/.../FileController.java` | CWE-400 | WEB-DOSREAD |
| 38 | Sensitive data in logs | `auth/.../AuthController.java` | CWE-532 | config |
| 39 | Verbose errors + env endpoint | `api/.../GlobalExceptionHandler.java` | CWE-209 | config |
| 40 | Actuator exposed with secrets | `api/src/main/resources/application.yml` | CWE-200 | config |
| 41 | Host-header injection | `auth/.../AuthController.java` | CWE-644 | WEB-HOSTHDR |
| 42 | Reflected XSS (error page) | `api/.../OrderController.java` | CWE-79 | WEB-XSS |
| 43 | Weak random session id | `auth/.../PasswordService.java` | CWE-338 | crypto |
| 44 | Buffer overflow / format string | `native/buffer.c` | CWE-120/CWE-134 | custom |
| 45 | Use-after-free / double free | `native/buffer.c` | CWE-416/CWE-415 | custom |
| 46 | Vulnerable Maven pins | `pom.xml` | CVE | DEP / VEX |
| 47 | Vulnerable Python tooling pins | `tools/requirements.txt` | CVE | DEP / VEX |
| 48 | Unpinned git dependency | `tools/requirements.txt` | CWE-1104 | DEP |
| 49 | Privileged + hostPath container | `k8s/api-deploy.yaml` | CWE-250 | infra |
| 50 | Over-provisioned replicas/resources | `k8s/payments-deploy.yaml` | cost | cost/waste |
| 51 | Public S3 / public RDS / 0.0.0.0/0 | `terraform/main.tf` | CWE-284 | infra |
| 52 | CI secret exposure + PR code execution | `.github/workflows/build.yml` | CWE-522 | infra |
| 53 | Docker root + latest + baked secrets | `Dockerfile` | CWE-250 | infra |
| 54 | Dead helpers / abandoned classes | `common/.../MoneyUtils.java` | dead code | DEAD |
| 55 | AI-slop verbose functions | `common/.../MoneyUtils.java`, `api/.../OrderService.java`, `payments/.../RefundService.java` | waste | waste/dead |

## Deliberate false-positive bait (safe)

- `api/.../OrderRepository.java::findByIdParameterized` uses `?` binding.
- `api/.../OrderService.java::safeRedirectTarget` allowlists local paths.
- `common/.../CryptoUtils.java::etag` uses MD5 for an ETag only.
- `auth/.../PasswordService.java::verifyBcrypt` uses BCrypt correctly.
- `files/.../StorageService.java::readInsideUploadDir` resolves and checks the
  prefix.
- `common/.../LegacyXml.java::parseSafe` disables DTDs and external entities.
