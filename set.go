package dice

import "fmt"

const minDices int = 1

type Set struct {
	dices []*Dice
}

func NewSet(dices ...*Dice) (*Set, error) {
	if len(dices) == 0 {
		return nil, fmt.Errorf("the number of dices (%d) must be greater than or equal to %d", len(dices), minDices)
	}
	return &Set{dices}, nil
}

func (s *Set) Throw() []string {
	result := make([]string, 0)
	for i := 0; i < len(s.dices); i++ {
		result = append(result, s.dices[i].Throw())
	}
	return result
}
