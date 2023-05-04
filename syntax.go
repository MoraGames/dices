package dice

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

/*
[nil|sum|sub|mul|tim[l|e|g]T] Nd[S|MIN-MAX] [[keep|drop][hig|low]R] [+|-|*|/[h|l|V]]
[nil|tim[l|e|g]T] Ndc{<sides>} [[keep|drop][hig|low]R]
*/

func NewSyntaxSet(syntax string) (*Set, error) {
	reg := regexp.MustCompile("^((0*[1-9]+)|[1-9]*)d(0*)([2-9]|[1-9][0-9]+)$")
	if reg.MatchString(syntax) == false {
		return nil, fmt.Errorf("")
	}

	var nDices, nSides int
	splittedSyntax := strings.Split(syntax, "d")
	if splittedSyntax[0] == "" {
		nDices = minDices
	} else {
		nDices, _ = strconv.Atoi(splittedSyntax[0])
	}
	if splittedSyntax[1] == "" {
		nSides = minSides
	} else {
		nSides, _ = strconv.Atoi(splittedSyntax[1])
	}

	var dices []*Dice
	for i := 0; i < nDices; i++ {
		d, _ := NewDice(nSides)
		dices = append(dices, d)
	}
	return &Set{dices}, nil
}
