package dices

import (
	"testing"
)

func TestNewSyntaxSet(t *testing.T) {
	// Try with a less than the minimum number of dices (should be rejected)
	s1, err := NewSyntaxSet(`0d6`)
	if err == nil {
		t.Fail()
	}
	t.Logf("NewSet() = %+v, %v\n", s1, err)

	// Try with a less than the minimum number of sides (should be rejected)
	s2, err := NewSyntaxSet(`2d1`)
	if err == nil {
		t.Fail()
	}
	t.Logf("NewSet() = %+v, %v\n", s2, err)

	// Try with a greater than the minimum number of dices and sides (should be accepted)
	s3, err := NewSyntaxSet(`2d6`)
	if err != nil {
		t.Fail()
	}
	t.Logf("NewSet() = %+v, %v\n", s3, err)

	// Try with the default number of dices (should be accepted)
	s4, err := NewSyntaxSet(`d6`)
	if err != nil {
		t.Fail()
	}
	t.Logf("NewSet() = %+v, %v\n", s4, err)

	// Try with the default number of sides (should be accepted)
	s5, err := NewSyntaxSet(`2d`)
	if err != nil {
		t.Fail()
	}
	t.Logf("NewSet() = %+v, %v\n", s5, err)

	// Try with the default number of dices and sides (should be accepted)
	s6, err := NewSyntaxSet(`d`)
	if err != nil {
		t.Fail()
	}
	t.Logf("NewSet() = %+v, %v\n", s6, err)

	// Try with invalid ranged sides (should be rejected)
	s7, err := NewSyntaxSet(`2d[1-0]`)
	if err != nil {
		t.Fail()
	}
	t.Logf("NewSet() = %+v, %v\n", s7, err)

	// Try with valid ranged sides (should be accepted)
	s8, err := NewSyntaxSet(`2d[-1,1]`)
	if err != nil {
		t.Fail()
	}
	t.Logf("NewSet() = %+v, %v\n", s8, err)
}
