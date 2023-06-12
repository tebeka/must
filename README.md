# must - Sprinkle Panic Everywhere

In some cases, you'd like to declare package level variables from a function call.
In this case, you can't check for an error and must panic.

For example, the `regexp` package has [MustCompile](https://pkg.go.dev/regexp#MustCompile) that will either return a `*regexp.Regexp` or panic on invalid regular expression.

The `must` package provide a `Wrap` function that will wrap a function that returns an value and error with a function that will return a value or panic.

For example, assume you have written a [ring buffer](https://en.wikipedia.org/wiki/Circular_buffer):

```go
type RingBuffer struct {
    size int
}

func New(size int) (*RingBuffer, error) {
    if size <= 0 {
        return nil, fmt.Errorf("size (%d) must be > 0", size)
    }
    r := RingBuffer{size: size}
    return &r, nil
}
```

But people would like to declare a package lever buffer to use.
You can write:

```go
var MustNew = must.Wrap(New)
```


### License

MIT