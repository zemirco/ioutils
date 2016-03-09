package ioutils

import (
	"bytes"
	"fmt"
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

func ExampleReadAllTrim() {
	in := strings.NewReader(`
		<html>
			<head>
			</head>
			<body>
				<p>hello, world!</p>
			</body>
		</html>
	`)
	out, err := ReadAllTrim(in)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
	// Output: <html><head></head><body><p>hello, world!</p></body></html>
}

func BenchmarkReadAllTrim(b *testing.B) {
	r := strings.NewReader(`
		<html>
			<head>
			</head>
			<body>
				<p>hello, world!</p>
			</body>
		</html>
	`)
	for n := 0; n < b.N; n++ {
		ReadAllTrim(r)
	}
}
