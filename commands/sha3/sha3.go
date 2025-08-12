package sha3

import (
	"context"
	"crypto/sha3"
	"flag"
	"fmt"
	"hash_utils/tools"
	"os"
	"strings"

	"github.com/google/subcommands"
)

type (
	SHA3Cmd struct {
		file   string
		length int
	}
)

func (*SHA3Cmd) Name() string     { return "sha3" }
func (*SHA3Cmd) Synopsis() string { return "" }
func (*SHA3Cmd) Usage() string {
	return `Usage: hash_utils sha2

Options:
`
}

func (p *SHA3Cmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.file, "file", "", "get the checksum of a file")
	f.IntVar(&p.length, "length", 256, "change the length, supported length is 224, 256, 384 and 512")
}

func (p *SHA3Cmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	h, err := parseBitLength(p.length)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		return subcommands.ExitUsageError
	}
	if len(p.file) > 0 {
		file, err := os.OpenFile(p.file, os.O_RDONLY, 0)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error: failed to open the file:", err)
			return subcommands.ExitFailure
		}
		defer file.Close()

		h, err := tools.Hash(h, file)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			return subcommands.ExitFailure
		}
		fmt.Printf("%s: %s\n", f.Name(), h)

		return subcommands.ExitSuccess
	}

	for _, arg := range f.Args() {
		h, err := tools.Hash(h, strings.NewReader(arg))
		if err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			return subcommands.ExitFailure
		}
		fmt.Printf("%s: %s\n", arg, h)
	}

	return subcommands.ExitSuccess
}

func parseBitLength(length int) (*sha3.SHA3, error) {
	switch length {
	case 224:
		return sha3.New224(), nil
	case 256:
		return sha3.New256(), nil
	case 384:
		return sha3.New384(), nil
	case 512:
		return sha3.New512(), nil
	}
	return nil, fmt.Errorf("invalid length: supported length is 224, 256, 384 and 512")
}
