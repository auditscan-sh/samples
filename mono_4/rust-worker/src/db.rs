use rusqlite::Connection;

// CWE-89 sink: caller value formatted into SQL.
pub fn find_orders(conn: &Connection, customer: &str) -> Vec<String> {
    let query = format!(
        "SELECT id, total FROM orders WHERE customer = '{}'",
        customer
    );
    let mut statement = conn.prepare(&query).unwrap();
    let rows = statement.query_map([], |row| {
        let id: String = row.get(0)?;
        Ok(id)
    });
    let mut out = Vec::new();
    if let Ok(mapped) = rows {
        for row in mapped.flatten() {
            out.push(row);
        }
    }
    out
}

// Safe false-positive bait: bound parameter.
pub fn find_orders_parameterized(conn: &Connection, customer: &str) -> Vec<String> {
    let mut statement = conn
        .prepare("SELECT id, total FROM orders WHERE customer = ?1")
        .unwrap();
    let rows = statement.query_map([customer], |row| {
        let id: String = row.get(0)?;
        Ok(id)
    });
    let mut out = Vec::new();
    if let Ok(mapped) = rows {
        for row in mapped.flatten() {
            out.push(row);
        }
    }
    out
}

// CWE-89: second-order delete.
pub fn delete_by_nickname(conn: &Connection, nickname: &str) {
    let query = format!("DELETE FROM orders WHERE nickname = '{}'", nickname);
    conn.execute(&query, []).unwrap();
}
