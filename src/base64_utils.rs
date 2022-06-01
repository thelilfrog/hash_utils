use clap::{Subcommand};
use std::str;

#[derive(Subcommand)]
pub enum Base64SubCommands {
    /// Encode the text in base64
    Encode {
        /// The text to encode
        text: Option<String>,
    },

    /// Decode the text
    Decode {
        /// Base64 encoded text
        text: Option<String>,
    },
}

pub fn do_base64_encode(text: String) {
    let b64 = base64::encode(text);
    println!("{}", b64);
}

pub fn do_base64_decode(text: String) {
    let res = base64::decode(text);
    if res.is_err() {
        eprintln!("E: {}", res.unwrap_err());
        return;
    }
    let bytes = res.unwrap();
    let utf8_text = str::from_utf8(&bytes);
    if utf8_text.is_err() {
        eprintln!("E: {}", utf8_text.unwrap_err());
        return;
    }
    println!("{}", utf8_text.unwrap());
}