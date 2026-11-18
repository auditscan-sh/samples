import Koa from "koa";
import Router from "koa-router";
import axios from "axios";
import ejs from "ejs";
import jwt from "jsonwebtoken";
import serialize from "node-serialize";
import fs from "fs";
import path from "path";
import { execFileSync } from "child_process";
import { findOrders } from "./db";

const app = new Koa();
const router = new Router();
const UPLOAD_ROOT = "/srv/uploads";

// Koa sources: ctx.query and ctx.request.body are not Express req.query/body.
router.get("/orders", async (ctx) => {
  const email = ctx.query.email;
  ctx.body = await findOrders(email);
});

router.get("/proxy", async (ctx) => {
  const target = ctx.query.url;
  const response = await axios.get(target);
  ctx.body = response.data;
});

router.post("/exec", (ctx) => {
  const command = ctx.request.body.cmd;
  // CWE-78: execFileSync is not exec/execSync.
  ctx.body = execFileSync("/bin/sh", ["-c", command]).toString();
});

router.post("/render", (ctx) => {
  // CWE-1336: template injection through EJS.
  ctx.body = ejs.render(ctx.request.body.template, { user: "guest" });
});

router.post("/import", (ctx) => {
  // CWE-502: node-serialize executes embedded functions.
  ctx.body = serialize.unserialize(ctx.request.body.blob);
});

router.post("/token", (ctx) => {
  // CWE-347: decode without verification.
  ctx.body = jwt.decode(ctx.request.body.token);
});

router.get("/read", (ctx) => {
  // CWE-22: traversal through createReadStream, no delete call.
  const target = path.join(UPLOAD_ROOT, ctx.query.file);
  ctx.body = fs.createReadStream(target).pipe(ctx.res);
});

router.get("/redirect", (ctx) => {
  // CWE-601: redirect through Koa context.
  ctx.redirect(ctx.query.next);
});

router.post("/merge", (ctx) => {
  // CWE-1321: prototype pollution via a shallow assign of request data.
  const target: any = {};
  Object.assign(target, ctx.request.body);
  ctx.body = target;
});

router.get("/debug", (ctx) => {
  // CWE-200: full environment exposure.
  ctx.body = process.env;
});

app.use(router.routes());
app.listen(3000, "0.0.0.0");
