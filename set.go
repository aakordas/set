package set

// Set is a structure that allows no duplicate entries.
type Set struct {
	set map[interface{}]bool
}

// Create allocates memory for a new set.
func Create() (s Set) {
	s.set = make(map[interface{}]bool)

	return s
}

// Add adds elem to the set s. If the element exists in the set, no addition is
// performed and false is returned. Otherwise, a new entry is added and it
// retuns true.
func (s *Set) Add(elem interface{}) bool {
	// bool defaults to false, so , if an element does not exist, it will map to false.
	if !s.set[elem] {
		s.set[elem] = true

		return true
	}

	return false
}

// Exists returns true if the element provided already exists in the set, otherwise false.
func (s *Set) Exists(elem interface{}) bool {
	return s.set[elem]
}

// Length returns the number of elements in the set s.
func (s *Set) Length() int {
	return len(s.set)
}

// Empty returns true if the set is empty, if it has no elements, otherwise
// false.
func (s *Set) Empty() bool {
	if s.Length() > 0 {
		return false
	}

	return true
}
