package dices

import (
	"testing"
)

func TestNewSet(t *testing.T) {
	// Try with a less than the minimum number of dices (should be rejected)
	var invalidDices []*Dice
	for i := 0; i < minimumDices-1; i++ {
		d, err := NewDice(minimumDices + 1)
		if err != nil {
			t.Fatal(err)
		}
		invalidDices = append(invalidDices, d)
	}
	if len(invalidDices) != minimumDices-1 {
		t.Fatal("invalidDices created with wrong length")
	}

	s1, err := NewSet(invalidDices...)
	if err == nil {
		t.Fail()
	}
	t.Logf("NewSet() = %+v, %v\n", s1, err)

	// Try with a greater than the minimum number of dices (should be accepted)
	var validDices []*Dice
	for i := 0; i < minimumDices+1; i++ {
		d, err := NewDice(minimumDices + 1)
		if err != nil {
			t.Fatal(err)
		}
		validDices = append(validDices, d)
	}
	if len(validDices) != minimumDices+1 {
		t.Fatal("validDices created with wrong length")
	}

	s2, err := NewSet(validDices...)
	if err != nil {
		t.Fail()
	}
	t.Logf("NewSet() = %+v, %v\n", s2, err)
}
