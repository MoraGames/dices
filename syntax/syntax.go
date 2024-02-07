package dices

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/MoraGames/dices"
	"github.com/MoraGames/dices/set"
	"github.com/MoraGames/dices/standardDice"
)

/*
DONE:
<N>d<S|\[MIN-MAX\]|\{SIDES\}>

TODO:
[nil|sum|sub|mul|tim[l|e|g]T] Nd[S|MIN-MAX] [[keep|drop][hig|low]R] [+|-|*|/[h|l|V]]
[nil|tim[l|e|g]T] Ndc{<sides>} [[keep|drop][hig|low]R]
*/
func NewSyntaxSet(syntax string) (*set.Set, error) {
	// Create, compile the regular expression and check if the syntax matches
	reg := regexp.MustCompile(`^([0-9]*)?d(([0-9]*)|(\[(\-|\+)?[0-9]+\,(\-|\+)?[0-9]+\])|(\{\N+(\,\N+)*\}))$`)
	if !reg.MatchString(syntax) {
		return nil, fmt.Errorf("the set of dices indicated is not supported by creation via syntax")
	}
	splittedSyntax := strings.Split(syntax, "d")

	// Check if the number of dices is a valid value or not
	var numberDices int
	if splittedSyntax[0] == "" {
		return nil, fmt.Errorf("the number of dices must be indicated")
	}
	var err error
	numberDices, err = strconv.Atoi(splittedSyntax[0])
	if err != nil {
		return nil, err
	}
	if minimumDices := set.GetMinimumDices(); numberDices < minimumDices {
		return nil, fmt.Errorf("the number of dices (%d) must be greater than or equal to %d", numberDices, minimumDices)
	}

	// Check if the number of sides is indicated as a number, a range, a list, or is not indicated
	var sides []dices.Side
	if splittedSyntax[1] == "" {
		return nil, fmt.Errorf("the number of sides must be indicated")
	} else if strings.Contains(splittedSyntax[1], "[") {
		// [MIN-MAX]
		var err error
		sides, err = getRangeSides(strings.Split(strings.Trim(splittedSyntax[1], "[]"), ","))
		if err != nil {
			return nil, err
		}
	} else if strings.Contains(splittedSyntax[1], "{") {
		// {SIDES}
		var err error
		sides, err = getCustomSides(strings.Split(strings.Trim(splittedSyntax[1], "{}"), ","))
		if err != nil {
			return nil, err
		}
	} else {
		// S (sides number)
		var err error
		sides, err = getNumberSides(splittedSyntax[1])
		if err != nil {
			return nil, err
		}
	}

	// Create the set of dices
	var dices []dices.Dice
	for i := 0; i < numberDices; i++ {
		d, err := standardDice.NewCustomDice(sides...)
		if err != nil {
			return nil, err
		}
		dices = append(dices, d)
	}
	return set.NewSet(dices...)
}

func getRangeSides(sidesSyntax []string) ([]dices.Side, error) {
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

	var sides []dices.Side
	for s := lowestSide; s <= highestSide; s++ {
		sides = append(sides, standardDice.NewStandardSide(s))
	}
	return sides, nil
}

func getNumberSides(sidesSyntax string) ([]dices.Side, error) {
	numberSides, err := strconv.Atoi(sidesSyntax)
	if err != nil {
		return nil, err
	}

	var sides []dices.Side
	for s := 1; s <= numberSides; s++ {
		sides = append(sides, standardDice.NewStandardSide(s))
	}
	return sides, nil
}

func getCustomSides(sidesSyntax []string) ([]dices.Side, error) {
	if minimumSides := dices.GetMinimumSides(); len(sidesSyntax) < minimumSides {
		return nil, fmt.Errorf("the number of sides must be greater than or equal to %d", minimumSides)
	}

	var sides []dices.Side
	for i := 0; i < len(sidesSyntax); i++ {
		sides = append(sides, standardDice.NewStandardSide(sidesSyntax[i]))
	}
	return sides, nil
}
