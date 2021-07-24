package main

import (
	"fmt"
	//"github.com/MoraGames/dice"
)

func main() {
	//Create a new dice with 6 faces and their respective values
	dice.NewDice(6, []string{"1", "2", "3", "4", "5", "6"})

	//Roll the dice one time and print the result
	fmt.Println("The result of the launch is:", dice.Throw(1))
}
