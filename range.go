package dices

import "fmt"

type (
	Range struct {
		lowest  int
		highest int
	}
)

func NewRange(lowest, highest int) (Range, error) {
	if lowest >= highest {
		return Range{}, fmt.Errorf("the lowest value (%d) must be less than highest value (%d)", lowest, highest)
	}
	return Range{lowest, highest}, nil
}

func (r Range) Lowest() int {
	return r.lowest
}

func (r Range) Highest() int {
	return r.highest
}
