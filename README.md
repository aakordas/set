# set
A simple set implementation in Go. It aims to be mathematically correct,
following all the relevant properties of set operations. It can accept any type
of element to store. It can even accept different types of elements in the same
set (check Examples). It does not allow duplicates. It is implemented with Go's
maps and reflect package, with whatever performance implications this comes
with.

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

s := set.NewSet()

s.Add(1)
s.Add(2)

if s.Add(2) == false {
	// The element is not added in the set.
	fmt.Println("The element already exists in the set)
}

s.Add(true) // TOTALLY FINE!!
s.Add("set") // MORE THAN TOTALLY FINE!!

// This creates a set with a starting element '1'.
// Moreover, this sets the type of the set to that of the argument passed, 'int'
// in this case.
intSet := set.CreateSet(1)

intSet.Add(2)
intSet.Add(2) // Nothing happens, since the element already exists in the set.
intSet.Add("3") // This does not get added, since the set has type 'int'.

intSet.Add(3)
s.Add(4)
union := s.Union(intSet) // This will result in the set {1, 2, 3, 4}
intersection := s.Intersection(intSet) // This will result in the set {1, 2}

// No sets of sets are allowed.
s1 := set.CreateSet()
s.Add(s1) // PANIC! Set is not hashable.
```

Also check the tests file for potentially more examples.
