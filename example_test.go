package must_test

import (
	"fmt"

	"github.com/tebeka/must"
)

func ExampleWrap() {
	type RingBuffer struct {
		size int
	}

	New := func(size int) (*RingBuffer, error) {
		if size <= 0 {
			return nil, fmt.Errorf("size (%d) must be > 0", size)
		}
		r := RingBuffer{size: size}
		return &r, nil
	}

	MustNew := must.Wrap(New)
	val := MustNew(7)
	fmt.Println(val)
	// Output:
	// &{7}
}
