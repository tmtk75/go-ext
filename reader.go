package osext

import (
	"io"
	"os"

	"code.google.com/p/go.crypto/ssh/terminal"
)

func OptionalReader(f func() io.Reader) io.Reader {
	if terminal.IsTerminal(int(os.Stdin.Fd())) {
		return f()
	}
	return os.Stdin
}
