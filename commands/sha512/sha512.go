package sha512

import (
	"context"
	"crypto/sha512"
	"flag"
	"fmt"
	"hash_utils/tools"
	"os"
	"strings"

	"github.com/google/subcommands"
)

type (
	SHA512Cmd struct {
		file string
	}
)

func (*SHA512Cmd) Name() string     { return "sha512" }
func (*SHA512Cmd) Synopsis() string { return "" }
func (*SHA512Cmd) Usage() string {
	return `Usage: hash_utils sha512

Options:
`
}

func (p *SHA512Cmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.file, "file", "", "get the checksum of a file")
}

func (p *SHA512Cmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if len(p.file) > 0 {
		file, err := os.OpenFile(p.file, os.O_RDONLY, 0)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error: failed to open the file:", err)
			return subcommands.ExitFailure
		}
		defer file.Close()

		h, err := tools.Hash(sha512.New(), file)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			return subcommands.ExitFailure
		}
		fmt.Printf("%s: %s\n", f.Name(), h)

		return subcommands.ExitSuccess
	}

	for _, arg := range f.Args() {
		h, err := tools.Hash(sha512.New(), strings.NewReader(arg))
		if err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			return subcommands.ExitFailure
		}
		fmt.Printf("%s: %s\n", arg, h)
	}

	return subcommands.ExitSuccess
}
