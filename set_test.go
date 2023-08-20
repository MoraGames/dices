package dices

import (
	"testing"
)

func TestNewSet(t *testing.T) {
	//Prepare the sides for the invalid case
	var invalidDices []*Dice
	for i := 0; i < minDices-1; i++ {
		d, err := NewDice(minSides + 1)
		if err != nil {
			t.Fatal(err)
		}
		invalidDices = append(invalidDices, d)
	}
	if len(invalidDices) != minDices-1 {
		t.Fatal("invalidDices created with wrong length")
	}

	// Try with an invalid number of dices (less than the minimum)
	s1, err := NewSet(invalidDices...)
	if err == nil {
		t.Fail()
	}
	t.Logf("NewSet() = %+v, %v\n", s1, err)

	//Prepare the sides for the valid case
	var validDices []*Dice
	for i := 0; i < minDices+1; i++ {
		d, err := NewDice(minSides + 1)
		if err != nil {
			t.Fatal(err)
		}
		validDices = append(validDices, d)
	}
	if len(validDices) != minDices+1 {
		t.Fatal("validDices created with wrong length")
	}

	//Try with a valid number of dices (greater than the minimum)
	s2, err := NewSet(validDices...)
	if err != nil {
		t.Fail()
	}
	t.Logf("NewSet() = %+v, %v\n", s2, err)
}
