package dices

import "fmt"

const minimumDices int = 1

type Set struct {
	dices []*Dice
}

func NewSet(dices ...*Dice) (*Set, error) {
	if len(dices) == 0 {
		return nil, fmt.Errorf("the number of dices (%d) must be greater than or equal to %d", len(dices), minimumDices)
	}
	return &Set{dices}, nil
}

func (s *Set) Throw() []Side {
	result := make([]Side, 0)
	for i := 0; i < len(s.dices); i++ {
		result = append(result, s.dices[i].Throw())
	}
	return result
}
