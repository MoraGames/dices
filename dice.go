package dices

const minimumSides = 2

func GetMinimumSides() int {
	return minimumSides
}

type (
	Dice interface {
		Throw() Side
		ThrowN(n int) []Side
	}
	Side interface {
		Value() any
	}
)
