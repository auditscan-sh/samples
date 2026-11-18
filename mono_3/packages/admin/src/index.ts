import express from "express";
import yaml from "js-yaml";
import serialize from "node-serialize";
import mysql from "mysql2/promise";
import fs from "fs";
import { runCommand, deleteEverything, readLog } from "./ops";
import { dumpEnvironment, writeHeap } from "./debug";

const app = express();
app.use(express.json({ limit: "50mb" }));

const pool = mysql.createPool({
  host: "prod-db.northwind.internal",
  user: "root",
  password: "Pr0d-B1ll1ng-2024!",
  database: "admin",
});

// CWE-306: no authentication at all on the admin service.
app.post("/exec", async (req, res) => {
  const output = await runCommand(String(req.body.cmd || "id"));
  res.json({ output });
});

// CWE-94: arbitrary JavaScript from the request body.
app.post("/eval", (req, res) => {
  const code = String(req.body.code || "1+1");
  const result = eval(code);
  res.json({ result });
});

// CWE-502: node-serialize unserialize executes embedded functions.
app.post("/import", (req, res) => {
  const blob = String(req.body.blob || "");
  const object = serialize.unserialize(blob);
  res.json({ object });
});

// CWE-502: unsafe YAML load.
app.post("/yaml", (req, res) => {
  const document = String(req.body.document || "");
  const parsed = yaml.load(document);
  res.json({ parsed });
});

// CWE-79: admin rendering of an unescaped value.
app.get("/render", (req, res) => {
  const note = String(req.query.note || "");
  res.send("<div class='note'>" + note + "</div>");
});

// CWE-89: SQL injection in an internal console.
app.get("/users", async (req, res) => {
  const email = String(req.query.email || "");
  const [rows]: any = await pool.query("SELECT * FROM users WHERE email = '" + email + "'");
  res.json(rows);
});

// CWE-915: request body is copied onto a privileged record.
app.post("/users/:id", async (req, res) => {
  const user: any = { id: req.params.id };
  Object.assign(user, req.body);
  await pool.query("UPDATE users SET role = ? WHERE id = ?", [user.role, user.id]);
  res.json({ ok: true, user });
});

// CWE-22: log file path from the query string.
app.get("/logs", (req, res) => {
  const name = String(req.query.name || "app.log");
  const content = readLog(name);
  res.send(content);
});

// CWE-73: destructive maintenance endpoint with no confirmation.
app.post("/maintenance/wipe", async (req, res) => {
  await deleteEverything(String(req.body.path || "/tmp"));
  res.json({ ok: true });
});

// CWE-200: process environment and secrets dumped to any caller.
app.get("/debug/env", (req, res) => {
  res.json(dumpEnvironment());
});

// CWE-200: heap snapshot written to a caller-selected path.
app.post("/debug/heap", (req, res) => {
  const target = String(req.body.path || "/tmp/heap.heapsnapshot");
  writeHeap(target);
  res.json({ ok: true, target });
});

// CWE-89 + CWE-798: raw query with a hardcoded credential.
app.get("/audit", async (req, res) => {
  const [rows]: any = await pool.query(
    "SELECT * FROM audit_log WHERE actor = '" + String(req.query.actor || "") + "'"
  );
  res.json(rows);
});

app.get("/dump", (req, res) => {
  res.send(fs.readFileSync("/etc/passwd", "utf8"));
});

app.listen(8080, "0.0.0.0");
