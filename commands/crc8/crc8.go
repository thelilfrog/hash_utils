package crc8

import (
	"context"
	"flag"
	"fmt"
	"hash_utils/tools"
	"os"

	"github.com/google/subcommands"
	"github.com/sigurn/crc8"
)

type (
	CRC8Cmd struct {
		file  string
		table string
	}
)

var (
	tables = []crc8.Params{
		crc8.CRC8,
		crc8.CRC8_CDMA2000,
		crc8.CRC8_DARC,
		crc8.CRC8_DVB_S2,
		crc8.CRC8_EBU,
		crc8.CRC8_I_CODE,
		crc8.CRC8_ITU,
		crc8.CRC8_MAXIM,
		crc8.CRC8_ROHC,
		crc8.CRC8_WCDMA,
	}
)

func (*CRC8Cmd) Name() string     { return "crc8" }
func (*CRC8Cmd) Synopsis() string { return "" }
func (*CRC8Cmd) Usage() string {
	return `Usage: hash_utils crc8

Options:
`
}

func (p *CRC8Cmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.file, "file", "", "get the checksum of a file")
	f.StringVar(&p.table, "table", crc8.CRC8.Name, "Predefined CRC-8 algorithms")
}

func (p *CRC8Cmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	var err error
	param := crc8.CRC8
	if len(p.table) > 0 {
		param, err = parse(p.table)
		if err != nil {
			fmt.Printf("Available tables: ")
			for _, table := range tables {
				fmt.Printf("%s ", table.Name)
			}
			fmt.Fprintln(os.Stderr, "error:", err)

			return subcommands.ExitFailure
		}
	}

	if len(p.file) > 0 {
		b, err := os.ReadFile(p.file)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error: failed to read file:", err)
			return subcommands.ExitFailure
		}

		h := tools.CRC8(b, param)

		fmt.Printf("%s: %s\n", f.Name(), h)

		return subcommands.ExitSuccess
	}

	for _, arg := range f.Args() {
		h := tools.CRC8([]byte(arg), param)
		fmt.Printf("%s: %s\n", arg, h)
	}

	return subcommands.ExitSuccess
}

func parse(tableName string) (crc8.Params, error) {
	for _, table := range tables {
		if table.Name == tableName {
			return table, nil
		}
	}
	return crc8.CRC8, fmt.Errorf("invalid table name: %s", tableName)
}
