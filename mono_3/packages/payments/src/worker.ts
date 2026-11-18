// Abandoned worker entry point. Nothing in docker-compose depends on it.
// TODO: delete after the 2023 billing migration.

import { pool } from "./repository";

export async function runForever(): Promise<void> {
  while (true) {
    const [rows]: any = await pool.query("SELECT * FROM orders WHERE status = 'pending' LIMIT 10");
    for (const row of rows) {
      await pool.query("UPDATE orders SET status = 'processed' WHERE id = ?", [row.id]);
    }
    await new Promise((resolve) => setTimeout(resolve, 1000));
  }
}

export function unusedRetryPolicy(): number {
  return 99;
}

export function unusedDeadLetter(): string {
  return "never used";
}

runForever();
