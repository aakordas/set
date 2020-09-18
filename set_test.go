// TODO: Check if the operations are commutative.
package set

import (
	"testing"
)

func TestCreateAndAdd(t *testing.T) {
	s := NewSet()

	if !s.Add(1) {
		t.Errorf("1 already exists in the set %v", s)
	}
}

func TestCreateWithElement(t *testing.T) {
	s := CreateSet(1)

	if !s.Has(1) {
		t.Errorf("1 is not in the set %v", s)
	}
}

func TestAddMultipleElements(t *testing.T) {
	s := NewSet()

	if !s.Add(1) {
		t.Errorf("1 already exists in the set %v", s)
	}

	if !s.Add(2) {
		t.Errorf("2 already exists in the set %v", s)
	}
}

func TestAddDifferentKindElements(t *testing.T) {
	s := NewSet()

	if !s.Add(1) {
		t.Errorf("1 already exists in the set %v", s)
	}

	if !s.Add("2") {
		t.Errorf("\"2\" already exists in the set %v", s)
	}
}

func TestAddDifferentKindElementsThanInitial(t *testing.T) {
	s := CreateSet(1)

	if !s.Add(2) {
		t.Errorf("2 already exists in the set %v", s)
	}

	if s.Add("2") {
		t.Errorf("\"2\" was succesfully added in the set %v", s)
	}
}

func TestAddDuplicate(t *testing.T) {
	s := NewSet()

	if !s.Add(1) {
		t.Errorf("1 already exists in the set %v", s)
	}

	if s.Add(1) {
		t.Errorf("1 does not exist in the set %v", s)
	}
}

func TestAddDifferentKindThanInitialCustomType(t *testing.T) {
	points := []struct {
		X, Y int
	}{
		{1, 2},
		{2, 3},
		{1, 2},
	}

	s := CreateSet(points[0])

	for point := range points {
		s.Add(point)
	}

	if s.Empty() {
		t.Errorf("The set is empty.")
	}

	if s.Add(1) {
		t.Errorf("1 was added succesfully in the set %v", s)
	}
}

func TestHas(t *testing.T) {
	s := CreateSet(1)

	if !s.Has(1) {
		t.Errorf("1 is not in the set %v", s)
	}

	if s.Has(2) {
		t.Errorf("2 is in the set %v", s)
	}

	if s.Has("3") {
		t.Errorf("\"3\" is in the set %v", s)
	}
}

func TestSetType(t *testing.T) {
	s := NewSet()

	s.Add(1)
	s.Add("2")

	var err error
	if err = s.SetType(1); err != nil {
		t.Errorf("Could not change the type of Set %v", s)
	}

	if s.Add("1") {
		t.Errorf("\"1\" was added in the set %v", s)
	}

	if !s.Add(2) {
		t.Errorf("2 was not added in the set %v", s)
	}

	if err = s.SetType("2"); err == nil {
		t.Errorf("Changed the type of the set %v succesfully.", s)
	}
}

func TestLength(t *testing.T) {
	s := NewSet()

	if s.Length() != 0 {
		t.Errorf("The set %v has elements in it?", s)
	}

	s.Add(1)
	s.Add(2)

	if s.Length() != 2 {
		t.Errorf("The set %v has more than two elements.", s)
	}
}

func TestEmpty(t *testing.T) {
	s := NewSet()

	if !s.Empty() {
		t.Errorf("The set %v is not empty", s)
	}

	s.Add(1)

	if s.Empty() {
		t.Errorf("The set %v is empty", s)
	}
}

func TestEqual(t *testing.T) {
	s1 := CreateSet(1)
	s2 := CreateSet(1)

	if !s1.Equal(s2) {
		t.Errorf("The set %v is not equal to the set %v.", s1, s2)
	}

	s2.Add(2)

	if s1.Equal(s2) {
		t.Errorf("The set %v is equal to the set %v.", s1, s2)
	}

	s1.Add(2)
	s1.Add(3)

	if s1.Equal(s2) {
		t.Errorf("The set %v is equal to the set %v.", s1, s2)
	}
}

func TestSubset(t *testing.T) {
	s1 := NewSet()
	s2 := NewSet()

	if !s1.Subset(s2) {
		t.Errorf("The set %v is not a subset of the set %v.", s1, s2)
	}

	s1.Add(1)

	if s1.Subset(s2) {
		t.Errorf("The set %v is a subset of the set %v.", s1, s2)
	}

	s2.Add(1)

	if !s1.Subset(s2) {
		t.Errorf("The set %v is not a subset of the set %v.", s1, s2)
	}

	s2.Add(2)

	if !s1.Subset(s2) {
		t.Errorf("The set %v is not a subset of the set %v.", s1, s2)
	}

	s1.Add(3)

	if s1.Subset(s2) {
		t.Errorf("The set %v is a subset of the set %v.", s1, s2)
	}
}

func TestSameType(t *testing.T) {
	s1 := NewSet()
	s2 := NewSet()

	if !s1.SameType(s2) || !s2.SameType(s1) {
		t.Errorf("The set %v does not have the same type as the set %v.", s1, s2)
	}

	s1.Add(1)
	s2.Add(2)

	if !s1.SameType(s2) || !s2.SameType(s1) {
		t.Errorf("The set %v does not have the same type as the set %v.", s1, s2)
	}

	s3 := CreateSet("a")

	if s1.SameType(s3) || s3.SameType(s1) {
		t.Errorf("The set %v has the same type as the set %v.", s1, s2)
	}
}

func TestUnion(t *testing.T) {
	s1 := CreateSet(1)
	s2 := CreateSet(2)

	got, err := s1.Union(s2)

	want := CreateSet(1)
	want.Add(2)

	if !got.Equal(want) || err != nil {
		t.Errorf("The union of %v and %v resulted in %v, instead of %v.", s1, s2, got, want)
	}

	s3 := CreateSet("a")

	got, err = s1.Union(s3)
	want = Set{}

	if !got.Equal(want) || err == nil {
		t.Errorf("The union of %v and %v succeeded with %v, instead of %v.", s1, s3, got, want)
	}
}

func TestIntersection(t *testing.T) {
	s1 := CreateSet(1)
	s2 := CreateSet(1)

	got, err := s1.Intersection(s2)
	want := CreateSet(1)

	if !got.Equal(want) || err != nil {
		t.Errorf("The intersection of %v and %v resulted in %v, instead of %v.", s1, s2, got, want)
	}

	s3 := CreateSet(2)

	got, err = s1.Intersection(s3)
	want = NewSet()

	if !got.Equal(want) || err != nil {
		t.Errorf("The intersection of %v and %v resulted in %v, instead of %v.", s1, s3, got, want)
	}

	s1.Add(2)
	s3.Add(3)

	got, err = s1.Intersection(s3)
	want = CreateSet(2)

	if !got.Equal(want) || err != nil {
		t.Errorf("The interesection of %v and %v resulted in %v, instead of %v.", s1, s3, got, want)
	}

	s4 := CreateSet("a")

	got, err = s1.Intersection(s4)
	want = NewSet()

	if !got.Equal(want) || err == nil {
		t.Errorf("The intersection of %v and %v succeeded with %v, instead of %v.", s1, s3, got, want)
	}
}
