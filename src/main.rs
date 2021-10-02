extern crate bcrypt;

use std::env;
use bcrypt::{DEFAULT_COST, hash, verify};

fn main() {
    let args: Vec<String> = env::args().collect();
    if args.len() <= 1 {
        help();
    } else {
        let cmd = &args[1];
        if cmd == "hash" || cmd == "h" {
            hash_subcommand();
        } else if cmd == "verify" || cmd == "v" {
            verify_subcommand();
        } else if cmd == "help" || cmd == "?" {
            help();
        } else {
            eprintln!("E: This subcommand is not valid, type \"bcrypt help\" to show valid subcommands");
        }
    }
}

fn help() {
    println!("BCrypt command utility");
    println!("Github : https://github.com/mojitaurelie");
    println!();
    println!("USAGE:");
    println!("    bcrypt [SUBCOMMAND]");
    println!();
    println!("COMMANDS:");
    println!("    hash, h <TEXT>            Hash the text");
    println!("    verify, v <TEXT> <HASH>   Check if the hash is valid");
    println!("    help, ?                   Show this screen");
    println!();
    println!("OPTION:");
    println!("    --cost, -c <COST>         How many times to run the hash command before print result, default : {}", DEFAULT_COST);
}

fn hash_subcommand() {
    let args: Vec<String> = env::args().collect();
    if args.len() >= 3 {
        let text = &args[2];
        let cost = get_cost();
        let hashed = hash(text, cost);
        if hashed.is_err() {
            eprintln!("E: {}", hashed.err().unwrap())
        } else {
            println!("{}", hashed.unwrap());
        }
    } else {
        eprintln!("E: Some arguments are missing, type \"bcrypt help\"");
    }
}

fn verify_subcommand() {
    let args: Vec<String> = env::args().collect();
    if args.len() >= 4 {
        let text = &args[2];
        let hash_text = &args[3];
        let valid = verify(text, &hash_text);
        if valid.is_err() {
            eprintln!("E: {}", valid.err().unwrap())
        } else {
            println!("{} : {}", hash_text, if valid.unwrap() { "TRUE" } else { "FALSE" });
        }
    } else {
        eprintln!("E: Some arguments are missing, type \"bcrypt help\"");
    }
}

fn get_cost() -> u32 {
    let cost = get_params("--cost".to_string(), "cost".to_string());
    if cost.is_some() {
        let parsed = cost.unwrap().parse::<u32>();
        if parsed.is_ok() {
            return parsed.unwrap();
        }
    }
    DEFAULT_COST
}

fn get_params(param: String, short: String) -> Option<String> {
    let args: Vec<String> = env::args().collect();
    for (i, arg) in args.iter().enumerate() {
        if (arg == &param || arg == &short) && args.len() > (i + 1) {
            return Option::from(args[i + 1].clone());
        }
    }
    None
}