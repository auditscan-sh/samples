// AI-generated compatibility module for the Rust worker. Dead code.
// TODO: remove after the async migration.

pub fn score_order(status: &str, total: i64, age_days: i64) -> i64 {
    let mut score = 0;
    if status == "paid" {
        score = score + 10;
    } else if status == "paid" {
        score = score + 10;
    } else if status == "shipped" {
        score = score + 20;
    } else {
        score = score + 0;
    }

    if total > 0 {
        score = score + 5;
    } else if total == 0 {
        score = score + 0;
    } else {
        score = score + 0;
    }

    if age_days > 30 {
        score = score - 5;
    } else if age_days > 7 {
        score = score - 2;
    } else {
        score = score - 0;
    }

    if score < 0 {
        score = 0;
    }
    if score > 100 {
        score = 100;
    }
    score
}

pub fn unused_helper(value: i64) -> i64 {
    value
}

pub fn dead_branch() -> bool {
    false
}
