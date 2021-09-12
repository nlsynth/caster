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
	lines := strings.Split(i.s, "\n")
	level := 0
	for _, line := range lines {
		line = i.trimHead(line)
		level = i.preUpdateLevel(level, line)
		line = i.addIndent(level, line)
		level = i.postUpdateLevel(level, line)
		_, err := i.w.Write([]byte(line + "\n"))
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *indenter) trimHead(s string) string {
	for strings.HasPrefix(s, " ") {
		s = s[1:]
	}
	return s
}

func (i *indenter) addIndent(d int, s string) string {
	for c := 0; c < d; c++ {
		s = " " + s
	}
	return s
}

func (i *indenter) preUpdateLevel(d int, line string) int {
	if strings.HasPrefix(line, "module") ||
		strings.HasPrefix(line, "endmodule") {
		return 0
	}
	if strings.HasPrefix(line, "end") {
		return d - 1
	}
	return d
}

func (i *indenter) postUpdateLevel(d int, line string) int {
	if strings.HasPrefix(line, "module") {
		return 1
	}
	if strings.HasPrefix(line, "case") {
		return d + 1
	}
	if strings.HasSuffix(line, "begin") {
		return d + 1
	}
	return d
}
