//go:build ignore

package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var sentinel = `
		if err != nil {
			panic(err)
		}
		return out
	}

	return wrapper
}
`

var header = `
// Generated by gen.go DO NOT EDIT

package must
`

func genArgs(n int, tmpl string) string {
	vals := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		vals = append(vals, fmt.Sprintf(tmpl, i))
	}
	return strings.Join(vals, ", ")
}

/*
Example function

func Wrap2[In1, In2, Out any](fn func(In1, In2) (Out, error)) func(In1, In2) Out{
	wrapper := func(arg1 In1, arg2 In2) Out {
		out, err := fn(arg1, arg2)
		if err != nil {
			panic(err)
		}
		return out
	}

	return wrapper
}
*/

func gen(n int, w io.Writer) {
	ins := genArgs(n, "In%[1]d")
	argsTyped := genArgs(n, "arg%[1]d In%[1]d")
	args := genArgs(n, "arg%[1]d")

	fmt.Fprintf(w, "// Wrap%d is a version of Wrap for %d arguments.\n", n, n)
	fmt.Fprintf(w,
		"func Wrap%d[%s, Out any](fn func(%s) (Out, error)) func(%s) Out{\n",
		n, ins, ins, ins,
	)
	fmt.Fprintf(w, "\twrapper := func(%s) Out {\n", argsTyped)
	fmt.Fprintf(w, "\t\tout, err := fn(%s)", args)
	fmt.Fprintln(w, sentinel)
}

func main() {
	file, err := os.Create("mustn.go")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer file.Close()
	fmt.Fprintln(file, header)

	// If you have a procedure with 10 parameters, you probably missed some. - Alan Perlis
	for i := 2; i <= 5; i++ {
		gen(i, file)
	}
}
