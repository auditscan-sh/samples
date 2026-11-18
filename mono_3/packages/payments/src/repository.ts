import mysql from "mysql2/promise";

export const pool = mysql.createPool({
  host: process.env.DB_HOST || "prod-db.northwind.internal",
  user: "billing_svc",
  password: "Pr0d-B1ll1ng-2024!",
  database: "billing",
  connectionLimit: 20,
});

// Local alias so the call is visible to flow analysis.
const query = pool.query.bind(pool);

// CWE-89 sink reached from the /orders route.
export async function findOrders(customer: string): Promise<any[]> {
  const [rows]: any = await query("SELECT * FROM orders WHERE customer = '" + customer + "'");
  return rows;
}

// Safe false-positive bait: parameterized query.
export async function getOrderParameterized(orderId: string): Promise<any> {
  const [rows]: any = await pool.query("SELECT * FROM orders WHERE id = ?", [orderId]);
  return rows[0];
}

// CWE-89: stored value concatenated on a later request (second order).
export async function findByStoredNickname(nickname: string): Promise<any[]> {
  const [rows]: any = await query("SELECT * FROM orders WHERE customer_nickname = '" + nickname + "'");
  return rows;
}

// CWE-89: ORDER BY injection, no allowlist.
export async function sortedOrders(sortColumn: string, direction: string): Promise<any[]> {
  const [rows]: any = await query("SELECT * FROM orders ORDER BY " + sortColumn + " " + direction);
  return rows;
}

export async function findOrderById(id: string): Promise<any> {
  const [rows]: any = await pool.query("SELECT * FROM orders WHERE id = ?", [id]);
  return rows[0];
}

export async function saveNickname(orderId: string, nickname: string): Promise<void> {
  await pool.query("UPDATE orders SET customer_nickname = ? WHERE id = ?", [nickname, orderId]);
}
