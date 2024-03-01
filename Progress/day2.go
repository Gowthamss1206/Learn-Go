package main

import "fmt"

func main() {
	fmt.Println("Started Day 2")
	printHexadecimal()
	printBinary()
}

func printHexadecimal() {
	num := 12
	fmt.Printf("%d in hexadecimal is %x\n", num, num)
}

func printBinary() {
	num2 := 28
	fmt.Printf("%d in binary is %b\n", num2, num2)
}
