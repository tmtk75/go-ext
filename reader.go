package osext

import (
	"io"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

func OptionalReader(f func() io.Reader) io.Reader {
	if terminal.IsTerminal(int(os.Stdin.Fd())) {
		return f()
	}
	return os.Stdin
}
