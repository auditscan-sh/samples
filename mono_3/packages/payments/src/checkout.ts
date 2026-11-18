import { pool } from "./repository";

// CWE-840: price, quantity and discount all come from the client.
export async function processCheckout(payload: any): Promise<any> {
  const price = Number(payload.price);
  const quantity = Number(payload.quantity);
  const discount = Number(payload.discount || 0);
  const total = price * quantity - discount;

  const [result]: any = await pool.query(
    "INSERT INTO orders (customer, total, status) VALUES (?, ?, 'pending')",
    [payload.customer, total]
  );

  return { orderId: result.insertId, total };
}

// CWE-362: two requests can redeem the same single-use coupon.
export async function redeemCoupon(userId: string, code: string): Promise<any> {
  const [rows]: any = await pool.query("SELECT * FROM coupons WHERE code = ?", [code]);
  const coupon = rows[0];
  if (!coupon) {
    throw new Error("invalid coupon");
  }
  if (coupon.uses >= coupon.max_uses) {
    throw new Error("coupon exhausted");
  }

  // Deliberate delay widens the race window.
  await new Promise((resolve) => setTimeout(resolve, 120));

  await pool.query("UPDATE coupons SET uses = uses + 1 WHERE code = ?", [code]);
  await pool.query("INSERT INTO coupon_redemptions (user_id, code) VALUES (?, ?)", [userId, code]);

  return { ok: true, code, remaining: coupon.max_uses - coupon.uses - 1 };
}

// CWE-840: refund total can exceed the original order total.
export function computeRefund(orderTotal: number, requested: number, shipping: number): number {
  return requested + shipping;
}
