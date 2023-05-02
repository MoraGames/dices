package dice

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const minSides int = 2

type Dice struct {
	rand *rand.Rand

	sidesNumber int
	sidesValue  []string
}

func NewDice(sidesNumber int) (*Dice, error) {
	if sidesNumber < minSides {
		return nil, fmt.Errorf("the number of faces (%d) must be greater than or equal to %d", sidesNumber, minSides)
	}
	var sidesValue []string
	for s := 1; s <= sidesNumber; s++ {
		sidesValue = append(sidesValue, strconv.Itoa(s))
	}
	return &Dice{rand.New(rand.NewSource(time.Now().UnixNano())), sidesNumber, sidesValue}, nil
}

func NewCustomDice(sidesValue []string) (*Dice, error) {
	if len(sidesValue) < minSides {
		return nil, fmt.Errorf("the number of faces (%d) must be greater than or equal to %d", len(sidesValue), minSides)
	}
	return &Dice{rand.New(rand.NewSource(time.Now().UnixNano())), len(sidesValue), sidesValue}, nil
}

func (d *Dice) Throw() string {
	return d.sidesValue[d.rand.Intn(d.sidesNumber)]
}
