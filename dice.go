package dices

import (
	"fmt"
	"math/rand"
	"time"
)

const minimumSides int = 2

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

func NewCustomDice(sidesValue ...Side) (*Dice, error) {
	if len(sidesValue) < minimumSides {
		return nil, fmt.Errorf("the number of sides (%d) must be greater than or equal to %d", len(sidesValue), minimumSides)
	}
	return &Dice{rand.New(rand.NewSource(time.Now().UnixNano())), len(sidesValue), sidesValue}, nil
}

func (d *Dice) Throw() Side {
	return d.sidesValue[d.rand.Intn(d.sidesNumber)]
}
