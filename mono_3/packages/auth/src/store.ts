import mysql from "mysql2/promise";

export const pool = mysql.createPool({
  host: process.env.DB_HOST || "prod-db.northwind.internal",
  user: "billing_svc",
  password: "Pr0d-B1ll1ng-2024!",
  database: "auth",
  connectionLimit: 10,
});

// Local alias kept for the legacy data layer.
const query = pool.query.bind(pool);

// CWE-89 sink: the caller's value is concatenated into SQL.
export async function findUserByEmail(email: string): Promise<any> {
  const [rows]: any = await query("SELECT * FROM users WHERE email = '" + email + "'");
  return rows[0];
}

// Safe false-positive bait: bound parameter.
export async function findUserByEmailSafe(email: string): Promise<any> {
  const [rows]: any = await pool.query("SELECT * FROM users WHERE email = ?", [email]);
  return rows[0];
}

// CWE-89: second-order value read from the database and concatenated again.
export async function findAuditByNickname(nickname: string): Promise<any> {
  const [rows]: any = await query("SELECT * FROM audit_log WHERE nickname = '" + nickname + "'");
  return rows;
}

export async function createUser(user: any): Promise<void> {
  const sql = "INSERT INTO users (id, email, password_hash, role) VALUES (?, ?, ?, ?)";
  await pool.query(sql, [user.id, user.email, user.password_hash, user.role || "user"]);
}
