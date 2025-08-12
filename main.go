package main

import (
	"context"
	"flag"
	"hash_utils/commands/argon2"
	"hash_utils/commands/base64"
	"hash_utils/commands/bcrypt"
	"hash_utils/commands/crc16"
	"hash_utils/commands/crc32"
	"hash_utils/commands/crc64"
	"hash_utils/commands/crc8"
	"hash_utils/commands/md5"
	"hash_utils/commands/sha1"
	"hash_utils/commands/sha256"
	"hash_utils/commands/sha512"
	"hash_utils/commands/version"
	"os"

	"github.com/google/subcommands"
)

func main() {
	subcommands.Register(subcommands.HelpCommand(), "help")
	subcommands.Register(subcommands.FlagsCommand(), "help")
	subcommands.Register(subcommands.CommandsCommand(), "help")
	subcommands.Register(&version.VersionCmd{}, "help")

	subcommands.Register(&md5.MD5Cmd{}, "unkeyed cryptographic hash functions")
	subcommands.Register(&sha1.SHA1Cmd{}, "unkeyed cryptographic hash functions")
	subcommands.Register(&sha256.SHA256Cmd{}, "unkeyed cryptographic hash functions")
	subcommands.Register(&sha512.SHA512Cmd{}, "unkeyed cryptographic hash functions")

	subcommands.Register(&crc8.CRC8Cmd{}, "cyclic redundancy checks")
	subcommands.Register(&crc16.CRC16Cmd{}, "cyclic redundancy checks")
	subcommands.Register(&crc32.CRC32Cmd{}, "cyclic redundancy checks")
	subcommands.Register(&crc64.CRC64Cmd{}, "cyclic redundancy checks")

	subcommands.Register(&bcrypt.BCryptCmd{}, "password hashing functions")
	subcommands.Register(&argon2.Argon2Cmd{}, "password hashing functions")

	subcommands.Register(&base64.Base64Cmd{}, "other")

	flag.Parse()
	ctx := context.Background()

	os.Exit(int(subcommands.Execute(ctx)))
}
