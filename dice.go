package dices

import (
	"fmt"
	"math/rand"
	"time"
)

const minimumSides int = 2

var defaultSides int = 6

func GetMinimumSides() int {
	return minimumSides
}

func GetDefaultSides() int {
	return defaultSides
}

func SetDefaultSides(numberSides int) error {
	if numberSides < minimumSides {
		return fmt.Errorf("the number of sides (%d) must be greater than or equal to %d", numberSides, minimumSides)
	}
	defaultSides = numberSides
	return nil
}

type (
	Side any
	Dice struct {
		rand *rand.Rand

		sidesNumber int
		sidesValue  []Side
	}
)

func NewDice(sidesNumber int) (*Dice, error) {
	if sidesNumber < minimumSides {
		return nil, fmt.Errorf("the number of sides (%d) must be greater than or equal to %d", sidesNumber, minimumSides)
	}
	var sidesValue []Side
	for s := 1; s <= sidesNumber; s++ {
		sidesValue = append(sidesValue, s)
	}
	return &Dice{rand.New(rand.NewSource(time.Now().UnixNano())), sidesNumber, sidesValue}, nil
}

func NewRangeDice(lowestSide, highestSide int) (*Dice, error) {
	if lowestSide >= highestSide {
		return nil, fmt.Errorf("the lowestSide (%d) must be less than highestSide (%d)", lowestSide, highestSide)
	}
	var sidesValue []Side
	for s := lowestSide; s <= highestSide; s++ {
		sidesValue = append(sidesValue, s)
	}
	return &Dice{rand.New(rand.NewSource(time.Now().UnixNano())), len(sidesValue), sidesValue}, nil
}

func NewCustomDice(sidesValue ...Side) (*Dice, error) {
	if len(sidesValue) < minimumSides {
		return nil, fmt.Errorf("the number of sides (%d) must be greater than or equal to %d", len(sidesValue), minimumSides)
	}
	return &Dice{rand.New(rand.NewSource(time.Now().UnixNano())), len(sidesValue), sidesValue}, nil
}

func (d *Dice) Throw() Side {
	return d.sidesValue[d.rand.Intn(d.sidesNumber)]
}
