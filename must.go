package must

//go:generate go run gen.go
//go:generate go fmt mustn.go

// Wrap wraps "fn" with a function that will panic instead of returning an error
func Wrap[In, Out any](fn func(In) (Out, error)) func(In) Out {
	wrapper := func(arg In) Out {
		out, err := fn(arg)
		if err != nil {
			panic(err)
		}
		return out
	}

	return wrapper
}

// WrapVariadic is a version of Wrap for variadic functions.
func WrapVariadic[In, Out any](fn func(...In) (Out, error)) func(...In) Out {
	wrapper := func(args ...In) Out {
		out, err := fn(args...)
		if err != nil {
			panic(err)
		}
		return out
	}

	return wrapper
}
