# Dice package
A GoLang package that allows complete management, autonomous or manual, of the use of dice for games of all kinds.

## Installation
First, make sure you have [GoLang 1.16](https://golang.org/doc/install) installed on your machine.<br>
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

func main ()() {
	//Create a standard n-sides dice
	sidesNumber := 6
	d1, err := dice.NewDice(sidesNumber)
	if err != nil {
		log.Panic(err)
	}
	//d1 is a 6-sided dice with faces valued ["1", "2", "3", "4", "5", "6"].

	//Roll the dice n-times and print the result
	rollTimes := 1
	result, err := d1.Throw(rollTimes)
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
	
	//Roll the dice n-times and print the result
	rollTimes := 3
	result, err := d2.Throw(rollTimes)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("The result of rolling the dice d2 is:", result)
}
```
**3. Roll more dice:** [GoPlayground]()
```Go
package main

import (
	"fmt"
	"log"
	"github.com/MoraGames/dice"
)

func main ()() {
	//Create a slice of dice
	d1, err := dice.NewDice(4)
	if err != nil {
		log.Panic(err)
	}
	d2, err := dice.NewDice(6)
	if err != nil {
		log.Panic(err)
	}
	d3, err := dice.NewDice(8)
	if err != nil {
		log.Panic(err)
	}
	d4, err := dice.NewCustomDice([]string{"-1", "0", "1"})
	if err != nil {
		log.Panic(err)
	}
	dices := dice.NewSliceOfDice(d1, d2, d3, d4)

	//Generate special rules
	frequency, err := dice.NewFrequency("")
	if err != nil {
		log.Panic(err)
	}
	options, err := dice.NewOptions(nil, frequency)
	if err != nil {
		log.Panic(err)
	}

	//Roll the dices n-times and print the result
	rollTimes := 1
	result, err := dice.Throw(dices, options, rollTimes)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("The result of the roll of the dices is:", result)
}
```
**4. roll more dice with special rules:** [GoPlayground]()
```Go
package main

import (
	"fmt"
	"log"
	"strconv"
	"github.com/MoraGames/dice"
)

func main ()() {
	//Create a slice of dice
	d1, err := dice.NewDice(4)
	if err != nil {
		log.Panic(err)
	}
	d2, err := dice.NewDice(6)
	if err != nil {
		log.Panic(err)
	}
	d3, err := dice.NewDice(8)
	if err != nil {
		log.Panic(err)
	}
	d4, err := dice.NewCustomDice([]string{"-1", "0", "1"})
	if err != nil {
		log.Panic(err)
	}
	dices := dice.NewSliceOfDice(d1, d2, d3, d4)

	//Generate special rules
	frequency, err := dice.NewFrequency("4") //Only after dice no.4 (dices[3])
	if err != nil {
		log.Panic(err)
	}
	operation := func (dicesResult ...string)(string, error){
		var sum int
		for i := 0; i < len(dicesResult)-1; i++ { //First, Second and Third dices
			result, err := strconv.Atoi(dicesResult[i])
			if err != nil {
				return "", err
			}
			sum += result
		}
		moltiplier, err := strconv.Atoi(dicesResult[len(dicesResult)-1]) //Fourth dice
		if err != nil {
			return "", err
		}
		return sum * moltiplier, nil
	}
	options, err := dice.NewOptions(operation, frequency)
	if err != nil {
		log.Panic(err)
	}

	//Roll the dices n-times and print the result
	rollTimes := 1
	result, err := dice.Throw(dices, options, rollTimes)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("The result of the roll of the dices is:", result)
}
```
## Functions
### Type: dice
Throw:
```Go
func (d *dice) Throw(rollTimes int) ([]string, error) {}
```