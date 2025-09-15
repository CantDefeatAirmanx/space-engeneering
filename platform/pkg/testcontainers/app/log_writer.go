package app

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type PrefixedWriter struct {
	writer io.Writer
	prefix string
}

func NewPrefixedWriter(writer io.Writer, prefix string) *PrefixedWriter {
	return &PrefixedWriter{
		writer: writer,
		prefix: prefix,
	}
}

func (pw *PrefixedWriter) Write(p []byte) (n int, err error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			_, err = fmt.Fprintf(pw.writer, "[%s] %s\n", pw.prefix, line)
			if err != nil {
				return 0, err
			}
		}
	}
	return len(p), nil
}
