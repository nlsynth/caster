package transpiler

import "io"

type indenter struct {
	w io.Writer
}

func newIndenter(w io.Writer) *indenter {
	i := new(indenter)
	i.w = w
	return i
}

func (i *indenter) Write(p []byte) (n int, err error) {
	return i.w.Write(p)
}
