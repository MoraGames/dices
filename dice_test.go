package dice

import (
	"strconv"
	"testing"
)

func TestNewDice(t *testing.T) {
	// Try with an invalid number of sides (less than the minimum)
	d1, err := NewDice(minSides - 1)
	if err == nil {
		t.Fail()
	}
	t.Logf("NewDice() = %+v, %v\n", d1, err)

	//Try with a valid number of sides (greater than the minimum)
	d2, err := NewDice(minSides + 1)
	if err != nil {
		t.Fail()
	}
	t.Logf("NewDice() = %+v, %v\n", d2, err)
}

func TestNewCustomDice(t *testing.T) {
	//Prepare the sides for the invalid case
	var invalidSides []string
	for i := 0; i < minSides-1; i++ {
		invalidSides = append(invalidSides, strconv.Itoa(i)+"s")
	}
	if len(invalidSides) != minSides-1 {
		t.Fatal("invalidSides created with wrong length")
	}

	// Try with an invalid number of sides (less than the minimum)
	d1, err := NewCustomDice(invalidSides...)
	if err == nil {
		t.Fail()
	}
	t.Logf("NewDice() = %+v, %v\n", d1, err)

	//Prepare the sides for the valid case
	var validSides []string
	for i := 0; i < minSides+1; i++ {
		validSides = append(validSides, strconv.Itoa(i)+"s")
	}
	if len(validSides) != minSides+1 {
		t.Fatal("validSides created with wrong length")
	}

	//Try with a valid number of sides (greater than the minimum)
	d2, err := NewCustomDice(validSides...)
	if err != nil {
		t.Fail()
	}
	t.Logf("NewDice() = %+v, %v\n", d2, err)
}
