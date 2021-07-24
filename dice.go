package dice

import (
	"fmt"
	"math/rand"
	"time"
)

type dice struct {
	sidesNumber int
	sidesValue  []string
}

func NewDice(sidesNumber int, sidesValue []string) (*dice, error) {
	if len(sidesValue) > sidesNumber {
		return nil, fmt.Errorf("The length of sidesValue (%d) must be less or equal than the sidesNumber (%d).", len(sidesValue), sidesNumber)
	}
	return &dice{sidesNumber, sidesValue}, nil
}

func (d *dice) Throw(rollTimes int) ([]string, error) {
	if rollTimes < 1 {
		return nil, fmt.Errorf("The number of dice rolls (%d) must be greater than or equal to 1.", rollTimes)
	}
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	var result []string
	for t := 0; t < rollTimes; t++ {
		result = append(result, d.sidesValue[random.Intn(d.sidesNumber)])
	}
	return result, nil
}
