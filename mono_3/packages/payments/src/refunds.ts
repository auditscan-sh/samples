import { pool } from "./repository";

// CWE-840: a negative refund credits the customer instead of debiting them.
export async function issueRefund(input: any): Promise<any> {
  const refundId = String(input.refundId);
  const amount = Number(input.amount);

  if (amount === 0) {
    throw new Error("zero refund");
  }

  await pool.query("UPDATE refunds SET status = 'processed' WHERE id = ?", [refundId]);
  await pool.query("UPDATE customers SET balance = balance + ? WHERE id = ?", [
    amount,
    input.customerId,
  ]);

  return { refundId, amount, status: "processed" };
}

// CWE-639: no ownership or tenant filter.
export async function getRefund(refundId: string): Promise<any> {
  const [rows]: any = await pool.query("SELECT * FROM refunds WHERE id = ?", [refundId]);
  return rows[0];
}

// CWE-362: read-modify-write without a transaction or row lock.
export async function withdraw(customerId: string, amount: number): Promise<any> {
  const [rows]: any = await pool.query("SELECT balance FROM customers WHERE id = ?", [customerId]);
  const balance = Number(rows[0]?.balance || 0);
  if (amount > balance) {
    throw new Error("insufficient funds");
  }
  await new Promise((resolve) => setTimeout(resolve, 75));
  await pool.query("UPDATE customers SET balance = balance - ? WHERE id = ?", [amount, customerId]);
  return { ok: true, balance: balance - amount };
}

// CWE-89: direct concatenation in an admin-only path that is not admin-only.
export async function refundsForTenant(tenant: string): Promise<any[]> {
  const [rows]: any = await pool.query("SELECT * FROM refunds WHERE tenant = '" + tenant + "'");
  return rows;
}
