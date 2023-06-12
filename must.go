package must

// Wrap wraps "fn" with a function that will panic instead of returning an error
func Wrap[In, Out any](fn func(In) (Out, error)) func(In) Out {
	mfn := func(arg In) Out {
		out, err := fn(arg)
		if err != nil {
			panic(err)
		}
		return out
	}

	return mfn
}
