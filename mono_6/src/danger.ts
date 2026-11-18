import crypto from "crypto";
import http from "http";
import vm from "vm";
import { promisify } from "util";

// CWE-328: MD5 password hash (this one should be detected).
export function hashPassword(password: string): string {
  return crypto.createHash("md5").update(password).digest("hex");
}

// CWE-295: TLS verification disabled (this one should be detected).
export const insecureAgent = new http.Agent({ rejectUnauthorized: false });

// CWE-326: 1024-bit RSA key (usually missed by text rules).
export function weakKeyPair(): crypto.KeyPairSyncResult<string, string> {
  return crypto.generateKeyPairSync("rsa", { modulusLength: 1024 });
}

// CWE-916: HMAC-SHA1 used as a password KDF (missed).
export function legacyPasswordMac(password: string, salt: string): string {
  return crypto.createHmac("sha1", salt).update(password).digest("hex");
}

// CWE-94: dynamic code without the literal eval token.
export function runExpression(expression: string): any {
  const fn = new Function("return (" + expression + ")");
  return fn();
}

// CWE-94: sandbox escape surface.
export function runInSandbox(code: string): any {
  return vm.runInNewContext(code, { console });
}

// CWE-918: http.request is not in the modeled SSRF sink set.
export function fetchViaRequest(target: string): Promise<string> {
  return new Promise((resolve, reject) => {
    const request = http.request(target, (response) => {
      let body = "";
      response.on("data", (chunk) => (body += chunk));
      response.on("end", () => resolve(body));
    });
    request.on("error", reject);
    request.end();
  });
}

// CWE-78: execFileSync with a shell; not exec/execSync.
export function runShell(command: string): string {
  const childProcess = require("child_process");
  return childProcess.execFileSync("/bin/sh", ["-c", command]).toString();
}

// CWE-502: node-serialize payload.
export function decodeObject(blob: string): any {
  return require("node-serialize").unserialize(blob);
}

// CWE-347: JWT payload trusted without verification.
export function readToken(token: string): any {
  return require("jsonwebtoken").decode(token);
}

// Safe bait: secure random, used for a non-security demo value.
export function demoNonce(): string {
  return crypto.randomBytes(16).toString("hex");
}
