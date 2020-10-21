/*
Package set provides a simple set implementation. Each element can be added
only once in this set. Optionally, the user can set a type for the set, so only
elements of this type are accepted.

A custom error type TypeError is defined, which can hold information about
the type of the set and a new type that tried to get enforced to it.

A set cannot have other sets (or maps) as elements, as they are not hashable
and the runtime panics.

In order to achieve the type enforcement, the reflect package is used, with
whatever performance penalties this might have.

This set implementation is not thread-safe.
*/

// TODO: Make add take variadic arguments?
package set

import (
	"fmt"
	"reflect"
)

// TypeError indicates an incongruity between the type the set has and another
// type the user tries to set or between two sets. It holds information about
// both the new and the old type, as well as an error message.
type TypeError struct {
	CurrentType reflect.Type // The type the set already has
	NewType     reflect.Type // The type that caused the error
	Err         string
}

func (e *TypeError) Error() string {
	e.Err = fmt.Sprintf("%s is not a valid type for the set with type %s.", e.NewType, e.CurrentType)

	return e.Err
}

// Set is a structure that allows no duplicate entries.
type Set struct {
	Set          map[interface{}]struct{}
	elementsType reflect.Type
}

var exists = struct{}{}

// NewSet allocates memory for a new Set. A Set created this way can have
// elements of varying types.
func NewSet() (s Set) {
	s.Set = make(map[interface{}]struct{})
	s.elementsType = nil

	return s
}

// SetType sets the type of the elements the set accepts. The type of the set
// can only be set once. Hence, if one tries to set it again, the function will
// return false and the type will not change. If the set already has elements of
// other type(s) in it when this function is called, nothing will happen, but
// future elements will have to be of the type specified here.
func (s *Set) SetType(elem interface{}) error {
	newType := reflect.ValueOf(elem).Type()

	if s.elementsType == nil {
		s.elementsType = newType
		return nil
	}

	return &TypeError{s.elementsType, newType, "Trying to re-set the set's type."}
}

// CreateSet creates a set and inserts elem in it. Moreover, it sets the type of
// the set to be that of the element. That means that Sets created this way will
// only accept elements of the same type as the initial element.
func CreateSet(elem interface{}) (s Set) {
	s.Set = make(map[interface{}]struct{})
	s.elementsType = reflect.ValueOf(elem).Type()

	s.Add(elem)

	return s
}

// properType checks if elem is the same type as Set.elementsType.
func (s *Set) properType(elem interface{}) bool {
	if s.elementsType == nil || reflect.ValueOf(elem).Type() == s.elementsType {
		return true
	}

	return false
}

// SameType checks if the set s1 is of the same type as the set s2. If it is,
// it returns true.
func (s1 *Set) SameType(s2 Set) bool {
	if s1.elementsType != s2.elementsType {
		return false
	}

	return true
}

// Add adds elem to the set s. If the element exists in the set or if the
// element is not of the correct type,, no addition is performed and false is
// returned. Otherwise, a new entry is added and it retuns true.
func (s *Set) Add(elem interface{}) bool {
	if !s.properType(elem) {
		return false
	}

	// bool defaults to false, so , if an element does not exist, it will map to false.
	if _, ok := s.Set[elem]; !ok {
		s.Set[elem] = exists

		return true
	}

	return false
}

// Has returns true if the element provided already exists in the set, otherwise false.
func (s *Set) Has(elem interface{}) bool {
	if !s.properType(elem) {
		return false
	}

	_, ok := s.Set[elem]
	return ok
}

// Length returns the number of elements in the set s.
func (s *Set) Length() int {
	return len(s.Set)
}

// Empty returns true if the set is empty, if it has no elements, otherwise
// false.
func (s *Set) Empty() bool {
	if s.Length() > 0 {
		return false
	}

	return true
}

// Equal returns true if the provided sets is equal to the messenger passing
// set. That means that both sets have the very same elements, in count and
// type.
func (s1 *Set) Equal(s2 Set) bool {
	if s1.Empty() && s2.Empty() {
		return true
	}

	if s1.Length() != s2.Length() {
		return false
	}

	if reflect.DeepEqual(s1.Set, s2.Set) {
		return true
	}

	return false
}

// Subset returns true if s1 is a subset of s2. That means that all elements of
// s1 are also elements of s2.
func (s1 *Set) Subset(s2 Set) bool {
	if s1.Equal(s2) {
		return true
	}

	if s1.Empty() {
		return true
	}

	if s1.Length() > s2.Length() {
		return false
	}

	s := make([]bool, len(s1.Set), len(s1.Set))
	i := 0
	// Testing in this order because s2 will always be bigger, otherwise the
	// length clause above will have caught it.
	for e1 := range s1.Set {
		for e2 := range s2.Set {
			if e1 == e2 {
				s[i] = true
			}
		}
		i++
	}

	for c := range s {
		// If some element of s is false, that means an element of s1
		// was not found in s2, so s1 is not a subset of s2. false is an
		// 'untyped boolean', hence the need to use 0.
		if c != 0 {
			return false
		}
	}

	return true
}

// Union returns the union of the two sets.
func (s1 *Set) Union(s2 Set) (Set, error) {
	// An empty set will have nil elementsType which will trigger this
	// clause. But it's a valid operation to make the union of a set with
	// the empty set.
	if !s1.Empty() && !s2.Empty() {
		if !s1.SameType(s2) {
			return Set{}, &TypeError{s1.elementsType, s2.elementsType,
			"The sets' types do not match."}
		}
	}

	s := NewSet()
	s.elementsType = s1.elementsType

	for v := range s1.Set {
		s.Add(v)
	}

	for v := range s2.Set {
		s.Add(v)
	}

	return s, nil
}

// Intersection returns the intersection of the two sets.
func (s1 *Set) Intersection(s2 Set) (Set, error) {
	if !s1.SameType(s2) {
		return Set{}, &TypeError{s1.elementsType, s2.elementsType,
			"The sets' types do not match."}
	}

	s := NewSet()
	s.elementsType = s1.elementsType

	for v1 := range s1.Set {
		for v2 := range s2.Set {
			if v1 == v2 {
				s.Add(v1)
			}
		}
	}

	return s, nil
}

// Difference returns a set that is the difference (also termed as relative
// complement) of s1 from s2. The resulting set is the s1\s2.
func (s1 *Set) Difference(s2 Set) (Set, error) {
	if !s1.SameType(s2) {
		return Set{}, &TypeError{s1.elementsType, s2.elementsType,
			"The sets' type do not match."}
	}

	s := NewSet()

	for v1 := range s1.Set {
		if !s2.Has(v1) {
			s.Add(v1)
		}
	}

	return s, nil
}
