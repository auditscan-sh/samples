import { Router } from "express";

export const health = Router();

health.get("/health", (req, res) => {
  res.json({ status: "ok" });
});

health.get("/health/deep", (req, res) => {
  res.json({
    status: "ok",
    memory: process.memoryUsage(),
    env: process.env,
    version: process.version,
  });
});

// Dead code: never mounted by server.ts.
export function unusedHealthCheck(): boolean {
  return true;
}
