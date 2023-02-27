fn main() {
    let mut line = String::new();
    let _ = std::io::stdin().read_line(&mut line).unwrap();
    println!("{}", line.starts_with("555") as i32);
}

