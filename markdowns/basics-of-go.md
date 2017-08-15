Basics of Go
===

Variables and Declarations
===

You can declare a variable of a specific type like this:

```go
var i int
```

by default that assigns a default value of `0` for this type. For other types for example booleans it is
`false`, for strings is `""` and so on.

Go has a shorthand variable declaration operator, :=, which can infer the type:

```go
var i := 1000 // notice => No semicolons!
```

note that you cannot `redeclare` the i using `:=` as it has already been declared:

```go
var i := 1000

// Compiler error
i := 1001
```

you can also use that for multiple declarations and assignments:

```go
name, money := "Theo", 1000
```

?[1. Which of the following are valid types in Go?]
-[x] int
-[x] int32
-[ ] complex32
-[x] float32

?[2 . What is the value of c after this statement?: var c complex128]
-[ ] nil
-[ ] (0,0i)
-[ ] (0.0+0.0i)
-[x] (0+0i)

Functions Declarations
====
Bellow are some examples of function declarations in Go:

```go
func log(message string) { // takes 1 parameter
}

func add(a int, b int) int { // takes 2 int parameters and returns an int
}

func power(name string) (int, bool) { // takes 1 string parameter and returns an int and a boolean
}
```


To use just simply call the functions. If you do not care about the return values use `_`

```go
_, isPowerful = power("Theo")
```

?[1. How can I call a function `cdd` that takes 2 doubles and return 3 ints and discard the second?]
-[x] a, _, c := cdd(d, e)
-[ ] a, _, _ = cdd(d, e, f)
-[ ] a, _, c = cdd(d, e, f)
-[ ] a, _, c := cdd(d, e, f)()
