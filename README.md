# Dice package
A GoLang package that allows complete management, autonomous or manual, of the use of dice for games of all kinds.

## Special Thanks:
	- @Vano2903 in the role of `Alpha-Tester`

## Installation
First, make sure you have [GoLang](https://golang.org/doc/install) installed on your machine.<br>
Proceed by downloading the package with the `go get https://github.com/MoraGames/dice` command.<br>

## Example
**1. Roll a standard 6-sides dice:**
```Go
package main

import (
	"fmt"
	"github.com/MoraGames/dice"
)

func main ()() {
	//Create a new dice with 6 faces and their respective values
	dice.NewDice(6, []string{"1", "2", "3", "4", "5", "6"})

	//Roll the dice one time and print the result
	fmt.Println("The result of the launch is:", dice.Throw(1))
}
```
## Functions
### Type: dice
Throw:
```Go
func (d *dice) Throw(rollTimes int) ([]string, error) {}
```