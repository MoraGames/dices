package weightedDice

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

func TestWeightedSideInterface(t *testing.T) {
	// Check if WeightedSide implements Side interface
	side, err := NewWeightedSide("side", 1)
	if err != nil {
		t.Fatalf("NewWeightedSide() = %v\n", err)
	}

	standardSide := reflect.TypeOf(side)
	interfaceSide := reflect.TypeOf((*dices.Side)(nil)).Elem()
	if !standardSide.Implements(interfaceSide) {
		t.Fatalf("WeightedSide does not implement Side\n")
	}
	t.Logf("WeightedSide implements Side\n")
}

func TestNewWeightedSide(t *testing.T) {
	// Try with a negative weight (should be rejected)
	side, err := NewWeightedSide("side", -1)
	if err == nil {
		t.Fatalf("NewWeightedSide() = %v\n", err)
	}
	t.Logf("NewWeightedSide() = %+v\n", side)

	// Try with a positive weight (should be accepted)
	side, err = NewWeightedSide("side", 1)
	if err != nil {
		t.Fatalf("NewWeightedSide() = %v\n", err)
	}
	t.Logf("NewWeightedSide() = %+v\n", side)
}

func TestWeightedDiceInterface(t *testing.T) {
	// Check if WeightedDice implements Dice interface
	dice, err := NewDice([]float64{1, 1}, 2)
	if err != nil {
		t.Fatalf("NewDice() = %v\n", err)
	}

	standardDice := reflect.TypeOf(dice)
	interfaceDice := reflect.TypeOf((*dices.Dice)(nil)).Elem()
	if !standardDice.Implements(interfaceDice) {
		t.Fatalf("WeightedDice does not implement Dice\n")
	}
	t.Logf("WeightedDice implements Dice\n")
}

func TestNewDice(t *testing.T) {
	// Try with a less than the minimum number of sides (should be rejected)
	w1 := make([]float64, 0)
	for i := 0; i < dices.GetMinimumSides()-1; i++ {
		w1 = append(w1, 1)
	}

	d1, err := NewDice(w1, dices.GetMinimumSides()-1)
	if err == nil {
		t.Fatalf("NewDice() = %v\n", err)
	}
	t.Logf("NewDice() = %+v\n", d1)

	// Try with a greater than the minimum number of sides (should be accepted)
	w2 := make([]float64, 0)
	for i := 0; i < dices.GetMinimumSides()+1; i++ {
		w2 = append(w2, 1)
	}

	d2, err := NewDice(w2, dices.GetMinimumSides()+1)
	if err != nil {
		t.Fatalf("NewDice() = %v\n", err)
	}
	t.Logf("NewDice() = %+v\n", d2)
}

func TestNewRangeDice(t *testing.T) {
	// Try with a lowestSide value greater than highestSide value (should be rejected)
	invalidRange, err := dices.NewRange(1, -1)
	if err == nil {
		t.Fatalf("NewRange() = %v\n", err)
	}
	w1 := make([]float64, 0)
	for i := invalidRange.Lowest(); i <= invalidRange.Highest(); i++ {
		w1 = append(w1, 1)
	}

	d1, err := NewRangeDice(w1, invalidRange)
	if err == nil {
		t.Fatalf("NewRangeDice() = %v\n", err)
	}
	t.Logf("NewRangeDice() = %+v\n", d1)

	// Try with a lowestSide value smaller than highestSide value (should be accepted)
	validRange, err := dices.NewRange(-1, 1)
	if err != nil {
		t.Fatalf("NewRange() = %v\n", err)
	}
	w2 := make([]float64, 0)
	for i := validRange.Lowest(); i <= validRange.Highest(); i++ {
		w2 = append(w2, 1)
	}

	d2, err := NewRangeDice(w2, validRange)
	if err != nil {
		t.Fatalf("NewRangeDice() = %v\n", err)
	}
	t.Logf("NewRangeDice() = %+v\n", d2)
}

func TestNewCustomDice(t *testing.T) {
	// Try with a less than the minimum number of sides (should be rejected)
	var invalidWeights []float64
	var invalidSides []dices.Side
	for i := 0; i < dices.GetMinimumSides()-1; i++ {
		side, err := NewWeightedSide("side-"+strconv.Itoa(i), 1)
		if err != nil {
			t.Fatal("failed to create WeightedSide")
		}
		invalidWeights = append(invalidWeights, 1)
		invalidSides = append(invalidSides, side)
	}
	if len(invalidSides) != dices.GetMinimumSides()-1 {
		t.Fatal("invalidSides created with wrong length")
	}
	if len(invalidWeights) != len(invalidSides) {
		t.Fatal("invalidWeights created with wrong length")
	}

	d1, err := NewCustomDice(invalidWeights, invalidSides...)
	if err == nil {
		t.Fail()
	}
	t.Logf("NewCustomDice() = %+v, %v\n", d1, err)

	// Try with a greater than the minimum number of sides (should be accepted)
	var validWeights []float64
	var validSides []dices.Side
	for i := 0; i < dices.GetMinimumSides()+1; i++ {
		side, err := NewWeightedSide("side-"+strconv.Itoa(i), 1)
		if err != nil {
			t.Fatal("failed to create WeightedSide")
		}
		validWeights = append(validWeights, 1)
		validSides = append(validSides, side)
	}
	if len(validSides) != dices.GetMinimumSides()+1 {
		t.Fatal("validSides created with wrong length")
	}
	if len(validWeights) != len(validSides) {
		t.Fatal("validWeights created with wrong length")
	}

	d2, err := NewCustomDice(validWeights, validSides...)
	if err != nil {
		t.Fail()
	}
	t.Logf("NewCustomDice() = %+v, %v\n", d2, err)

	//Try with a custom struct data type (should be accepted)
	var validStructWeights []float64
	var validStructSides []dices.Side
	for i := 0; i < dices.GetMinimumSides()+1; i++ {
		side, err := NewWeightedSide(TestSide{i, "side-" + strconv.Itoa(i)}, 1)
		if err != nil {
			t.Fatal("failed to create WeightedSide")
		}
		validStructWeights = append(validStructWeights, 1)
		validStructSides = append(validStructSides, side)
	}
	if len(validStructSides) != dices.GetMinimumSides()+1 {
		t.Fatal("validStructSides created with wrong length")
	}
	if len(validStructWeights) != len(validStructSides) {
		t.Fatal("validStructWeights created with wrong length")
	}

	d3, err := NewCustomDice(validStructWeights, validStructSides...)
	if err != nil {
		t.Fail()
	}
	t.Logf("NewCustomDice() = %+v, %v\n", d3, err)
}
