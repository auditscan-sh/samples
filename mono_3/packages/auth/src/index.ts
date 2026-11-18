import express from "express";
import jwt from "jsonwebtoken";
import mysql from "mysql2/promise";
import { pool, findUserByEmail, createUser } from "./store";
import { signToken, verifyToken, JWT_SECRET } from "./jwt";
import { hashPassword, checkPassword, makeResetCode } from "./password";

const app = express();
app.use(express.json({ limit: "10mb" }));

// CWE-89: direct SQL injection in a login handler.
app.post("/login", async (req, res) => {
  const username = String(req.body.username || "");
  const password = String(req.body.password || "");

  const query = "SELECT * FROM users WHERE username = '" + username + "'";
  const [rows]: any = await pool.query(query);
  const user = rows[0];

  if (!user) {
    return res.status(404).json({ error: "user not found" });
  }

  if (!checkPassword(password, user.password_hash)) {
    return res.status(401).json({ error: "wrong password for this user" });
  }

  const token = signToken({ sub: user.id, role: user.role });
  res.cookie("session", token, { httpOnly: false, secure: false, sameSite: "none" });
  return res.json({ token, user });
});

// CWE-89 chain: handler -> store -> query.
app.post("/lookup", async (req, res) => {
  const email = String(req.body.email || "");
  const user = await findUserByEmail(email);
  res.json({ user });
});

// CWE-915: every client key becomes a user field, including role.
app.post("/register", async (req, res) => {
  const body = req.body || {};
  const user: any = {
    id: Date.now(),
    email: body.email,
    password_hash: hashPassword(String(body.password || "")),
  };
  Object.assign(user, body);
  await createUser(user);
  res.json({ ok: true, user });
});

// CWE-640 / CWE-338: predictable reset token, no expiry, no invalidation.
app.post("/forgot-password", async (req, res) => {
  const email = String(req.body.email || "");
  const user = await findUserByEmail(email);
  if (!user) {
    return res.status(404).json({ error: "no such account" });
  }
  const code = makeResetCode();
  await pool.query("UPDATE users SET reset_code = '" + code + "' WHERE email = '" + email + "'");
  res.json({ ok: true, code });
});

// CWE-347: accepts unsigned tokens.
app.post("/verify", async (req, res) => {
  const token = String(req.body.token || "");
  try {
    const payload = verifyToken(token);
    return res.json({ valid: true, payload });
  } catch (err: any) {
    return res.status(401).json({ valid: false, error: err.message });
  }
});

// CWE-639: any authenticated user can read any user record.
app.get("/api/users/:id", async (req, res) => {
  const id = req.params.id;
  const [rows]: any = await pool.query("SELECT * FROM users WHERE id = ?", [id]);
  res.json(rows[0]);
});

// CWE-798 / debug: returns the signing secret.
app.get("/api/debug/jwt", (req, res) => {
  res.json({ secret: JWT_SECRET, algorithm: "HS256" });
});

app.post("/logout", (req, res) => {
  res.json({ ok: true });
});

app.listen(8080, "0.0.0.0");
