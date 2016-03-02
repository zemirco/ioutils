package ioutils

import (
	"bytes"
	"strings"
	"testing"
)

func TestReadAllTrim(t *testing.T) {
	in := strings.NewReader(`
		<html>
			<head>
			</head>
			<body>
				<p>hello, world!</p>
			</body>
		</html>
	`)
	expected := []byte("<html><head></head><body><p>hello, world!</p></body></html>")
	out, err := ReadAllTrim(in)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(out, expected) {
		t.Errorf("got %v but expected %v", out, expected)
	}
}
