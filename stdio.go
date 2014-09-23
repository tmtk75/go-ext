package osext

import (
	"bytes"
	"io"
	"os"
)

type Pipe struct {
	old  *os.File
	outC chan string
	//r    *os.File
	w      *os.File
	assign func(f *os.File)
}

func PipeStdout() (*Pipe, error) {
	return pipe(func(c *Pipe) {
		c.old = os.Stdout
	}, func(f *os.File) {
		os.Stdout = f
	})
}

func pipe(save func(c *Pipe), assign func(f *os.File)) (*Pipe, error) {
	var c Pipe
	save(&c)
	r, w, err := os.Pipe()
	if err != nil {
		return nil, err
	}
	c.assign = assign
	c.assign(w)
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

func (c *Pipe) Close() string {
	c.w.Close()
	c.assign(c.old)
	out := <-c.outC
	return out
}
