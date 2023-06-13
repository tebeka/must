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

func TestWrapOK(t *testing.T) {
	fn := must.Wrap(NewRing[int])
	size := 7
	buf := fn(size)
	if buf.size != size {
		t.Fatal(buf)
	}
}

func TestWrapPanic(t *testing.T) {
	fn := must.Wrap(NewRing[int])
	size := -1
	defer func() {
		if err := recover(); err == nil {
			t.Fatalf("size %d - no panic", size)
		}
	}()

	fn(size)
}

func Max(values ...int) (int, error) {
	if len(values) == 0 {
		return 0, fmt.Errorf("max with no values")
	}

	m := values[0]
	for _, v := range values[1:] {
		if v > m {
			m = v
		}
	}

	return m, nil
}

func TestWrapVariadicOK(t *testing.T) {
	fn := must.WrapVariadic(Max)
	v := fn(1, 2, 3)
	if v != 3 {
		t.Fatal(v)
	}
}

func TestWrapVariadicPanic(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Fatalf("no panic")
		}
	}()
	fn := must.WrapVariadic(Max)
	fn()
}
