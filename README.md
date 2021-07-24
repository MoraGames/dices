# Dice package
A GoLang package that allows complete management, autonomous or manual, of the use of dice for games of all kinds.

## Installation
First, make sure you have [GoLang](https://golang.org/doc/install) installed on your machine.<br>
Proceed by downloading the package with the `go get https://github.com/MoraGames/dice` command.<br>

## Examples
**1. Roll a standard 6-sides dice:**
```Go
package main

import (
	"fmt"
	"github.com/MoraGames/dice"
)

func main ()() {
	//Create a standard n-sides dice
	sidesNumber := 6
	d1, err := dice.NewDice(sidesNumber)
	if err != nil {
		log.Panic(err)
	}
	//d1 is a 6-sided dice with faces valued ["1", "2", "3", "4", "5", "6"].

	//Roll the dice n-times and print the result
	rollTimes = 1
	fmt.Println("The result of rolling the dice d1 is:", d1.Throw(rollTimes))
}
```
**2. Roll a custom dice:**
```Go
package main

import (
	"fmt"
	"github.com/MoraGames/dice"
)

func main ()() {
	//Create a custom n-sides dice and their respective values
	sidesValue := []string{"Apple", "Banana", "Cherry", "Dates", "Elderberry"}
	d2, err := dice.NewCustomDice(sidesValue)
	if err != nil {
		log.Panic(err)
	}
	//d2 is a 5-sided dice with faces valued ["Apple", "Banana", "Cherry", "Dates", "Elderberry"]
	
	//Roll the dice n-times and print the result
	rollTimes := 3
	fmt.Println("The result of rolling the dice d2 is:", d2.Throw(rollTimes))
}
```
## Functions
### Type: dice
Throw:
```Go
func (d *dice) Throw(rollTimes int) ([]string, error) {}
```