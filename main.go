// Initial Program

package main

import "fmt"

func main() {
	fmt.Println("Hello Gophers! 🥳")
	main2()
	stringLiteral()
}

// Format Printing

func main2() {
	const name, age = "Gowtham", 22
	fmt.Printf("%s is %d years old.\n", name, age)
}

func stringLiteral() {
	fmt.Println(`This 
	is a
	string 
	literal`)
}
