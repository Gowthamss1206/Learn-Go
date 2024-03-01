package main

import (
	"fmt"
	"math"
	"math/rand"
)

var globalScopeVar string = "Testing the global scope variable"

func main() {
	x := 28
	name := "Gowtham's"

	fmt.Printf("%d is %s one of the fav number\n", x, name)

	fmt.Println("Accessing global scoped variable inside the function:", globalScopeVar)

	randomNumber()

	squareRooting()

	addedResult := add(28, 6)
	fmt.Printf("Result of addition is %d\n", addedResult)

	resultStr := nakedReturn("Good Learning Today")
	fmt.Println(resultStr)
}

func randomNumber() {
	number := rand.Intn(12)
	fmt.Printf("%d will also be Gowtham's fav number\n", number)
}

func squareRooting() {
	squareNum := 28
	fmt.Printf("Square Rooting the fav number: %f\n", math.Sqrt(float64(squareNum)))
}

func add(x, y int) int {
	result := x + y
	return result
}

func nakedReturn(inputStr string) (resultStr string) {
	resultStr = inputStr
	return
}
