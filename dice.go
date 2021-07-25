package dice

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func Throw(dices []dice, options options, rollTimes int) ([][]string, error) {
	var totalResults [][]string
	if rollTimes < 1 {
		return nil, fmt.Errorf("The number of dices rolls (%d) must be greater than or equal to 1.", rollTimes)
	}
	for t := 0; t < rollTimes; t++ {
		var results []string
		for i, dice := range dices {
			result, _ := dice.Throw(1)
			results = append(results, result...)
			if options.operation != nil && options.frequency.contains(strconv.Itoa(i+1)) == true {
				realResult, err := options.operation(results...)
				if err != nil {
					return nil, err
				}
				results = append(results[:len(results)-1], realResult)
			}
		}
		totalResults = append(totalResults, results)
	}
	return totalResults, nil
}

const minSides int = 2

type dice struct {
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
	return dice{sidesNumber, sidesValue}, nil
}

func NewCustomDice(sidesValue []string) (dice, error) {
	if len(sidesValue) < minSides {
		return dice{}, fmt.Errorf("The number of faces (%d) must be greater than or equal to %d.", len(sidesValue), minSides)
	}
	return dice{len(sidesValue), sidesValue}, nil
}

func NewSliceOfDice(dices ...dice) []dice {
	return dices
}

func (d dice) Throw(rollTimes int) ([]string, error) {
	if rollTimes < 1 {
		return nil, fmt.Errorf("The number of dice rolls (%d) must be greater than or equal to 1.", rollTimes)
	}
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	var results []string
	for t := 0; t < rollTimes; t++ {
		results = append(results, d.sidesValue[random.Intn(d.sidesNumber)])
	}
	return results, nil
}

// Ciao
