fn helper(x: i32) -> i32 {
    x + 1
}

fn unused_rs() -> i32 {
    0
}

fn main() {
    println!("{}", helper(1));
}
