package dice

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type frequency []int

func (frq frequency) contains(element string) bool {
	e, _ := strconv.Atoi(element)
	for _, v := range frq {
		if e == v {
			return true
		}
	}
	return false
}

func NewFrequency(pattern string) (frequency, error) {
	var frq frequency

	elements := strings.Split(pattern, ",")
	for _, element := range elements {
		if matched, _ := regexp.MatchString(`^[0-9]+$`, element); matched == true {
			// 1 // 9069
			if frq.contains(element) == true {
				return nil, fmt.Errorf("The number %q must not be repeated within the pattern.", element)
			}
			v, _ := strconv.Atoi(element)
			frq = append(frq, v)
			continue
		}
		if matched, _ := regexp.MatchString(`([0-9]+)-([0-9]+)`, element); matched == true {
			// 1-3 // 1-31 // 13-1 // 13-31
			values := strings.Split(element, "-")
			vStart, _ := strconv.Atoi(values[0])
			vFinish, _ := strconv.Atoi(values[1])
			if vStart > vFinish {
				// 13-1
				return nil, fmt.Errorf("The set %q must have the start number (%d) lower than the finish number (%d).", element, vStart, vFinish)
			}
			for v := vStart; v <= vFinish; v++ {
				if frq.contains(strconv.Itoa(v)) == true {
					return nil, fmt.Errorf("The number %q must not be repeated within the pattern.", strconv.Itoa(v))
				}
				frq = append(frq, v)
			}
			continue
		}
	}
	return frq, nil
}
