package transpiler

import (
	"io"
	"strings"
)

type indenter struct {
	w io.Writer
	s string
}

func newIndenter(w io.Writer) *indenter {
	i := new(indenter)
	i.w = w
	return i
}

func (i *indenter) Write(p []byte) (n int, err error) {
	i.s += string(p)
	return len(p), nil
}

func (i *indenter) flush() error {
	tokens := strings.Split(i.s, "\n")
	for _, t := range tokens {
		_, err := i.w.Write([]byte(t + "\n"))
		if err != nil {
			return err
		}
	}
	return nil
}
