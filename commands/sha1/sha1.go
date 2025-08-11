package sha1

import (
	"context"
	"crypto/sha1"
	"flag"
	"fmt"
	"hash_utils/tools"
	"os"
	"strings"

	"github.com/google/subcommands"
)

type (
	SHA1Cmd struct {
		file string
	}
)

func (*SHA1Cmd) Name() string     { return "sha1" }
func (*SHA1Cmd) Synopsis() string { return "" }
func (*SHA1Cmd) Usage() string {
	return `Usage: hash_utils sha1

Options:
`
}

func (p *SHA1Cmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.file, "file", "", "get the checksum of a file")
}

func (p *SHA1Cmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if len(p.file) > 0 {
		file, err := os.OpenFile(p.file, os.O_RDONLY, 0)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error: failed to open the file:", err)
			return subcommands.ExitFailure
		}
		defer file.Close()

		h, err := tools.Hash(sha1.New(), file)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			return subcommands.ExitFailure
		}
		fmt.Printf("%s: %s\n", f.Name(), h)

		return subcommands.ExitSuccess
	}

	for _, arg := range f.Args() {
		h, err := tools.Hash(sha1.New(), strings.NewReader(arg))
		if err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			return subcommands.ExitFailure
		}
		fmt.Printf("%s: %s\n", arg, h)
	}

	return subcommands.ExitSuccess
}
