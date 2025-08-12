package argon2

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/google/subcommands"
	"golang.org/x/crypto/argon2"
)

type (
	Argon2Cmd struct {
		memory      int
		iterations  int
		parallelism int
		saltLength  int
		keyLength   int
		version     bool
	}
)

func (*Argon2Cmd) Name() string     { return "argon2" }
func (*Argon2Cmd) Synopsis() string { return "" }
func (*Argon2Cmd) Usage() string {
	return `Usage: hash_utils argon2

Options:
`
}

func (p *Argon2Cmd) SetFlags(f *flag.FlagSet) {
	f.IntVar(&p.memory, "memory", 64*1024, "the amount of memory used by the algorithm (in kibibytes)")
	f.IntVar(&p.iterations, "iterations", 3, "the number of iterations over the memory.")
	f.IntVar(&p.parallelism, "parallelism", runtime.NumCPU(), "the number of threads used by the algorithm.")
	f.IntVar(&p.saltLength, "salt-length", 16, "length of the random salt. 16 bytes is recommended for password hashing")
	f.IntVar(&p.keyLength, "key-length", 16, "length of the generated key (or password hash). 16 bytes or more is recommended")
	f.BoolVar(&p.version, "v", false, "show the version of argon2")
}

func (p *Argon2Cmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if p.version {
		fmt.Println(argon2.Version)
		return subcommands.ExitSuccess
	}
	if f.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "error: too many or no argument. this command required 1 argument")
		return subcommands.ExitUsageError
	}
	salt, err := salt(p.saltLength)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		return subcommands.ExitUsageError
	}
	hash := argon2.IDKey([]byte(f.Arg(0)), salt, uint32(p.iterations), uint32(p.memory), uint8(p.parallelism), uint32(p.keyLength))

	fmt.Println(p.toString(salt, hash))
	return subcommands.ExitSuccess
}

func salt(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (p *Argon2Cmd) toString(s, h []byte) string {
	return fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		p.memory,
		p.iterations,
		p.parallelism,
		base64.RawStdEncoding.EncodeToString(s),
		base64.RawStdEncoding.EncodeToString(h))
}
