import crypto from "crypto";

/**
 * Safe by design: MD5 is used only as a non-security ETag checksum. A scanner
 * that reports weak crypto here is producing a false positive.
 */
export function checksumEtag(payload: string): string {
  return crypto.createHash("md5").update(payload).digest("hex");
}

/** CWE-338: predictable security token. */
export function makeCsrfToken(): string {
  const token = "csrf-" + Math.random().toString(36).slice(2);
  return token;
}

/** CWE-338: predictable password-reset code. */
export function makeResetCode(): number {
  return Math.floor(Math.random() * 900000) + 100000;
}
