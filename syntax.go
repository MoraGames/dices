package dices

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

/*
DONE:
[N]d[<S>|<\[MIN-MAX\]>|<\{SIDES\}>]

TODO:
[nil|sum|sub|mul|tim[l|e|g]T] Nd[S|MIN-MAX] [[keep|drop][hig|low]R] [+|-|*|/[h|l|V]]
[nil|tim[l|e|g]T] Ndc{<sides>} [[keep|drop][hig|low]R]
*/

func NewSyntaxSet(syntax string) (*Set, error) {
	reg := regexp.MustCompile(`^([0-9]*)?d(([0-9]*)|(\[(\-|\+)?[0-9]+\,(\-|\+)?[0-9]+\])|(\{\N+(\,\N+)*\}))$`)
	if !reg.MatchString(syntax) {
		return nil, fmt.Errorf("the set of dices indicated is not supported by creation via syntax")
	}

	var numberDices int
	var sides []Side
	splittedSyntax := strings.Split(syntax, "d")

	// Check if the number of dices is a valid value or not indicated (in this last case use default value)
	if splittedSyntax[0] == "" {
		numberDices = defaultDices
	} else {
		var err error
		numberDices, err = strconv.Atoi(splittedSyntax[0])
		if err != nil {
			return nil, err
		}
		if numberDices < minimumDices {
			return nil, fmt.Errorf("the number of dices (%d) must be greater than or equal to %d", numberDices, minimumDices)
		}
	}

	// Check if the number of sides is a valid value or not indicated (in this last case use default value)
	if splittedSyntax[1] == "" {
		for s := 1; s <= defaultSides; s++ {
			sides = append(sides, s)
		}
	} else {
		if strings.Contains(splittedSyntax[1], "[") {
			// [MIN-MAX]
			sidesSyntax := strings.Split(strings.Trim(splittedSyntax[1], "[]"), ",")
			if len(sidesSyntax) != 2 {
				return nil, fmt.Errorf("the number of sides indicated is not supported by creation via syntax")
			}
			lowestSide, err := strconv.Atoi(sidesSyntax[0])
			if err != nil {
				return nil, err
			}
			highestSide, err := strconv.Atoi(sidesSyntax[1])
			if err != nil {
				return nil, err
			}
			if lowestSide >= highestSide {
				return nil, fmt.Errorf("the lowestSide (%d) must be less than highestSide (%d)", lowestSide, highestSide)
			}

			for s := lowestSide; s <= highestSide; s++ {
				sides = append(sides, s)
			}
		} else if strings.ContainsAny(splittedSyntax[1], "{") {
			// {SIDES}
			sidesSyntax := strings.Split(strings.Trim(splittedSyntax[1], "{}"), ",")
			for _, s := range sidesSyntax {
				sides = append(sides, s)
			}
		} else {
			// S (sides number)
			numberSides, err := strconv.Atoi(splittedSyntax[1])
			if err != nil {
				return nil, err
			}
			for s := 1; s <= numberSides; s++ {
				sides = append(sides, s)
			}
		}
	}

	var dices []*Dice
	for i := 0; i < numberDices; i++ {
		d, err := NewCustomDice(sides)
		if err != nil {
			return nil, err
		}
		dices = append(dices, d)
	}
	return &Set{dices}, nil
}
