package osext

import (
	"bytes"
	"io"
	"os"
)

type StdoutContext struct {
	old  *os.File
	outC chan string
	//r    *os.File
	w *os.File
}

func CaptureStdout() (*StdoutContext, error) {
	var c StdoutContext
	c.old = os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		return nil, err
	}
	os.Stdout = w
	//c.r = r
	c.w = w

	c.outC = make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		c.outC <- buf.String()
	}()
	return &c, nil
}

func (c *StdoutContext) End() string {
	c.w.Close()
	os.Stdout = c.old
	out := <-c.outC
	return out
}
