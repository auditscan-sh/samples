# mono_6 — AuditScan evasion probe

**DO NOT DEPLOY.** Seven files, all deliberately hostile. This target is a
teardown probe, not a realistic app: it is built around what static scanners
commonly miss (alternate frameworks, aliases, lockfile-only versions, encoded
secrets, indirect sinks).

## Planted issues (expected ground truth)

1. **Koa taint sources** — `ctx.query` / `ctx.request.body` are not the
   Express-style `req.query` / `req.body`. SQLi, SSRF, SSTI, open redirect and
   command execution all receive request data through Koa context objects.
2. **Lockfile-only CVEs** — `package.json` uses ranges/aliases
   (`^4.17.20`, `npm:minimist@1.2.5`, `0.21.x`), while the exact vulnerable
   versions (`lodash 4.17.20`, `minimist 1.2.5`, `axios 0.21.1`) exist only in
   `package-lock.json`.
3. **Command execution through `execFileSync`** — not `exec`/`execSync`.
4. **SSRF through `http.request`** — not `http.get`/`fetch`.
5. **Code execution through `new Function` and `vm.runInNewContext`** — not a
   literal `eval(...)` call.
6. **Deserialization through `node-serialize` and `jwt.decode`** — the scanner
   looks for pickle/YAML/marshal patterns.
7. **Path traversal through `fs.createReadStream`** — no delete/unlink call.
8. **Base64-encoded credentials** — AWS key, JWT secret, database URL and a
   test-mode Stripe key live in `config/settings.json`, none in provider-prefix
   form.
9. **Secrets in documentation** — this README contains a fake GitHub token and
   a fake PEM block, which are outside the code/config scan surface.
10. **Obvious controls that should still be found** — MD5 password hashing,
    `rejectUnauthorized: false`, and a plaintext database password in
    `src/db.ts` exist as a sanity check.

## Fake credentials present in this file (documentation-only)

- GitHub: `ghp_Q7mK2pLxV4nT7wZ1cB6sH3jF5gA0eY2uI9oP`
- AWS: `AKIAQ7M2X9PL4RT8VW3N`

```
-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEA7QmP4vL8nT3wZ6cB1sH5jF0gA9eY2uI7oP4qR8tVxK2mQ9pLxV
THIS-IS-NOT-A-REAL-KEY-DO-NOT-USE
-----END RSA PRIVATE KEY-----
```

## Expected scanner result

The goal is a deliberately lopsided report: a handful of crypto findings from
`src/danger.ts`, and silence for the Koa request paths, lockfile CVEs, encoded
secrets and documentation secrets.
