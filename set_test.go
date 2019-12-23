package set

import (
	"testing"
)

func TestCreateAndAdd(t *testing.T) {
	s := CreateSet()

	if !s.Add(1) {
		t.Errorf("1 already exists in the set %v", s)
	}
}

func TestCreateWithElement(t *testing.T) {
	s := Create(1)

	if !s.Exists(1) {
		t.Errorf("1 is not in the set %v", s)
	}
}

func TestAddMultipleElements(t *testing.T) {
	s := CreateSet()

	if !s.Add(1) {
		t.Errorf("1 already exists in the set %v", s)
	}

	if !s.Add(2) {
		t.Errorf("2 already exists in the set %v", s)
	}
}

func TestAddDifferentKindElements(t *testing.T) {
	s := CreateSet()

	if !s.Add(1) {
		t.Errorf("1 already exists in the set %v", s)
	}

	if !s.Add("2") {
		t.Errorf("\"2\" already exists in the set %v", s)
	}
}

func TestAddDifferentKindElementsThanInitial(t *testing.T) {
	s := Create(1)

	if !s.Add(2) {
		t.Errorf("2 already exists in the set %v", s)
	}

	if s.Add("2") {
		t.Errorf("\"2\" was succesfully added in the set %v", s)
	}
}

func TestAddDuplicate(t *testing.T) {
	s := CreateSet()

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

	s := Create(points[0])

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
	s := CreateSet()
	s.Add(1)

	if !s.Exists(1) {
		t.Errorf("1 is not in the set %v", s)
	}

	if s.Exists(2) {
		t.Errorf("2 is in the set %v", s)
	}
}

func TestEmpty(t *testing.T) {
	s := CreateSet()

	if !s.Empty() {
		t.Errorf("The set %v is not empty", s)
	}

	s.Add(1)

	if s.Empty() {
		t.Errorf("The set %v is empty", s)
	}
}
