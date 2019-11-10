# set
A simple set implementation in Go. It can accept any type
of element to store. It can even accept different types of elements in the same set! (check Examples). It does not allow duplicates. It is
implemented with Go's maps, with whatever performance
implications this comes with.

## Installation
To install, type
```
go get github.com/aakordas/set
```

## Usage
To use it in a program, do
```
import "github.com/aakordas/set"
```

## Examples

```go
import "github.com/aakordas/set"

s := set.Create()

s.Add(1)
s.Add(2)

if s.Add(3) == false {
    fmt.Println("The element already exists in the set)
}

s.Add(2) // Nothing happens, `2` is not added in the set and the call returns false.

s.Add(true) // TOTALLY FINE!!
s.Add("set") // MORE THAN TOTALLY FINE!!

// No sets of sets are allowed.
s1 := set.Create()
s.Add(s1) // PANIC! Set is not hashable.
```

Also check the tests file for potentially more examples.
