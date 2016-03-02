
# ioutils

[![Build Status](https://travis-ci.org/zemirco/ioutils.svg)](https://travis-ci.org/zemirco/ioutils)
[![GoDoc](https://godoc.org/github.com/zemirco/ioutils?status.svg)](https://godoc.org/github.com/zemirco/ioutils)

Package ioutils provides some helpers methods in addition to [io/ioutil](https://golang.org/pkg/io/ioutil/).

## ReadAllTrim

Useful when looking for some string inside http response body, i.e. during tests. You don't have to worry about whitespace between tags.

```go
body, err := ioutils.ReadAllTrim(res.Body)
if err != nil {
	panic(err)
}
// do something with body
```

## Test

```bash
$ go test
```

## License

MIT
