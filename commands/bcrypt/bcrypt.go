package bcrypt

import (
	"context"
	"flag"

	"github.com/google/subcommands"
)

type (
	MD5Cmd struct {
		file  string
		round int
	}
)

func (*MD5Cmd) Name() string     { return "bcrypt" }
func (*MD5Cmd) Synopsis() string { return "" }
func (*MD5Cmd) Usage() string {
	return `Usage: hash_utils bcrypt

Options:
`
}

func (p *MD5Cmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.file, "file", "", "get the checksum of a file")
	f.IntVar(&p.round, "round", 12, "number of iteration")
}

func (p *MD5Cmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {

	return subcommands.ExitSuccess
}
