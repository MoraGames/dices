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
Proceed by downloading the package with the `go get -u github.com/MoraGames/dice` command.<br>

## Examples
**1. Roll a standard 6-sides dice:** [GoPlayground](https://go.dev/play/p/uh8fWjBEAn_0)
```Go
package main

import (
	"fmt"
	"log"
	
	"github.com/MoraGames/dice"
)

func main() {
	//Create a standard n-sides dice
	sidesNumber := 6
	d1, err := dice.NewDice(sidesNumber)
	if err != nil {
		log.Panic(err)
	}
	//d1 is a 6-sided dice with faces valued ["1", "2", "3", "4", "5", "6"].

	//Roll the dice and print the result
	result  := d1.Throw()
	fmt.Println("The result of rolling the dice d1 is:", result)
}
```
**2. Roll a custom dice:** [GoPlayground](https://go.dev/play/p/CoWKl382p2O)
```Go
package main

import (
	"fmt"
	"log"
	
	"github.com/MoraGames/dice"
)

func main() {
	//Create a custom n-sides dice and their respective values
	sidesValue := []string{"Apple", "Banana", "Cherry", "Dates", "Elderberry"}
	d2, err := dice.NewCustomDice(sidesValue...)
	if err != nil {
		log.Panic(err)
	}
	//d2 is a 5-sided dice with faces valued ["Apple", "Banana", "Cherry", "Dates", "Elderberry"]
	
	//Roll the dice and print the result
	result := d2.Throw()
	fmt.Println("The result of rolling the dice d2 is:", result)
}
```
**3. Roll a set of dices:** [GoPlayground](https://go.dev/play/p/0KeDyKLMLXJ)
```Go
package main

import (
	"fmt"
	"log"

	"github.com/MoraGames/dice"
)

func main() {
	//Create a standard n-sides dice
	sidesNumber := 6
	d3a, err := dice.NewDice(sidesNumber)
	if err != nil {
		log.Panic(err)
	}
	//d3a is a 6-sided dice with faces valued ["1", "2", "3", "4", "5", "6"].

	//Create a custom n-sides dice and their respective values
	sidesValue1 := []string{"Apple", "Banana", "Cherry", "Dates", "Elderberry"}
	d3b, err := dice.NewCustomDice(sidesValue1...)
	if err != nil {
		log.Panic(err)
	}
	//d3b is a 5-sided dice with faces valued ["Apple", "Banana", "Cherry", "Dates", "Elderberry"]

	//Create a custom n-sides dice and their respective values
	sidesValue2 := []string{"-1", "0", "1"}
	d3c, err := dice.NewCustomDice(sidesValue2...)
	if err != nil {
		log.Panic(err)
	}
	//d3c is a 3-sided dice with faces valued ["-1", "0", "1"]

	//Create a custom set of n-dices
	s1, err := dice.NewSet(d3a, d3b, d3c)
	if err != nil {
		log.Panic(err)
	}

	//Roll all the dices in the set
	result := s1.Throw()
	fmt.Println("The result of rolling the set s1 is:", result)
}
```

[contributors-shield]: https://img.shields.io/github/contributors/MoraGames/dice.svg?style=for-the-badge
[contributors-url]: https://github.com/MoraGames/dice/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/MoraGames/dice.svg?style=for-the-badge
[forks-url]: https://github.com/MoraGames/dice/network/members
[stars-shield]: https://img.shields.io/github/stars/MoraGames/dice.svg?style=for-the-badge
[stars-url]: https://github.com/MoraGames/dice/stargazers
[issues-shield]: https://img.shields.io/github/issues/MoraGames/dice.svg?style=for-the-badge
[issues-url]: https://github.com/MoraGames/dice/issues
[license-shield]: https://img.shields.io/github/license/MoraGames/dice.svg?style=for-the-badge
[license-url]: https://github.com/MoraGames/dice/blob/master/LICENSE.md
[godoc-shield]: https://godoc.org/github.com/MoraGames/dice?status.svg
[godoc-url]: https://pkg.go.dev/github.com/MoraGames/dice
[gocard-shield]: https://goreportcard.com/badge/github.com/MoraGames/dice
[gocard-url]: https://goreportcard.com/report/github.com/MoraGames/dice
