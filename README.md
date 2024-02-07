[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]

[![GoDoc][godoc-shield]][godoc-url]
[![Go Report Card][gocard-shield]][gocard-url]

# Dice package
A GoLang package that allows complete management, autonomous or manual, of the use of dice for games of all kinds.

## Installation
First, make sure you have [GoLang](https://golang.org/doc/install) installed on your machine.<br>
Proceed by downloading the package with the `go get -u github.com/MoraGames/dices` command.<br>

## Examples
**1. Roll a standard 6-sides dice:** [GoPlayground](https://go.dev/play/p/IXjnsYI6lZX)
```Go
package main

import (
	"fmt"
	"log"

	"github.com/MoraGames/dices/standardDice"
)

func main() {
	//Create a standard n-sides dice
	sidesNumber := 6
	d1, err := standardDice.NewDice(sidesNumber)
	if err != nil {
		log.Panic(err)
	}
	//d1 is a 6-sided dice with faces valued [1, 2, 3, 4, 5, 6]

	//Roll the dice and print the result
	result := d1.Throw()
	fmt.Println("The result of rolling the dice d1 is:", result)
}
```
**2. Roll a range dice:** [GoPlayground](https://go.dev/play/p/h35WkwlQ0ic)
```Go
package main

import (
	"fmt"
	"log"

	"github.com/MoraGames/dices"
	"github.com/MoraGames/dices/standardDice"
)

func main() {
	//Create a special ranged dice
	lowestSide, highestSide := -1, 1
	r1, err := dices.NewRange(lowestSide, highestSide)
	if err != nil {
		log.Panic(err)
	}
	d2, err := standardDice.NewRangeDice(r1)
	if err != nil {
		log.Panic(err)
	}
	//d2 is a 3-sided dice with faces valued [-1, 0, 1]

	//Roll the dice and print the result
	result := d2.Throw()
	fmt.Println("The result of rolling the dice d2 is:", result)
}
```
**3. Roll a custom dice:** [GoPlayground](https://go.dev/play/p/GyTUWdyG0_j)
```Go
package main

import (
	"fmt"
	"log"

	"github.com/MoraGames/dices"
	"github.com/MoraGames/dices/standardDice"
)

func main() {
	//Create a custom n-sides dice and their respective values
	sidesValue := []dices.Side{
		standardDice.NewStandardSide("Apple"),
		standardDice.NewStandardSide("Banana"),
		standardDice.NewStandardSide("Cherry"),
		standardDice.NewStandardSide("Dates"),
		standardDice.NewStandardSide("Elderberry"),
	}
	d3, err := standardDice.NewCustomDice(sidesValue...)
	if err != nil {
		log.Panic(err)
	}
	//d3 is a 5-sided dice with faces valued ["Apple", "Banana", "Cherry", "Dates", "Elderberry"]

	//Roll the dice and print the result
	result := d3.Throw()
	fmt.Println("The result of rolling the dice d3 is:", result)
}
```
**4. Roll a set of dices:** [GoPlayground](https://go.dev/play/p/Pe6mMGzOJVU)
```Go
package main

import (
	"fmt"
	"log"

	"github.com/MoraGames/dices"
	"github.com/MoraGames/dices/set"
	"github.com/MoraGames/dices/standardDice"
)

func main() {
	//Create a standard n-sides dice
	sidesNumber := 6
	d4a, err := standardDice.NewDice(sidesNumber)
	if err != nil {
		log.Panic(err)
	}
	//d4a is a 6-sided dice with faces valued [1, 2, 3, 4, 5, 6]

	//Create a special ranged dice
	lowestSide, highestSide := -1, 1
	r2, err := dices.NewRange(lowestSide, highestSide)
	if err != nil {
		log.Panic(err)
	}
	d4b, err := standardDice.NewRangeDice(r2)
	if err != nil {
		log.Panic(err)
	}
	//d4b is a 3-sided dice with faces valued [-1, 0, 1]

	//Create a custom n-sides dice and their respective values
	sidesValue := []dices.Side{
		standardDice.NewStandardSide("Apple"),
		standardDice.NewStandardSide("Banana"),
		standardDice.NewStandardSide("Cherry"),
		standardDice.NewStandardSide("Dates"),
		standardDice.NewStandardSide("Elderberry"),
	}
	d4c, err := standardDice.NewCustomDice(sidesValue...)
	if err != nil {
		log.Panic(err)
	}
	//d4c is a 5-sided dice with faces valued ["Apple", "Banana", "Cherry", "Dates", "Elderberry"]

	//Create a custom set of n-dices
	s1, err := dices.NewSet(d4a, d4b, d4c)
	if err != nil {
		log.Panic(err)
	}

	//Roll all the dices in the set
	result := s1.Throw()
	fmt.Println("The result of rolling the set s1 is:", result)
}
```

[contributors-shield]: https://img.shields.io/github/contributors/MoraGames/dices.svg?style=for-the-badge
[contributors-url]: https://github.com/MoraGames/dices/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/MoraGames/dices.svg?style=for-the-badge
[forks-url]: https://github.com/MoraGames/dices/network/members
[stars-shield]: https://img.shields.io/github/stars/MoraGames/dices.svg?style=for-the-badge
[stars-url]: https://github.com/MoraGames/dices/stargazers
[issues-shield]: https://img.shields.io/github/issues/MoraGames/dices.svg?style=for-the-badge
[issues-url]: https://github.com/MoraGames/dices/issues
[license-shield]: https://img.shields.io/github/license/MoraGames/dices.svg?style=for-the-badge
[license-url]: https://github.com/MoraGames/dices/blob/master/LICENSE.md
[godoc-shield]: https://godoc.org/github.com/MoraGames/dices?status.svg
[godoc-url]: https://pkg.go.dev/github.com/MoraGames/dices
[gocard-shield]: https://goreportcard.com/badge/github.com/MoraGames/dices
[gocard-url]: https://goreportcard.com/report/github.com/MoraGames/dices
