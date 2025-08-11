package tools

import (
	"encoding/hex"
	"fmt"
	"hash"
	"io"

	"github.com/sigurn/crc8"
)

func Hash(h hash.Hash, r io.Reader) (string, error) {
	if _, err := io.Copy(h, r); err != nil {
		return "", fmt.Errorf("failed to read string: %w", err)
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

func CRC8(s []byte, param crc8.Params) string {
	table := crc8.MakeTable(param)
	crc := crc8.Checksum(s, table)
	return fmt.Sprintf("%X", crc)
}
