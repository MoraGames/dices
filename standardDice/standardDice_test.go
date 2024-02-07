package standardDice

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/MoraGames/dices"
)

// TestSide is struct that implement Side interface used to test custom data types
type TestSide struct {
	v1 int
	v2 string
}

func (s TestSide) Value() interface{} {
	return s
}

func TestStandardSideInterface(t *testing.T) {
	// Check if StandardSide implements Side interface
	standardSide := reflect.TypeOf(NewStandardSide(1))
	interfaceSide := reflect.TypeOf((*dices.Side)(nil)).Elem()
	if !standardSide.Implements(interfaceSide) {
		t.Fatalf("StandardSide does not implement Side\n")
	}
	t.Logf("StandardSide implements Side\n")
}

func TestStandardDiceInterface(t *testing.T) {
	// Check if StandardDice implements Dice interface
	dice, err := NewDice(6)
	if err != nil {
		t.Fatalf("NewDice() = %v\n", err)
	}

	standardDice := reflect.TypeOf(dice)
	interfaceDice := reflect.TypeOf((*dices.Dice)(nil)).Elem()
	if !standardDice.Implements(interfaceDice) {
		t.Fatalf("StandardDice does not implement Dice\n")
	}
	t.Logf("StandardDice implements Dice\n")
}

func TestNewDice(t *testing.T) {
	// Try with a less than the minimum number of sides (should be rejected)
	d1, err := NewDice(dices.GetMinimumSides() - 1)
	if err == nil {
		t.Fatalf("NewDice() = %v\n", err)
	}
	t.Logf("NewDice() = %+v\n", d1)

	// Try with a greater than the minimum number of sides (should be accepted)
	d2, err := NewDice(dices.GetMinimumSides() + 1)
	if err != nil {
		t.Fatalf("NewDice() = %v\n", err)
	}
	t.Logf("NewDice() = %+v\n", d2)
}

func TestNewRangeDice(t *testing.T) {
	// Try with a lowestSide value greater than highestSide value (should be rejected)
	invalidRange, err := dices.NewRange(1, -1)
	if err != nil {
		t.Fatalf("NewRange() = %v\n", err)
	}

	d1, err := NewRangeDice(invalidRange)
	if err == nil {
		t.Fatalf("NewRangeDice() = %v\n", err)
	}
	t.Logf("NewRangeDice() = %+v\n", d1)

	// Try with a lowestSide value smaller than highestSide value (should be accepted)
	validRange, err := dices.NewRange(-1, 1)
	if err != nil {
		t.Fatalf("NewRange() = %v\n", err)
	}

	d2, err := NewRangeDice(validRange)
	if err != nil {
		t.Fatalf("NewRangeDice() = %v\n", err)
	}
	t.Logf("NewRangeDice() = %+v\n", d2)
}

func TestNewCustomDice(t *testing.T) {
	// Try with a less than the minimum number of sides (should be rejected)
	var invalidSides []dices.Side
	for i := 0; i < dices.GetMinimumSides()-1; i++ {
		invalidSides = append(invalidSides, NewStandardSide("side-"+strconv.Itoa(i)))
	}
	if len(invalidSides) != dices.GetMinimumSides()-1 {
		t.Fatal("invalidSides created with wrong length")
	}

	d1, err := NewCustomDice(invalidSides...)
	if err == nil {
		t.Fail()
	}
	t.Logf("NewCustomDice() = %+v, %v\n", d1, err)

	// Try with a greater than the minimum number of sides (should be accepted)
	var validSides []dices.Side
	for i := 0; i < dices.GetMinimumSides()+1; i++ {
		validSides = append(validSides, NewStandardSide("side-"+strconv.Itoa(i)))
	}
	if len(validSides) != dices.GetMinimumSides()+1 {
		t.Fatal("validSides created with wrong length")
	}

	d2, err := NewCustomDice(validSides...)
	if err != nil {
		t.Fail()
	}
	t.Logf("NewCustomDice() = %+v, %v\n", d2, err)

	//Try with a custom struct data type (should be accepted)
	var validStructSides []dices.Side
	for i := 0; i < dices.GetMinimumSides()+1; i++ {
		validStructSides = append(validStructSides, NewStandardSide(TestSide{i, "side-" + strconv.Itoa(i)}))
	}
	if len(validStructSides) != dices.GetMinimumSides()+1 {
		t.Fatal("validStructSides created with wrong length")
	}

	d3, err := NewCustomDice(validStructSides...)
	if err != nil {
		t.Fail()
	}
	t.Logf("NewCustomDice() = %+v, %v\n", d3, err)
}
