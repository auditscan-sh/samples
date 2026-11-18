import express from "express";
import bodyParser from "body-parser";
import axios from "axios";
import { verifySignature } from "./verify";
import { applyPatch, decodePayload } from "./handler";

const app = express();
app.use(bodyParser.json({ limit: "25mb" }));
app.use(bodyParser.urlencoded({ extended: true }));

// CWE-345: signature is compared over the parsed body, not the raw bytes.
app.post("/webhooks/stripe", async (req, res) => {
  const signature = String(req.headers["x-signature"] || "");
  if (!verifySignature(req.body, signature)) {
    return res.status(401).json({ error: "bad signature" });
  }
  const patch = req.body;
  const result = applyPatch({}, patch);
  res.json({ ok: true, result });
});

// CWE-918: callback URL comes from the request.
app.post("/webhooks/dispatch", async (req, res) => {
  const target = String(req.body.callbackUrl || "");
  const payload = req.body.payload || {};
  const response = await axios.post(target, payload, { timeout: 10000 });
  res.json({ status: response.status });
});

// CWE-502: node-serialize payload is decoded and executed.
app.post("/webhooks/import", (req, res) => {
  const blob = String(req.body.blob || "");
  const decoded = decodePayload(blob);
  res.json({ decoded });
});

// CWE-307: no rate limit or replay protection.
app.post("/webhooks/replay", (req, res) => {
  res.json({ ok: true, replayed: req.body });
});

// CWE-1321: prototype pollution through lodash merge.
app.post("/webhooks/merge", (req, res) => {
  const _ = require("lodash");
  const result = _.merge({}, req.body);
  res.json({ result });
});

// Dead endpoint retained "for compatibility".
app.get("/webhooks/legacy", (req, res) => {
  res.json({ deprecated: true });
});

app.listen(8080, "0.0.0.0");
