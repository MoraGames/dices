package dice

import (
	"fmt"
	"regexp"
)

/*
[nil|sum|sub|mul|tim[l|e|g]T] Nd[S|MIN-MAX] [[keep|drop][hig|low]R] [+|-|*|/[h|l|V]]
[nil|tim[l|e|g]T] Ndc{<sides>} [[keep|drop][hig|low]R]
*/

func NewSyntaxSet(syntax string) (*Set, error) {
	reg := regexp.MustCompile("^(nil |sum |sub |mul |enu [leg]\\d+ )?\\d+d(\\d+|\\d+\\-\\d+)( (k|d)?(h|l)?\\d+)?( [\\+\\-\\*\\/](h|l|\\d+))?$")
	if reg.MatchString(syntax) == false {
		return nil, fmt.Errorf("syntax is not valid")
	}

	/* TODO: complete the syntax check

	var enuVal, nDices, nSides, subVal int
	macroGroups := strings.Split(syntax, " ")

	if len(macroGroups) == 1 { //composition

	} else if len(macroGroups) == 2 { //composition + (exclusion | operation)
		exclusionReg := regexp.MustCompile("^(k|d)?(h|l)?[1-9]\\d*$")
		if exclusionReg.MatchString(macroGroups[1]) == true { //composition + exclusion

		} else { //composition + operation
			switch macroGroups[1][:3] {
			case "nil", "avg":
			case "sum", "sub", "mul", "div":
			case "enu":
			default: //both exclusion nor operation invalid
				return nil, fmt.Errorf("")
			}
		}

		if macroGroups[1][0] != 'k' && macroGroups[1][0] != 'd' {

		}
	} else if len(macroGroups) == 3 { //composition + exclusion + operation

	}

	if strings.HasPrefix(macroGroups[0], "enu") == true {

	}
	if len(macroGroups) == 4 { //Numbered Dices

	} else if len(macroGroups) == 3 { //Custom Dices

	} else {
		return nil, fmt.Errorf("syntax is not knowed")
	}

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
	*/

	var dices []*Dice
	for i := 0; i < nDices; i++ {
		d, _ := NewDice(nSides)
		dices = append(dices, d)
	}
	return &Set{dices}, nil
}

/*
^NdS|[](([1-9]\d*)|(\[[-+]?\d+\:[-+]?\d+\]))( (k|d)?(h|l)?[1-9]\d*)?(( (sum|sub|mul|div)?(hg|lw|[1-9]\d*))| nil| avg| (enu(lt|le|eq|ge|gt)?(lw|av|hg|\d+)))?$

N


0d6
1d0
1d6
2d8
3d[-1:1]
3d[-3:+5]
4d6 kh0
4d7 kh3
4d8 dl2
5d6 sumhg
5d7 sub8
5d8 mul12
11d[-4:4] dh4 nil
10d[-3:+3] kl3 sum17
12d[-4:4] dh4 enulehg
13d[-4:4] dh4 enugt8

15d{"a","b","c"} dh4
*/
