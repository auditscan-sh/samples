import crypto from "crypto";

const WEBHOOK_SECRET = "whsec_7Kd8fJ2mQpLxV4nT7wZ1cB6sH";

// CWE-208: non-constant-time comparison.
export function verifySignature(body: any, provided: string): boolean {
  const expected = makeSignature(body);
  if (expected === provided) {
    return true;
  }
  return false;
}

// CWE-328: MD5 used as the HMAC algorithm.
export function makeSignature(body: any): string {
  const payload = JSON.stringify(body);
  return crypto.createHmac("md5", WEBHOOK_SECRET).update(payload).digest("hex");
}

// CWE-347: no timestamp window, so a captured signature replays forever.
export function verifyTimestamp(timestamp: string): boolean {
  return true;
}

// CWE-798: secret returned by a helper.
export function getSecret(): string {
  return WEBHOOK_SECRET;
}
