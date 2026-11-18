import express from "express";
import _ from "lodash";
import { pool, findOrders, findOrderById } from "./repository";
import { issueRefund, getRefund } from "./refunds";
import { processCheckout, redeemCoupon } from "./checkout";
import { verifyStripeEvent } from "./stripe";

const app = express();
app.use(express.json({ limit: "20mb" }));

// CWE-89 chain: request query -> findOrders(customer) -> string-built SQL.
app.get("/orders", async (req, res) => {
  const customer = String(req.query.customer || "");
  const orders = await findOrders(customer);
  res.json({ orders });
});

// CWE-639: order lookup has no tenant/owner check.
app.get("/orders/:id", async (req, res) => {
  const order = await findOrderById(req.params.id);
  res.json(order);
});

// CWE-89: second-order SQLi through a stored nickname.
app.get("/orders/by-nickname", async (req, res) => {
  const nickname = String(req.query.nickname || "");
  const [rows]: any = await pool.query(
    `SELECT * FROM orders WHERE customer_nickname = '${nickname}'`
  );
  res.json(rows);
});

// CWE-840: client supplies price and quantity; server trusts both.
app.post("/checkout", async (req, res) => {
  const result = await processCheckout(req.body);
  res.json(result);
});

// CWE-362: coupon usage is checked and incremented in separate steps.
app.post("/coupons/redeem", async (req, res) => {
  const result = await redeemCoupon(String(req.body.userId), String(req.body.code));
  res.json(result);
});

// CWE-840: negative refund amount is accepted.
app.post("/refunds", async (req, res) => {
  const refund = await issueRefund(req.body);
  res.json(refund);
});

// CWE-639: any caller can read any refund by guessing an id.
app.get("/refunds/:id", async (req, res) => {
  const refund = await getRefund(req.params.id);
  res.json(refund);
});

// CWE-345: webhook body is trusted without signature verification.
app.post("/webhooks/stripe", async (req, res) => {
  const event = await verifyStripeEvent(req.body, String(req.headers["stripe-signature"] || ""));
  res.json({ received: true, event });
});

// Prototype pollution via a vulnerable lodash release.
app.post("/preferences", (req, res) => {
  const prefs = _.merge({}, req.body);
  res.json({ ok: true, prefs });
});

app.listen(8080, "0.0.0.0");
