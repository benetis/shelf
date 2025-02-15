package internal

import (
	"io"
	"sync"
)

type DebugMsg string

type DebugWriter struct {
	ch   chan string
	lock sync.Mutex
}

func NewDebugWriter() *DebugWriter {
	return &DebugWriter{
		ch: make(chan string, 100), // buffered channel to prevent blocking.
	}
}

func (dw *DebugWriter) Write(p []byte) (n int, err error) {
	dw.lock.Lock()
	defer dw.lock.Unlock()
	s := string(p)
	dw.ch <- s
	return len(p), nil
}

func (dw *DebugWriter) Channel() <-chan string {
	return dw.ch
}

var _ io.Writer = (*DebugWriter)(nil)
