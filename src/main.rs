mod bcrypt_utils;
mod md5_utils;
mod base64_utils;

pub use crate::bcrypt_utils::{do_bcrypt_hash, do_bcrypt_verify, BCryptSubCommands};
pub use crate::md5_utils::{do_md5_hash, do_md5_verify, MD5SubCommands};
pub use crate::base64_utils::{do_base64_encode, do_base64_decode, Base64SubCommands};
use clap::{Parser, Subcommand};

#[derive(Parser)]
#[clap(author, version, about, long_about = None)]
struct Cli {
    #[clap(subcommand)]
    command: Option<Commands>,
}

#[derive(Subcommand)]
enum Commands {
    /// Use the bcrypt algorithm
    Bcrypt {
        #[clap(subcommand)]
        command: Option<BCryptSubCommands>,
    },
    /// Use the md5 algorithm
    MD5 {
        #[clap(subcommand)]
        command: Option<MD5SubCommands>,
    },
    /// Encode and decode base64 text (I know, it's not a hash but I need it in my tools)
    Base64 {
        #[clap(subcommand)]
        command: Option<Base64SubCommands>,
    },
}

fn main() {
    let cmd = Cli::parse();

    match &cmd.command {
        Some(Commands::Bcrypt { command }) => {
            match &command {
                Some(BCryptSubCommands::Hash { cost, text }) => {
                    if let Some(value) = text {
                        do_bcrypt_hash(String::from(value), *cost);
                    } else {
                        eprintln!("E: Missing the text value")
                    }
                }
                Some(BCryptSubCommands::Verify { hash, text }) => {
                    if let Some(value) = text {
                        if let Some(hashed_value) = hash {
                            do_bcrypt_verify(String::from(hashed_value), String::from(value));
                        } else {
                            eprintln!("E: Missing the hash")
                        }
                    } else {
                        eprintln!("E: Missing the text value")
                    }
                }
                None => {eprintln!("E: Subcommand invalid, use 'help' to show subcommand available")}
            }
        }
        Some(Commands::MD5 { command }) => {
            match &command {
                Some(MD5SubCommands::Hash { text }) => {
                    if let Some(value) = text {
                        do_md5_hash(String::from(value));
                    } else {
                        eprintln!("E: Missing the text value")
                    }
                }
                Some(MD5SubCommands::Verify { hash, text }) => {
                    if let Some(value) = text {
                        if let Some(hashed_value) = hash {
                            do_md5_verify(String::from(hashed_value), String::from(value));
                        } else {
                            eprintln!("E: Missing the hash")
                        }
                    } else {
                        eprintln!("E: Missing the text value")
                    }
                }
                None => {eprintln!("E: Subcommand invalid, use 'help' to show subcommand available")}
            }
        }
        Some(Commands::Base64 { command }) => {
            match &command {
                Some(Base64SubCommands::Encode { text }) => {
                    if let Some(value) = text {
                        do_base64_encode(String::from(value));
                    } else {
                        eprintln!("E: Missing the text value")
                    }
                }
                Some(Base64SubCommands::Decode { text }) => {
                    if let Some(value) = text {
                        do_base64_decode(String::from(value));
                    } else {
                        eprintln!("E: Missing the encoded value")
                    }
                }
                None => {eprintln!("E: Subcommand invalid, use 'help' to show subcommand available")}
            }
        }
        None => {eprintln!("E: Subcommand invalid, use 'help' to show subcommand available")}
    }
}