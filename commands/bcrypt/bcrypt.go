package bcrypt

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/google/subcommands"
	"golang.org/x/crypto/bcrypt"
)

type (
	BCryptCmd struct {
		cost  int
		check bool
	}
)

func (*BCryptCmd) Name() string     { return "bcrypt" }
func (*BCryptCmd) Synopsis() string { return "" }
func (*BCryptCmd) Usage() string {
	return `Usage: hash_utils bcrypt

Options:
`
}

func (p *BCryptCmd) SetFlags(f *flag.FlagSet) {
	f.IntVar(&p.cost, "cost", 12, "number of iteration")
	f.BoolVar(&p.check, "check", false, "check a password")
}

func (p *BCryptCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if p.check {
		if f.NArg() != 2 {
			fmt.Fprintln(os.Stderr, "error: invalid number of parameters")
			return subcommands.ExitUsageError
		}
		if err := bcrypt.CompareHashAndPassword([]byte(f.Arg(0)), []byte(f.Arg(1))); err != nil {
			fmt.Println("no match")
			return subcommands.ExitSuccess
		}
		fmt.Println("match")
		return subcommands.ExitSuccess
	}
	if p.cost < bcrypt.MinCost {
		fmt.Fprintln(os.Stderr, "error: invalid cost settings: cost should be over", bcrypt.MinCost)
		return subcommands.ExitUsageError
	}
	if p.cost > bcrypt.MaxCost {
		fmt.Fprintln(os.Stderr, "error: invalid cost settings: cost should be under", bcrypt.MaxCost)
		return subcommands.ExitUsageError
	}
	h, err := bcrypt.GenerateFromPassword([]byte(f.Arg(0)), p.cost)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		return subcommands.ExitFailure
	}
	fmt.Println(string(h))
	return subcommands.ExitSuccess
}
