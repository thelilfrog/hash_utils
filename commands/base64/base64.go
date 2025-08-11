package base64

import (
	"context"
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/google/subcommands"
)

type (
	Base64Cmd struct {
		file        string
		decoderMode bool
	}
)

func (*Base64Cmd) Name() string     { return "base64" }
func (*Base64Cmd) Synopsis() string { return "" }
func (*Base64Cmd) Usage() string {
	return `Usage: hash_utils base64

Options:
`
}

func (p *Base64Cmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.file, "file", "", "get the checksum of a file")
	f.BoolVar(&p.decoderMode, "decode", false, "decode instead of encode")
}

func (p *Base64Cmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	if p.decoderMode {
		return p.decode(ctx, f, args)
	}
	return p.encode(ctx, f, args)
}

func (p *Base64Cmd) encode(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if len(p.file) > 0 {
		b, err := os.ReadFile(p.file)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error: failed to open the file:", err)
			return subcommands.ExitFailure
		}

		e := base64.NewEncoder(base64.RawStdEncoding, os.Stdout)
		if _, err := e.Write(b); err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			return subcommands.ExitFailure
		}
		fmt.Println()
		return subcommands.ExitSuccess
	}

	for _, arg := range f.Args() {
		fmt.Print(arg, ": ")
		e := base64.NewEncoder(base64.RawStdEncoding, os.Stdout)
		if _, err := e.Write([]byte(arg)); err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			return subcommands.ExitFailure
		}
		fmt.Println()
	}

	return subcommands.ExitSuccess
}

func (p *Base64Cmd) decode(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if len(p.file) > 0 {
		file, err := os.OpenFile(p.file, os.O_RDONLY, 0)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error: failed to open the file:", err)
			return subcommands.ExitFailure
		}
		defer file.Close()

		e := base64.NewDecoder(base64.RawStdEncoding, file)
		if _, err := io.Copy(os.Stdout, e); err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			return subcommands.ExitFailure
		}
		fmt.Println()
		return subcommands.ExitSuccess
	}

	for _, arg := range f.Args() {
		r := strings.NewReader(arg)
		e := base64.NewDecoder(base64.RawStdEncoding, r)
		if _, err := io.Copy(os.Stdout, e); err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			return subcommands.ExitFailure
		}
		fmt.Println()
	}

	return subcommands.ExitSuccess
}
