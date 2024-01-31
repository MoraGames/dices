package dices

import (
	"strconv"
	"testing"
)

func TestNewDice(t *testing.T) {
	// Try with a less than the minimum number of sides (should be rejected)
	d1, err := NewDice(minimumSides - 1)
	if err == nil {
		t.Fail()
	}
	t.Logf("NewDice() = %+v, %v\n", d1, err)

	// Try with a greater than the minimum number of sides (should be accepted)
	d2, err := NewDice(minimumSides + 1)
	if err != nil {
		t.Fail()
	}
	t.Logf("NewDice() = %+v, %v\n", d2, err)
}

func TestNewRangeDice(t *testing.T) {
	// Try with a lowestSide value greater than highestSide value (should be rejected)
	d1, err := NewRangeDice(3, 2)
	if err == nil {
		t.Fail()
	}
	t.Logf("NewRangeDice() = %+v, %v\n", d1, err)

	// Try with a lowestSide value smaller than highestSide value (should be accepted)
	d2, err := NewRangeDice(-1, 1)
	if err != nil {
		t.Fail()
	}
	t.Logf("NewRangeDice() = %+v, %v\n", d2, err)
}

func TestNewCustomDice(t *testing.T) {
	// Try with a less than the minimum number of sides (should be rejected)
	var invalidSides []Side
	for i := 0; i < minimumSides-1; i++ {
		invalidSides = append(invalidSides, strconv.Itoa(i)+"s")
	}
	if len(invalidSides) != minimumSides-1 {
		t.Fatal("invalidSides created with wrong length")
	}

	d1, err := NewCustomDice(invalidSides...)
	if err == nil {
		t.Fail()
	}
	t.Logf("NewCustomDice() = %+v, %v\n", d1, err)

	// Try with a greater than the minimum number of sides (should be accepted)
	var validSides []Side
	for i := 0; i < minimumSides+1; i++ {
		validSides = append(validSides, strconv.Itoa(i)+"s")
	}
	if len(validSides) != minimumSides+1 {
		t.Fatal("validSides created with wrong length")
	}

	d2, err := NewCustomDice(validSides...)
	if err != nil {
		t.Fail()
	}
	t.Logf("NewCustomDice() = %+v, %v\n", d2, err)

	//Try with a custom struct data type (should be accepted)
	type TempSide struct {
		v1 int
		v2 string
	}

	var validStructSides []Side
	for i := 0; i < minimumSides+1; i++ {
		validStructSides = append(validStructSides, TempSide{i, strconv.Itoa(i) + "s"})
	}
	if len(validStructSides) != minimumSides+1 {
		t.Fatal("validStructSides created with wrong length")
	}

	d3, err := NewCustomDice(validStructSides...)
	if err != nil {
		t.Fail()
	}
	t.Logf("NewCustomDice() = %+v, %v\n", d3, err)
}
