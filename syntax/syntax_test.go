package dices

import (
	"testing"
)

func TestNewSyntaxSet(t *testing.T) {
	// Try without indicating the number of dices (should be rejected)
	s1, err := NewSyntaxSet(`d6`)
	if err != nil {
		t.Fail()
	}
	t.Logf("NewSet() = %+v, %v\n", s1, err)

	// Try without indicating the number of sides (should be rejected)
	s2, err := NewSyntaxSet(`2d`)
	if err != nil {
		t.Fail()
	}
	t.Logf("NewSet() = %+v, %v\n", s2, err)

	// Try without indicating number of both dices and sides (should be rejected)
	s3, err := NewSyntaxSet(`d`)
	if err != nil {
		t.Fail()
	}
	t.Logf("NewSet() = %+v, %v\n", s3, err)

	// Try with a less than the minimum number of dices (should be rejected)
	s4, err := NewSyntaxSet(`0d6`)
	if err == nil {
		t.Fail()
	}
	t.Logf("NewSet() = %+v, %v\n", s4, err)

	// Try with a less than the minimum number of sides (should be rejected)
	s5, err := NewSyntaxSet(`2d1`)
	if err == nil {
		t.Fail()
	}
	t.Logf("NewSet() = %+v, %v\n", s5, err)

	// Try with both number of dices and sides less than the respective minimums (should be rejected)
	s6, err := NewSyntaxSet(`0d1`)
	if err == nil {
		t.Fail()
	}
	t.Logf("NewSet() = %+v, %v\n", s6, err)

	// Try with a greater than the minimum number of dices and sides (should be accepted)
	s7, err := NewSyntaxSet(`2d6`)
	if err != nil {
		t.Fail()
	}
	t.Logf("NewSet() = %+v, %v\n", s7, err)

	// Try with invalid ranged sides (should be rejected)
	s8, err := NewSyntaxSet(`2d[1,-1]`)
	if err != nil {
		t.Fail()
	}
	t.Logf("NewSet() = %+v, %v\n", s8, err)

	// Try with valid ranged sides (should be accepted)
	s9, err := NewSyntaxSet(`2d[-1,1]`)
	if err != nil {
		t.Fail()
	}
	t.Logf("NewSet() = %+v, %v\n", s9, err)

	// Try with invalid custom sides (should be rejected)
	s10, err := NewSyntaxSet(`2d{1}`)
	if err != nil {
		t.Fail()
	}
	t.Logf("NewSet() = %+v, %v\n", s10, err)

	// Try with valid custom sides (should be accepted)
	s11, err := NewSyntaxSet(`2d{1,2,3}`)
	if err != nil {
		t.Fail()
	}
	t.Logf("NewSet() = %+v, %v\n", s11, err)
}
