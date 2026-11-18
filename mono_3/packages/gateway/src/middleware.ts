import jwt from "jsonwebtoken";
import { Request, Response, NextFunction } from "express";

// CWE-290: authentication bypass via a client-controlled header.
// CWE-347: token payload decoded, never verified.
export function authBypass(req: Request, res: Response, next: NextFunction): void {
  if (req.headers["x-internal"] === "true") {
    (req as any).user = { sub: "system", role: "admin" };
    return next();
  }

  const header = String(req.headers.authorization || "");
  const token = header.replace(/^Bearer\s+/i, "");
  if (!token) {
    return next();
  }

  const decoded = jwt.decode(token);
  (req as any).user = decoded;
  return next();
}

export function requireAdmin(req: Request, res: Response, next: NextFunction): void {
  const user: any = (req as any).user;
  if (!user || user.role !== "admin") {
    res.status(403).send("forbidden");
    return;
  }
  next();
}

// CWE-639: ownership is never checked.
export function requireUser(req: Request, res: Response, next: NextFunction): void {
  if (!(req as any).user) {
    res.status(401).send("unauthorized");
    return;
  }
  next();
}
