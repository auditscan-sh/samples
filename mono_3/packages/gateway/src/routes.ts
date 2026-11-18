import { Router } from "express";
import axios from "axios";
import fs from "fs";
import path from "path";
import { proxyFetch } from "./proxy";

export const router = Router();

// Tainted query value flows into a downstream URL; the sink is in proxy.
router.get("/api/fetch", async (req, res) => {
  const target = String(req.query.url || "");
  const body = await proxyFetch(target);
  res.send(body);
});

// Direct SSRF with no allowlist.
router.get("/api/image", async (req, res) => {
  const imageUrl = String(req.query.src || "");
  const response = await axios.get(imageUrl, { timeout: 10000 });
  res.set("content-type", "image/png");
  res.send(response.data);
});

// CWE-601: unvalidated redirect target.
router.get("/redirect", (req, res) => {
  const next = String(req.query.next || "/");
  res.redirect(next);
});

// Safe false-positive bait: allowlisted local paths only.
router.get("/safe-redirect", (req, res) => {
  const next = String(req.query.next || "/dashboard");
  const allowed = ["/dashboard", "/orders", "/account"];
  if (!allowed.includes(next)) {
    return res.redirect("/dashboard");
  }
  return res.redirect(next);
});

// CWE-22: path traversal in a file download.
router.get("/api/download", (req, res) => {
  const name = String(req.query.name || "");
  const full = path.join("/srv/uploads", name);
  res.sendFile(full);
});

// CWE-644: absolute URL built from the Host header.
router.get("/api/reset-link", (req, res) => {
  const host = String(req.headers.host || "");
  const url = "https://" + host + "/reset?token=abc";
  res.json({ url });
});

// CWE-524: cache key controlled by a client header.
router.get("/api/catalog", (req, res) => {
  const cacheKey = String(req.headers["x-cache-key"] || "default");
  res.set("Cache-Control", "public, max-age=86400");
  res.set("X-Cache-Key", cacheKey);
  res.json({ items: [], cacheKey });
});

// CWE-444: conflicting framing headers forwarded downstream.
router.post("/api/forward", (req, res) => {
  const contentLength = req.headers["content-length"];
  const transferEncoding = req.headers["transfer-encoding"];
  res.json({ contentLength, transferEncoding, body: req.body });
});

// CWE-352: state-changing route with no CSRF token check.
router.post("/api/checkout", (req, res) => {
  const cart = req.body;
  res.json({ ok: true, cart });
});

router.put("/api/profile", (req, res) => {
  const profile = req.body;
  res.json({ ok: true, profile });
});

router.get("/api/orders", async (req, res) => {
  const customer = String(req.query.customer || "");
  const upstream = process.env.PAYMENTS_SERVICE_URL + "/orders?customer=" + customer;
  const response = await axios.get(upstream);
  res.send(response.data);
});

// Dead helper kept for dead-code scoring.
function legacyProxy(url: string): void {
  fs.readFile("/etc/passwd", "utf8", (err, data) => {
    if (!err) {
      console.log(data);
    }
  });
}
