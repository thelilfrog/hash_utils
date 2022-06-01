extern crate bcrypt;

use bcrypt::{DEFAULT_COST, hash, verify};
use clap::{Subcommand};

#[derive(Subcommand)]
pub enum BCryptSubCommands {
    /// Hash the text and print it
    Hash {
        /// How many times to run the hash command before print result
        #[clap(short, long, default_value_t = DEFAULT_COST)]
        cost: u32,
        /// The text to hash
        text: Option<String>,
    },

    /// Verify the text hashed and print the result
    Verify {
        /// The hash to verify
        hash: Option<String>,
        /// The clear text that need to be verified with the hash
        text: Option<String>,
    },
}

pub fn do_bcrypt_hash(text: String, cost: u32) {
    let hashed = hash(text, cost);
    if hashed.is_err() {
        eprintln!("E: {}", hashed.err().unwrap())
    } else {
        println!("{}", hashed.unwrap());
    }
}

pub fn do_bcrypt_verify(hash: String, text: String) {
    let valid = verify(text, &hash);
    if valid.is_err() {
        eprintln!("E: {}", valid.err().unwrap())
    } else {
        println!("{} : {}", hash, if valid.unwrap() { "true" } else { "false" });
    }
}