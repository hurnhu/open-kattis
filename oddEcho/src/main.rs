fn main() {
    let mut line = String::new();
    let _ = std::io::stdin().read_line(&mut line).unwrap();

    let myInt: i32 = line.trim_end().parse().unwrap();
    let mut i = 1;
    let mut x: Vec<String> = Vec::new();
    while i <= myInt {
        line = String::new();
        std::io::stdin().read_line(&mut line).unwrap();
        x.push(line.trim_end().parse().unwrap());
        i = i + 1;
    }
    i = 0;
    while i < myInt {
        if((i+1) % 2 > 0){
            println!(" {}",x[i as usize]);
        }
        i = i + 1;
    }
   
}
