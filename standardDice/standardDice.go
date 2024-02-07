package standardDice

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/MoraGames/dices"
)

var _ dices.Side = new(StandardSide)
var _ dices.Dice = new(StandardDice)

type (
	StandardSide struct {
		value any
	}
	StandardDice struct {
		rand *rand.Rand

		sidesNumber int
		sidesValue  []dices.Side
	}
)

func NewStandardSide(value any) StandardSide {
	return StandardSide{value}
}

func (s StandardSide) Value() any {
	return s.value
}

func NewDice(sidesNumber int) (*StandardDice, error) {
	if minimumSides := dices.GetMinimumSides(); sidesNumber < minimumSides {
		return nil, fmt.Errorf("the number of sides (%d) must be greater than or equal to %d", sidesNumber, minimumSides)
	}

	var sidesValue []dices.Side
	for s := 1; s <= sidesNumber; s++ {
		sidesValue = append(sidesValue, NewStandardSide(s))
	}
	return &StandardDice{rand.New(rand.NewSource(time.Now().UnixNano())), sidesNumber, sidesValue}, nil
}

func NewRangeDice(sidesRange dices.Range) (*StandardDice, error) {
	var sidesValue []dices.Side
	for s := sidesRange.Lowest(); s <= sidesRange.Highest(); s++ {
		sidesValue = append(sidesValue, NewStandardSide(s))
	}
	return &StandardDice{rand.New(rand.NewSource(time.Now().UnixNano())), len(sidesValue), sidesValue}, nil
}

func NewCustomDice(sidesValue ...dices.Side) (*StandardDice, error) {
	if minimumSides := dices.GetMinimumSides(); len(sidesValue) < minimumSides {
		return nil, fmt.Errorf("the number of sides (%d) must be greater than or equal to %d", len(sidesValue), minimumSides)
	}

	return &StandardDice{rand.New(rand.NewSource(time.Now().UnixNano())), len(sidesValue), sidesValue}, nil
}

func (d *StandardDice) Throw() dices.Side {
	return d.sidesValue[d.rand.Intn(d.sidesNumber)]
}

func (d *StandardDice) ThrowN(n int) []dices.Side {
	result := make([]dices.Side, 0)
	for i := 0; i < n; i++ {
		result = append(result, d.Throw())
	}
	return result
}
