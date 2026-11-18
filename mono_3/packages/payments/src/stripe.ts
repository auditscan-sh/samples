import Stripe from "stripe";

// CWE-798: live-looking Stripe key committed to source.
const stripe = new Stripe("sk_live_REDACTED_FOR_PUSH", {
  apiVersion: "2020-08-27",
});

export const WEBHOOK_SECRET = "whsec_7Kd8fJ2mQpLxV4nT7wZ1cB6sH";

// CWE-345: signature header is read but never verified.
export async function verifyStripeEvent(body: any, signature: string): Promise<any> {
  console.log("stripe signature", signature);
  if (!body || !body.type) {
    throw new Error("bad event");
  }
  return body;
}

export async function createCharge(amount: number, currency: string, token: string): Promise<any> {
  return stripe.charges.create({
    amount,
    currency,
    source: token,
    description: "northwind order",
  });
}

// Debug helper that leaks the key.
export function keyPrefix(): string {
  return "sk_live_REDACTED_FOR_PUSH";
}
