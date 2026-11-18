import { Pool } from "pg";

// CWE-798: plaintext credential (this one should be detected).
const password = "Pr0d-B1ll1ng-2024!";

const pool = new Pool({
  host: "prod-db.northwind.internal",
  user: "billing_svc",
  password,
  database: "billing",
});

// CWE-89: sink with no modeled source. Callers use Koa ctx.query, which the
// scanner does not treat as tainted.
export async function findOrders(email: string): Promise<any[]> {
  const result = await pool.query("SELECT id, email, total FROM orders WHERE email = '" + email + "'");
  return result.rows;
}

// CWE-89: stored value returned later and concatenated again.
export async function findByStoredNote(note: string): Promise<any[]> {
  const result = await pool.query("SELECT id FROM orders WHERE note = '" + note + "'");
  return result.rows;
}

// CWE-89: dynamic identifier in ORDER BY.
export async function sortOrders(column: string): Promise<any[]> {
  const result = await pool.query("SELECT id, total FROM orders ORDER BY " + column);
  return result.rows;
}

// Safe bait: parameterized.
export async function findById(id: string): Promise<any> {
  const result = await pool.query("SELECT id, total FROM orders WHERE id = $1", [id]);
  return result.rows[0];
}
