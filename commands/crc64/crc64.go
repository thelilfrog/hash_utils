package crc64

import (
	"context"
	"flag"
	"fmt"
	"hash/crc64"
	"hash_utils/tools"
	"os"

	"github.com/google/subcommands"
)

type (
	CRC64Cmd struct {
		file  string
		table string
	}
)

func (*CRC64Cmd) Name() string     { return "crc64" }
func (*CRC64Cmd) Synopsis() string { return "" }
func (*CRC64Cmd) Usage() string {
	return `Usage: hash_utils crc64

Algorithm available:
  - ECMA
  - ISO

Options:
`
}

func (p *CRC64Cmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.file, "file", "", "get the checksum of a file")
	f.StringVar(&p.table, "table", "ISO", "Predefined CRC-64 algorithms")
}

func (p *CRC64Cmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	var err error
	var param uint64 = crc64.ISO
	if len(p.table) > 0 {
		param, err = parse(p.table)
		if err != nil {
			fmt.Println("Available tables: ISO ECMA")
			return subcommands.ExitFailure
		}
	}

	if len(p.file) > 0 {
		b, err := os.ReadFile(p.file)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error: failed to read file:", err)
			return subcommands.ExitFailure
		}

		h := tools.CRC64(b, param)

		fmt.Printf("%s: %s\n", f.Name(), h)

		return subcommands.ExitSuccess
	}

	for _, arg := range f.Args() {
		h := tools.CRC64([]byte(arg), param)
		fmt.Printf("%s: %s\n", arg, h)
	}

	return subcommands.ExitSuccess
}

func parse(tableName string) (uint64, error) {
	switch tableName {
	case "ISO":
		return crc64.ISO, nil
	case "ECMA":
		return crc64.ECMA, nil
	}
	return 0, fmt.Errorf("invalid table name: %s", tableName)
}
