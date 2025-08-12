package tools

import (
	"encoding/hex"
	"fmt"
	"hash"
	"hash/crc32"
	"hash/crc64"
	"io"

	"github.com/sigurn/crc16"
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

func CRC16(s []byte, param crc16.Params) string {
	table := crc16.MakeTable(param)
	crc := crc16.Checksum(s, table)
	return fmt.Sprintf("%X", crc)
}

func CRC32(s []byte, param uint32) string {
	table := crc32.MakeTable(param)
	crc := crc32.Checksum(s, table)
	return fmt.Sprintf("%X", crc)
}

func CRC64(s []byte, param uint64) string {
	table := crc64.MakeTable(param)
	crc := crc64.Checksum(s, table)
	return fmt.Sprintf("%X", crc)
}
