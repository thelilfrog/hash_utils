package md5

import (
	"context"
	"crypto/md5"
	"flag"
	"fmt"
	"hash_utils/tools"
	"os"
	"strings"

	"github.com/google/subcommands"
)

type (
	MD5Cmd struct {
		file string
	}
)

func (*MD5Cmd) Name() string     { return "md5" }
func (*MD5Cmd) Synopsis() string { return "" }
func (*MD5Cmd) Usage() string {
	return `Usage: hash_utils md5

Options:
`
}

func (p *MD5Cmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.file, "file", "", "get the checksum of a file")
}

func (p *MD5Cmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if len(p.file) > 0 {
		file, err := os.OpenFile(p.file, os.O_RDONLY, 0)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error: failed to open the file:", err)
			return subcommands.ExitFailure
		}
		defer file.Close()

		h, err := tools.Hash(md5.New(), file)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			return subcommands.ExitFailure
		}
		fmt.Printf("%s: %s\n", f.Name(), h)

		return subcommands.ExitSuccess
	}

	for _, arg := range f.Args() {
		h, err := tools.Hash(md5.New(), strings.NewReader(arg))
		if err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			return subcommands.ExitFailure
		}
		fmt.Printf("%s: %s\n", arg, h)
	}

	return subcommands.ExitSuccess
}
