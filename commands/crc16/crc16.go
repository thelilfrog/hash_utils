package crc16

import (
	"context"
	"flag"
	"fmt"
	"hash_utils/tools"
	"os"

	"github.com/google/subcommands"
	"github.com/sigurn/crc16"
)

type (
	CRC16Cmd struct {
		file  string
		table string
	}
)

var (
	tables = []crc16.Params{
		crc16.CRC16_ARC,
		crc16.CRC16_AUG_CCITT,
		crc16.CRC16_BUYPASS,
		crc16.CRC16_CCITT_FALSE,
		crc16.CRC16_CDMA2000,
		crc16.CRC16_CRC_A,
		crc16.CRC16_DDS_110,
		crc16.CRC16_DECT_R,
		crc16.CRC16_DECT_X,
		crc16.CRC16_DNP,
		crc16.CRC16_EN_13757,
		crc16.CRC16_GENIBUS,
		crc16.CRC16_KERMIT,
		crc16.CRC16_MAXIM,
		crc16.CRC16_MCRF4XX,
		crc16.CRC16_MODBUS,
		crc16.CRC16_RIELLO,
		crc16.CRC16_T10_DIF,
		crc16.CRC16_TELEDISK,
		crc16.CRC16_TMS37157,
		crc16.CRC16_USB,
		crc16.CRC16_XMODEM,
		crc16.CRC16_X_25,
	}
)

func (*CRC16Cmd) Name() string     { return "crc16" }
func (*CRC16Cmd) Synopsis() string { return "" }
func (*CRC16Cmd) Usage() string {
	res := "Usage: hash_utils crc16\n\n"
	res += "Algorithm available:\n"
	for _, table := range tables {
		res += "  - " + table.Name + "\n"
	}
	res += "\nOptions:\n"
	return res
}

func (p *CRC16Cmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.file, "file", "", "get the checksum of a file")
	f.StringVar(&p.table, "table", crc16.CRC16_ARC.Name, "Predefined CRC-16 algorithms")
}

func (p *CRC16Cmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	var err error
	param := crc16.CRC16_ARC
	if len(p.table) > 0 {
		param, err = parse(p.table)
		if err != nil {
			fmt.Printf("Available tables: ")
			for _, table := range tables {
				fmt.Printf("%s ", table.Name)
			}
			fmt.Println()
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

		h := tools.CRC16(b, param)

		fmt.Printf("%s: %s\n", f.Name(), h)

		return subcommands.ExitSuccess
	}

	for _, arg := range f.Args() {
		h := tools.CRC16([]byte(arg), param)
		fmt.Printf("%s: %s\n", arg, h)
	}

	return subcommands.ExitSuccess
}

func parse(tableName string) (crc16.Params, error) {
	for _, table := range tables {
		if table.Name == tableName {
			return table, nil
		}
	}
	return crc16.CRC16_ARC, fmt.Errorf("invalid table name: %s", tableName)
}
