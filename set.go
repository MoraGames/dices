package dices

import "fmt"

const minimumDices int = 1

var defaultDices int = 2

func GetMinimumDices() int {
	return minimumDices
}

func GetDefaultDices() int {
	return defaultDices
}

func SetDefaultDices(numberDices int) error {
	if numberDices < minimumDices {
		return fmt.Errorf("the number of dices (%d) must be greater than or equal to %d", numberDices, minimumDices)
	}
	defaultDices = numberDices
	return nil
}

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
