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

	if !s.Exists(1) {
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

func TestExists(t *testing.T) {
	s := CreateSet(1)

	if !s.Exists(1) {
		t.Errorf("1 is not in the set %v", s)
	}

	if s.Exists(2) {
		t.Errorf("2 is in the set %v", s)
	}

	if s.Exists("3") {
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
