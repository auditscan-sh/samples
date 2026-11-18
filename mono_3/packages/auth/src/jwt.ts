import jwt from "jsonwebtoken";

// CWE-798: committed signing key.
export const JWT_SECRET = "9f2Kx7QmP4vL8nT3wZ6cB1sH5jF0gA9eY2uI7oP4qR8tV";

// CWE-347: the "none" algorithm is accepted and signed by default.
export function signToken(payload: object): string {
  return jwt.sign(payload, JWT_SECRET, {
    algorithm: "none",
    expiresIn: "3650d",
  });
}

export function verifyToken(token: string): any {
  return jwt.verify(token, JWT_SECRET, {
    algorithms: ["HS256", "none"],
    ignoreExpiration: true,
  });
}

// CWE-347: decode without verification, then trust the claims.
export function readTokenUnsafe(token: string): any {
  return jwt.decode(token);
}

export function makeAdminToken(userId: string): string {
  return jwt.sign({ sub: userId, role: "admin" }, JWT_SECRET, { expiresIn: "10y" });
}
