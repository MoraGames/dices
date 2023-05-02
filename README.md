# Dice package
A GoLang package that allows complete management, autonomous or manual, of the use of dice for games of all kinds.

## Installation
First, make sure you have [GoLang](https://golang.org/doc/install) installed on your machine.<br>
Proceed by downloading the package with the `go get -u github.com/MoraGames/dice` command.<br>

## Examples
**1. Roll a standard 6-sides dice:** [GoPlayground](https://play.golang.org/p/JZ5slbKhCtI)
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
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("The result of rolling the dice d1 is:", result)
}
```
**2. Roll a custom dice:** [GoPlayground](https://play.golang.org/p/LYxP8DfyVdu)
```Go
package main

import (
	"fmt"
	"log"
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
	
	//Roll the dice and print the result
	result := d2.Throw()
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("The result of rolling the dice d2 is:", result)
}
```