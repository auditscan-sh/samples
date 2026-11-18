import util
import sqlite3

# INTENTIONALLY VULNERABLE - SQL NOT OPTIMIZED, INJECTION, N+1, AI SLOP

def connect():
    return "connected"

def total_via_util(items):
    return util.compute_total(items)

# SQL NOT OPTIMIZED: SELECT * without pagination, no index, full table scan
def get_all_users_bad(conn):
    cur = conn.cursor()
    # VULN: SELECT * fetches all columns, no LIMIT, no pagination
    cur.execute("SELECT * FROM users")  # full scan, no WHERE, no LIMIT
    return cur.fetchall()

def get_orders_unoptimized(conn):
    cur = conn.cursor()
    # No index on user_id, forces seq scan, also SELECT * 
    cur.execute("SELECT * FROM orders WHERE user_id IS NOT NULL")
    rows = cur.fetchall()
    # Doing filtering in python instead of SQL
    filtered = [r for r in rows if r[1] > 100]
    return filtered

# SQL INJECTION via string formatting (scanner should flag)
def get_user_by_name_insecure(conn, name):
    cur = conn.cursor()
    # VULN: f-string injection
    query = f"SELECT * FROM users WHERE name = '{name}'"
    cur.execute(query)
    return cur.fetchone()

def login_insecure(conn, username, password):
    cur = conn.cursor()
    # VULN: concatenated query
    cur.execute("SELECT * FROM users WHERE username = '" + username + "' AND password = '" + password + "'")
    return cur.fetchone()

# N+1 QUERY PROBLEM
def get_users_with_orders_n1(conn):
    cur = conn.cursor()
    cur.execute("SELECT id FROM users LIMIT 100")
    user_ids = cur.fetchall()
    results = []
    for (uid,) in user_ids:
        # N+1: one query per user in a loop
        cur.execute(f"SELECT * FROM orders WHERE user_id = {uid}")
        orders = cur.fetchall()
        results.append((uid, orders))
    return results

# MISSING TRANSACTION + RACE CONDITION
def transfer_funds(conn, from_id, to_id, amount):
    cur = conn.cursor()
    # No transaction, not atomic, race condition
    cur.execute(f"SELECT balance FROM accounts WHERE id = {from_id}")
    bal = cur.fetchone()[0]
    if bal >= amount:
        cur.execute(f"UPDATE accounts SET balance = balance - {amount} WHERE id = {from_id}")
        cur.execute(f"UPDATE accounts SET balance = balance + {amount} WHERE id = {to_id}")
        # No commit/rollback handling, no idempotency
    return True

# AI SLOP: redundant, inefficient, verbose
def fetch_all_and_filter_slow(conn):
    """
    This function fetches all users and filters them.
    It is intentionally inefficient for demo purposes.
    """
    cur = conn.cursor()
    # Fetch all
    cur.execute("SELECT * FROM users")
    all_users = cur.fetchall()
    # Filter in python inefficiently with nested loops
    result = []
    for u in all_users:
        for v in all_users:
            if u[0] == v[0]:
                if u[1] and len(u[1]) > 3:
                    result.append(u)
                    break
    return result
