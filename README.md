# ptr

Generic pointer functions.

## Usage

```
go get github.com/0xch4z/ptr
```

## Why is this even a thing?

#### `ptr.To`

There is no function to create a reference to a literal value in Go's built-ins
or it's standard library. Before generics, I found myself copy-pasting the same
`stringPtr(string) *string`, `intPtr(int) *int`, etc. function declarations in
every code base I started. This is obviously not ideal and gets increasingly less
maintainable as the types of literal values you need to make references to vary
and get more complex.

```go
takesOptionalString(&"hi") // illegal
takesOptionalString(aws.String("hi")) // what does AWS have to do with this???
func strPtr(s string) *string { return &s }; takesOptionalString(strPtr("hi")) // unmaintainable

takesOptionalInt(ptr.From(5)) // noice
```

#### `ptr.ValueFrom`

Another frustration is that while Go makes it easy to handle missing input data
with zero-values for all types in the language, there is no mechanism for safely
dereferencing a pointer if it's not nil or interpreting the absence of value as
a zero-value.

```go
type QueryOptions struct { Page *int } // pretend we don't maintain this struct

var opts QueryOptions
var page int

page = *opts.Page // might panic
if opts.Page != nil { page = *opts.Page } // clunky

page = ptr.ValueFrom(opts.Page) // noice
```
