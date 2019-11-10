package set

import (
	"testing"
)

func TestCreateAndAdd(t *testing.T) {
	s := Create()

	if s.Add(1) == false {
		t.Errorf("1 already exists in the set %v", s)
	}
}

func TestAddMultipleElements(t *testing.T) {
	s := Create()

	if s.Add(1) == false {
		t.Errorf("1 already exists in the set %v", s)
	}

	if s.Add(2) == false {
		t.Errorf("2 already exists in the set %v", s)
	}
}

func TestAddDifferentKindElements(t *testing.T) {
	s := Create()

	if s.Add(1) == false {
		t.Errorf("1 already exists in the set %v", s)
	}

	if s.Add("2") == false {
		t.Errorf("\"2\" already exists in the set %v", s)
	}
}

func TestAddDuplicate(t *testing.T) {
	s := Create()

	if s.Add(1) == false {
		t.Errorf("1 already exists in the set %v", s)
	}

	if s.Add(1) == true {
		t.Errorf("1 does not exist in the set %v", s)
	}
}

func TestExists(t *testing.T) {
	s := Create()
	s.Add(1)

	if s.Exists(1) == false {
		t.Errorf("1 is not in the set %v", s)
	}

	if s.Exists(2) == true {
		t.Errorf("2 is in the set %v", s)
	}
}

func TestEmpty(t *testing.T) {
	s := Create()

	if s.Empty() == false {
		t.Errorf("The set %v is not empty", s)
	}

	s.Add(1)

	if s.Empty() == true {
		t.Errorf("The set %v is empty", s)
	}
}

func TestSetInSet(t *testing.T) {
	s1 := Create()
	s2 := Create()

	if s1.Add(s2) == false {
		t.Errorf("%v already exists in the set %v", s2, s1)
	}
}
