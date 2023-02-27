fn main() {
    let mut line = String::new();
    let _ = std::io::stdin().read_line(&mut line).unwrap();
    println!("{} {} {}", line.trim_end(), line.trim_end(), line.trim_end());
}
