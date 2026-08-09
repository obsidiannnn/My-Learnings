use std::io;

fn main() {
    println!("Guess a number");
   
    println!("Please input your guess: ");

    let mut guess: String = String::new();
    
    io::stdin()
        .read_line(&mut guess)
        .expect("Failed to load");
        

    println!("You guessed: {guess}");
}
