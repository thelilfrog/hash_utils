package crc32

import (
	"context"
	"flag"
	"fmt"
	"hash/crc32"
	"hash_utils/tools"
	"os"

	"github.com/google/subcommands"
)

type (
	CRC32Cmd struct {
		file  string
		table string
	}
)

func (*CRC32Cmd) Name() string     { return "crc32" }
func (*CRC32Cmd) Synopsis() string { return "" }
func (*CRC32Cmd) Usage() string {
	return `Usage: hash_utils crc32

Algorithm available:
  - Castagnoli
  - Koopman
  - IEEE

Options:
`
}

func (p *CRC32Cmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.file, "file", "", "get the checksum of a file")
	f.StringVar(&p.table, "table", "IEEE", "Predefined CRC-32 algorithms")
}

func (p *CRC32Cmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	var err error
	var param uint32 = crc32.IEEE
	if len(p.table) > 0 {
		param, err = parse(p.table)
		if err != nil {
			fmt.Println("Available tables: IEEE Castagnoli Koopman")
			return subcommands.ExitFailure
		}
	}

	if len(p.file) > 0 {
		b, err := os.ReadFile(p.file)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error: failed to read file:", err)
			return subcommands.ExitFailure
		}

		h := tools.CRC32(b, param)

		fmt.Printf("%s: %s\n", f.Name(), h)

		return subcommands.ExitSuccess
	}

	for _, arg := range f.Args() {
		h := tools.CRC32([]byte(arg), param)
		fmt.Printf("%s: %s\n", arg, h)
	}

	return subcommands.ExitSuccess
}

func parse(tableName string) (uint32, error) {
	switch tableName {
	case "Castagnoli":
		return crc32.Castagnoli, nil
	case "Koopman":
		return crc32.Koopman, nil
	case "IEEE":
		return crc32.IEEE, nil
	}
	return 0, fmt.Errorf("invalid table name: %s", tableName)
}
