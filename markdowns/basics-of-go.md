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
var i := 1000
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

Functions Declarations
====
