package set

import (
	"fmt"

	"github.com/MoraGames/dices"
)

const minimumDices int = 1

func GetMinimumDices() int {
	return minimumDices
}

type Set struct {
	dices []dices.Dice
}

func NewSet(dices ...dices.Dice) (*Set, error) {
	if len(dices) == 0 {
		return nil, fmt.Errorf("the number of dices (%d) must be greater than or equal to %d", len(dices), minimumDices)
	}
	return &Set{dices}, nil
}

func (s *Set) Throw() []dices.Side {
	result := make([]dices.Side, 0)
	for i := 0; i < len(s.dices); i++ {
		result = append(result, s.dices[i].Throw())
	}
	return result
}
