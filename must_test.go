package must_test

import (
	"fmt"
	"testing"

	"github.com/tebeka/must"
)

type RingBuffer[T any] struct {
	size int
}

func NewRing[T any](size int) (*RingBuffer[T], error) {
	if size <= 0 {
		return nil, fmt.Errorf("%d - invalid size, must be > 0", size)
	}

	r := RingBuffer[T]{
		size: size,
	}
	return &r, nil
}

func TestMustOK(t *testing.T) {
	fn := must.Wrap(NewRing[int])
	size := 7
	buf := fn(size)
	if buf.size != size {
		t.Fatal(buf)
	}
}

func TestMustPanic(t *testing.T) {
	fn := must.Wrap(NewRing[int])
	size := -1
	defer func() {
		if err := recover(); err == nil {
			t.Fatalf("size %d - no panic", size)
		}
	}()

	fn(size)
}
