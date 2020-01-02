/* This implementation, with anonymous struct, is copied from online. I could
link the latest source I found it from, but it's in many places. Same goes for
the boolean version.  */
package set

import (
	"reflect"
        "fmt"
)

type TypeError struct {
        newType reflect.Type // The type that caused the error
        err string
}

func (e *TypeError) Error() string {
        return fmt.Sprintf("%s is not a valid type for the set.", e.newType)
}

// Set is a structure that allows no duplicate entries.
type Set struct {
	Set          map[interface{}]struct{}
	elementsType reflect.Type
}

var exists = struct{}{}

// CreateSet allocates memory for a new Set. A Set created this way can have
// elements of varying types.
func CreateSet() (s Set) {
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

	return &TypeError{newType, "Invalid type"}
}

// Create creates a set and inserts elem in it. Moreover, it sets the type of
// the set to be that of the element. That means that Sets created this way will
// only accept elements of the same type as the initial element.
func Create(elem interface{}) (s Set) {
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

// Exists returns true if the element provided already exists in the set, otherwise false.
func (s *Set) Exists(elem interface{}) bool {
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
