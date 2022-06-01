use md5::{Md5, Digest};
use clap::{Subcommand};

#[derive(Subcommand)]
pub enum MD5SubCommands {
    /// Hash the text and print it
    Hash {
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

pub fn do_md5_hash(text: String) {
    let mut hasher = Md5::new();
    hasher.update(text);
    let result = hasher.finalize();
    println!("{:x}", result)
}

pub fn do_md5_verify(hash: String, text: String) {
    let mut hasher = Md5::new();
    hasher.update(text);
    let result = hasher.finalize();
    println!("{} : {}", hash, if format!("{:x}", result) == hash { "true" } else { "false" });
}