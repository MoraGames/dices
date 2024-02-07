package weightedDice

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/MoraGames/dices"
)

var _ dices.Side = new(WeightedSide)
var _ dices.Dice = new(WeightedDice)

type (
	WeightedSide struct {
		value  any
		weight float64
	}
	WeightedDice struct {
		rand *rand.Rand

		sidesNumber int
		sidesValue  []dices.Side
		//sidesWeight []float64
	}
)

func NewWeightedSide(value any, weight float64) (WeightedSide, error) {
	if weight < 0 {
		return WeightedSide{}, fmt.Errorf("the weight (%.2f) must be greater than or equal to 0", weight)
	}
	return WeightedSide{value, weight}, nil
}

func (s WeightedSide) Value() any {
	return s.value
}

func (s WeightedSide) Weight() float64 {
	return s.weight
}

func NewDice(sidesWeight []float64, sidesNumber int) (*WeightedDice, error) {
	if minimumSides := dices.GetMinimumSides(); sidesNumber < minimumSides {
		return nil, fmt.Errorf("the number of sides (%d) must be greater than or equal to %d", sidesNumber, minimumSides)
	}
	if len(sidesWeight) != sidesNumber {
		return nil, fmt.Errorf("the number of weight (%d) must be equal to sides number (%d)", len(sidesWeight), sidesNumber)
	}

	var sidesValue []dices.Side
	for s := 1; s <= sidesNumber; s++ {
		side, err := NewWeightedSide(s, sidesWeight[s-1])
		if err != nil {
			return nil, err
		}
		sidesValue = append(sidesValue, side)
	}
	return &WeightedDice{rand.New(rand.NewSource(time.Now().UnixNano())), sidesNumber, sidesValue}, nil
}

func NewRangeDice(sidesWeight []float64, sidesRange dices.Range) (*WeightedDice, error) {
	if sidesNumber := sidesRange.Highest() - sidesRange.Lowest() + 1; len(sidesWeight) != sidesNumber {
		return nil, fmt.Errorf("the number of weight (%d) must be equal to sides number (%d)", len(sidesWeight), sidesNumber)
	}

	var sidesValue []dices.Side
	for s := sidesRange.Lowest(); s <= sidesRange.Highest(); s++ {
		side, err := NewWeightedSide(s, sidesWeight[s-1])
		if err != nil {
			return nil, err
		}
		sidesValue = append(sidesValue, side)
	}
	return &WeightedDice{rand.New(rand.NewSource(time.Now().UnixNano())), len(sidesValue), sidesValue}, nil
}

func NewCustomDice(sidesWeight []float64, sidesValue ...dices.Side) (*WeightedDice, error) {
	if minimumSides := dices.GetMinimumSides(); len(sidesValue) < minimumSides {
		return nil, fmt.Errorf("the number of sides (%d) must be greater than or equal to %d", len(sidesValue), minimumSides)
	}
	if len(sidesWeight) != len(sidesValue) {
		return nil, fmt.Errorf("the number of weight (%d) must be equal to sides number (%d)", len(sidesWeight), len(sidesValue))
	}

	return &WeightedDice{rand.New(rand.NewSource(time.Now().UnixNano())), len(sidesValue), sidesValue}, nil
}

func (d *WeightedDice) Throw() dices.Side {
	return d.sidesValue[d.rand.Intn(d.sidesNumber)]
}

func (d *WeightedDice) ThrowN(n int) []dices.Side {
	result := make([]dices.Side, 0)
	for i := 0; i < n; i++ {
		result = append(result, d.Throw())
	}
	return result
}
