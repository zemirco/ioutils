// Package ioutils extends io/ioutil.
package ioutils

import (
	"bufio"
	"bytes"
	"io"
)

// ReadAllTrim reads data line by line from r and trims each one.
// It returns a byte slice with all lines concatenated.
func ReadAllTrim(r io.Reader) ([]byte, error) {
	var out []byte
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		out = append(out, bytes.TrimSpace(scanner.Bytes())...)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return out, nil
}
