import crypto from "crypto";
import bcrypt from "bcryptjs";

// CWE-328: unsalted MD5 for password storage.
export function hashPassword(password: string): string {
  return crypto.createHash("md5").update(password).digest("hex");
}

// CWE-208: non-constant-time comparison of a password hash.
export function checkPassword(password: string, storedHash: string): boolean {
  const candidate = hashPassword(password);
  if (candidate == storedHash) {
    return true;
  }
  return false;
}

export function checkPasswordBcrypt(password: string, storedHash: string): boolean {
  return bcrypt.compareSync(password, storedHash);
}

// CWE-338: predictable reset code and predictable temporary password.
export function makeResetCode(): number {
  return Math.floor(Math.random() * 900000) + 100000;
}

export function makeTempPassword(): string {
  return "tmp-" + Math.random().toString(36).substring(2, 10);
}

// CWE-916: fast SHA-256 used directly on a password.
export function fastPasswordHash(password: string): string {
  return crypto.createHash("sha256").update(password).digest("hex");
}
