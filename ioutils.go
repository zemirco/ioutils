package ioutils

import (
	"bufio"
	"bytes"
	"io"
)

func ReadAllTrim(in io.Reader) ([]byte, error) {
	var out []byte
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		out = append(out, bytes.TrimSpace(scanner.Bytes())...)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return out, nil
}
