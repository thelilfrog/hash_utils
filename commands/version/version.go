package version

import (
	"context"
	"flag"
	"fmt"
	"runtime"

	"github.com/google/subcommands"
)

const (
	version string = "2.0.0"
)

type (
	VersionCmd struct {
	}
)

func (*VersionCmd) Name() string     { return "version" }
func (*VersionCmd) Synopsis() string { return "" }
func (*VersionCmd) Usage() string {
	return `Usage: hash_utils version
`
}

func (p *VersionCmd) SetFlags(f *flag.FlagSet) {
}

func (p *VersionCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	fmt.Printf("hash_utils %s %s/%s\n", version, runtime.GOOS, runtime.GOARCH)
	return subcommands.ExitSuccess
}
