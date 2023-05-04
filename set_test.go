package dice

import "testing"

func TestNewSet(t *testing.T) {
	d, err := NewDice(6)
	if err != nil {
		t.Fatal(err)
	}

	s1, err := NewSet()
	if err == nil {
		t.Fail()
	}
	t.Logf("NewSet() = %+v, %v\n", s1, err)

	s2, err := NewSet(d)
	if err != nil {
		t.Fail()
	}
	t.Logf("NewSet() = %+v, %v\n", s2, err)
}
