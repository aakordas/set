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
	if !s2.Equal(s1) {
		t.Errorf("The set %v is not equal to the set %v.", s2, s1)
	}

	s2.Add(2)

	if s1.Equal(s2) {
		t.Errorf("The set %v is equal to the set %v.", s1, s2)
	}
	if s2.Equal(s1) {
		t.Errorf("The set %v is equal to the set %v.", s2, s1)
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
	if !s2.Subset(s1) {
		t.Errorf("The set %v is not a subset of the set %v.", s2, s1)
	}

	s1.Add(1)

	if s1.Subset(s2) {
		t.Errorf("The set %v is a subset of the set %v.", s1, s2)
	}
	if !s2.Subset(s1) {
		t.Errorf("The set %v is not a subset of the set %v.", s2, s1)
	}

	s2.Add(1)

	if !s1.Subset(s2) {
		t.Errorf("The set %v is not a subset of the set %v.", s1, s2)
	}
	if !s2.Subset(s1) {
		t.Errorf("The set %v is not a subset of the set %v.", s2, s1)
	}

	s2.Add(2)

	if !s1.Subset(s2) {
		t.Errorf("The set %v is not a subset of the set %v.", s1, s2)
	}

	s1.Add(3)

	if s1.Subset(s2) {
		t.Errorf("The set %v is a subset of the set %v.", s1, s2)
	}
	if s2.Subset(s1) {
		t.Errorf("The set %v is a subset of the set %v.", s2, s1)
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

	got1, err := s1.Union(s2)
	got2, err := s2.Union(s1)

	want := CreateSet(1)
	want.Add(2)

	if !got1.Equal(want) || err != nil {
		t.Errorf("The union of %v and %v resulted in %v, instead of %v.", s1, s2, got1, want)
	}
	if !got1.Equal(got2) {
		t.Errorf("Union is not commutative!")
	}

	s3 := CreateSet("a")

	got, err := s1.Union(s3)
	want = Set{}

	if !got.Equal(want) || err == nil {
		t.Errorf("The union of %v and %v succeeded with %v, instead of %v.", s1, s3, got, want)
	}
}

func TestIntersection(t *testing.T) {
	s1 := CreateSet(1)
	s2 := CreateSet(1)

	got1, err := s1.Intersection(s2)
	got2, err := s2.Intersection(s1)
	want := CreateSet(1)

	if !got1.Equal(want) || err != nil {
		t.Errorf("The intersection of %v and %v resulted in %v, instead of %v.", s1, s2, got1, want)
	}
	if !got1.Equal(got2) {
		t.Errorf("Intersection is not commutative!")
	}

	s3 := CreateSet(2)

	got1, err = s1.Intersection(s3)
	got2, err = s3.Intersection(s1)
	want = NewSet()

	if !got1.Equal(want) || err != nil {
		t.Errorf("The intersection of %v and %v resulted in %v, instead of %v.", s1, s3, got1, want)
	}
	if !got1.Equal(got2) {
		t.Errorf("Intersection is not commutative!")
	}

	s1.Add(2)
	s3.Add(3)

	got1, err = s1.Intersection(s3)
	got2, err = s3.Intersection(s1)
	want = CreateSet(2)

	if !got1.Equal(want) || err != nil {
		t.Errorf("The interesection of %v and %v resulted in %v, instead of %v.", s1, s3, got1, want)
	}
	if !got1.Equal(got2) {
		t.Errorf("Intersection is not commutative!")
	}

	s4 := CreateSet("a")

	got, err := s1.Intersection(s4)
	want = NewSet()

	if !got.Equal(want) || err == nil {
		t.Errorf("The intersection of %v and %v succeeded with %v, instead of %v.", s1, s3, got, want)
	}
}

func TestUnionSubsetProperties(t *testing.T) {
	s1 := CreateSet(1)
	s2 := CreateSet(2)

	union, err := s1.Union(s2)
	if err != nil {
		t.Errorf("There was an error trying to make the union of %v and %v.\n%v", s1, s2, err)
	}

	// s1 ⊆ s1 ∪ s2
	if !s1.Subset(union) {
		t.Errorf("The set %v is not a subset of the set %v.", s1, union)
	}
	// s2 ⊆ s1 ∪ s2
	if !s2.Subset(union) {
		t.Errorf("The set %v is not a subset of the set %v.", s2, union)
	}

	s3 := CreateSet(1)
	s3.Add(2)

	// if s1 ⊆ s3 and s2 ⊆ s3 then s1 ∪ s2 ⊆ s3
	if !union.Subset(s3) {
		t.Errorf("The set %v is not a subset of the set %v.", union, s3)
	}

	// s1 ∪ s1 = s1
	union, err = s1.Union(s1)
	if err != nil {
		t.Errorf("There was an error trying to make the union of %v with itself.\n%v", s1, err)
	}

	if !union.Equal(s1) {
		t.Errorf("The set %v is not the same as the set %v.", union, s1)
	}

	// s1 ∪ s2 = s2 ∪ s1
	union1, err1 := s1.Union(s2)
	if err1 != nil {
		t.Errorf("There was an error trying to make the union of %v with %v.", s1, s2)
	}

	union2, err2 := s2.Union(s1)
	if err2 != nil {
		t.Errorf("There was an error trying to make the union of %v with %v.", s2, s1)
	}

	if !union1.Equal(union2) {
		t.Errorf("The set %v is not equal to %v, that means the union is not commutative.", union1, union2)
	}

	// (s1 ∪ s2) ∪ s3 = s1 ∪ (s2 ∪ s3)
	s3 = CreateSet(3)
	union1, err = s1.Union(s2)
	if err != nil {
		t.Errorf("There was an error trying to make the union of %v with %v.", s1, s2)
	}

	got1, err := union1.Union(s3)
	if err != nil {
		t.Errorf("There was an error trying to make the union of %v with %v.", union1, s3)
	}

	union2, err = s2.Union(s3)
	if err != nil {
		t.Errorf("There was an error trying to make the union of %v with %v.", s2, s3)
	}

	got2, err := s1.Union(union2)
	if err != nil {
		t.Errorf("There was an error trying to make the union of %v with %v.", s1, union2)
	}

	if !got1.Equal(got2) {
		t.Errorf("The set %v is not equal to the set %v. This means union is not associative.", got1, got2)
	}
}

func TestIntersectionSubsetProperties(t *testing.T) {
	s1 := CreateSet(1)
	s1.Add(2)
	s2 := CreateSet(1)
	s2.Add(3)

	intersection, err := s1.Intersection(s2)
	if err != nil {
		t.Errorf("There was an error trying to make the intersection of %v and %v.\n%v", s1, s2, err)
	}

	// s1 ∩ s2 ⊆ s1
	if !intersection.Subset(s1) {
		t.Errorf("The set %v is not a subset of the set %v.", intersection, s1)
	}
	// s1 ∩ s2 ⊆ s2
	if !intersection.Subset(s2) {
		t.Errorf("The set %v is not a subset of the set %v.", intersection, s2)
	}

	// if s3 ⊆ s1 and s3 ⊆ s2 then s3 ⊆ s1 ∩ s2
	s3 := CreateSet(1)

	if !s3.Subset(s1) {
		t.Errorf("The set %v is not a subset of the set %v.", s3, s1)
	}
	if !s3.Subset(s2) {
		t.Errorf("The set %v is not a subset of the set %v.", s3, s2)
	}

	if !s3.Subset(intersection) {
		t.Errorf("The set %v is not a subset of the set %v.", s3, intersection)
	}

	// s1 ∩ s1 = s1
	intersection, err = s1.Intersection(s1)
	if err != nil {
		t.Errorf("There was an error trying to make the intersection of %v with itself.", s1)
	}

	if !intersection.Equal(s1) {
		t.Errorf("The set %v is not the same as the set %v.", intersection, s1)
	}

	// s1 ∩ s2 = s2 ∩ s1
	intersection1, err1 := s1.Intersection(s2)
	if err1 != nil {
		t.Errorf("There was an error trying to make the intersection of %v with %v.", s1, s2)
	}

	intersection2, err2 := s2.Intersection(s1)
	if err2 != nil {
		t.Errorf("There was an error trying to make the intersection of %v with %v.", s2, s1)
	}

	if !intersection1.Equal(intersection2) {
		t.Errorf("The set %v is not equal to %v, that means the intersection is not commutative.", intersection1, intersection2)
	}

	// (s1 ∩ s2) ∩ s3 = s1 ∩ (s2 ∩ s3)
	s3 = CreateSet(1)
	s3.Add(4)
	intersection1, err = s1.Intersection(s2)
	if err != nil {
		t.Errorf("There was an error trying to make the intersection of %v with %v.", s1, s2)
	}

	got1, err := intersection1.Intersection(s3)
	if err != nil {
		t.Errorf("There was an error trying to make the intersection of %v with %v.", intersection1, s3)
	}

	intersection2, err = s2.Intersection(s3)
	if err != nil {
		t.Errorf("There was an error trying to make the intersection of %v with %v.", s2, s3)
	}

	got2, err := s1.Intersection(intersection2)
	if err != nil {
		t.Errorf("There was an error trying to make the interesection of %v with %v.", s1, intersection2)
	}

	if !got1.Equal(got2) {
		t.Errorf("The set %v is not equal to the set %v. This means intersection is not associative.", got1, got2)
	}
}

func TestSetProperties(t *testing.T) {
	// s1 ∩ (s2 ∪ s3) = (s1 ∩ s2) ∪ (s1 ∩ s3)
	s1 := CreateSet(1)
	s1.Add(2)
	s1.Add(3)
	s2 := CreateSet(2)
	s3 := CreateSet(3)

	union, err := s2.Union(s3)
	if err != nil {
		t.Errorf("There was an error trying to make the union of %v and %v.", s2, s3)
	}

	got1, err := s1.Intersection(union)
	if err != nil {
		t.Errorf("There was an error trying to make the intersection of %v and %v.", s1, union)
	}

	intersection1, err := s1.Intersection(s2)
	if err != nil {
		t.Errorf("There was an error trying to make the intersection of %v and %v.", s1, s2)
	}

	intersection2, err := s1.Intersection(s3)
	if err != nil {
		t.Errorf("There was an error trying to make the intersection of %v and %v.", s1, s3)
	}

	got2, err := intersection1.Union(intersection2)
	if err != nil {
		t.Errorf("There was an error trying to make the union of %v and %v.", intersection1, intersection2)
	}

	if !got1.Equal(got2) {
		t.Errorf("The set %v is not equal to the set %v. This means the operations are not distributive.", got1, got2)
	}

	// s1 ∪ (s2 ∩ s3) = (s1 ∪ s2) ∩ (s1 ∪ s3)
	intersection, err := s2.Intersection(s3)
	if err != nil {
		t.Errorf("There was an error trying to make the intersection of %v and %v.", s2, s3)
	}

	got1, err = s1.Union(intersection)
	if err != nil {
		t.Errorf("There was an error trying to make the union of %v and %v.", s1, intersection)
	}

	union1, err := s1.Union(s2)
	if err != nil {
		t.Errorf("There was an error trying to make the union of %v and %v.", s1, s2)
	}

	union2, err := s1.Union(s3)
	if err != nil {
		t.Errorf("There was an error trying to make the union of %v and %v.", s1, s3)
	}

	got2, err = union1.Intersection(union2)
	if err != nil {
		t.Errorf("There was an error trying to make the intersection of %v and %v.", union1, union2)
	}

	if !got1.Equal(got2) {
		t.Errorf("The set %v is not equal to the set %v. This means the operations are not distributive.", got1, got2)
	}

	// s1 ∩ (s1 ∪ s2) = s1
	union, err = s1.Union(s2)
	if err != nil {
		t.Errorf("There was an error trying to make the union of %v and %v.", s1, s2)
	}

	got, err := s1.Intersection(union)
	if err != nil {
		t.Errorf("There was an error trying to make the intersection of %v and %v.", s1, union)
	}

	if !got.Equal(s1) {
		t.Errorf("The set %v is not equal to the set %v. This means the operations are not absorptive.", got, s1)
	}
}

func TestDifference(t *testing.T) {
	s1 := CreateSet(1)
	s1.Add(2)
	s2 := CreateSet(2)

	got, err := s1.Difference(s2)
	if err != nil {
		t.Errorf("There was an error trying to make the difference of %v and %v.", s1, s2)
	}

	want := CreateSet(1)

	if !got.Equal(want) {
		t.Errorf("The difference of set %v from set %v is %v instead of %v.", s2, s1, got, want)
	}

	got, err = s2.Difference(s1)
	want = NewSet()

	if !got.Equal(want) {
		t.Errorf("The difference of set %v from set %v is %v instead of %v.", s1, s2, got, want)
	}
}
