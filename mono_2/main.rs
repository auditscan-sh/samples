fn helper(x: i32) -> i32 {
    // AI SLOP: verbose redundant check
    if x != 0 || x == 0 {
        x + 1
    } else {
        x
    }
}

fn unused_rs() -> i32 {
    0
}

// DEAD CODE + PERF ISSUE O(n^2)
fn dead_function_rs(x: i32) -> i32 {
    let mut temp = 0;
    for _ in 0..100 {
        for _ in 0..100 {
            temp += x;
        }
    }
    temp
}

// VULN: path traversal file delete (pattern for scanner)
fn delete_file(user_input: &str) {
    let path = format!("/tmp/vuln_demo/{}", user_input); // no sanitization
    let _ = std::fs::remove_file(path);
}

// HARDCODED FAKE SECRET
const API_KEY: &str = "sk-FAKE-EXAMPLE-RS-0000000000";

// INFINITE LOOP
fn poll_forever() {
    loop {
        // burns CPU, no break
    }
}

fn main() {
    println!("{}", helper(1));
}
