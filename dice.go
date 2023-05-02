package dice

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const minSides int = 2

type dice struct {
	rand *rand.Rand

	sidesNumber int
	sidesValue  []string
}

func NewDice(sidesNumber int) (dice, error) {
	if sidesNumber < minSides {
		return dice{}, fmt.Errorf("The number of faces (%d) must be greater than or equal to %d.", sidesNumber, minSides)
	}
	var sidesValue []string
	for s := 1; s <= sidesNumber; s++ {
		sidesValue = append(sidesValue, strconv.Itoa(s))
	}
	return dice{rand.New(rand.NewSource(time.Now().UnixNano())), sidesNumber, sidesValue}, nil
}

func NewCustomDice(sidesValue []string) (dice, error) {
	if len(sidesValue) < minSides {
		return dice{}, fmt.Errorf("The number of faces (%d) must be greater than or equal to %d.", len(sidesValue), minSides)
	}
	return dice{rand.New(rand.NewSource(time.Now().UnixNano())), len(sidesValue), sidesValue}, nil
}

func (d dice) Throw() string {
	return d.sidesValue[d.rand.Intn(d.sidesNumber)]
}
