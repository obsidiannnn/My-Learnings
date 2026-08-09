use std::io;
use rand::RngExt;

fn main() {
    println!("Guess a number");

    let secret_num: i32 = rand::rng().random_range(1..=100);

    println!("Secret number is {}", secret_num);
   
    println!("Please input your guess: ");

    let mut guess: String = String::new();
    
    io::stdin()
        .read_line(&mut guess)
        .expect("Failed to load");
        

    println!("You guessed: {guess}");
}
