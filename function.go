package main

import "fmt"

func add(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	result := a - b
	return result
}
func mul(a, b int) (result int) {
	result = a * b
	return 
}
func div(a, b int) (int) {
	return a / b 
}
func main(){
	fmt.Println("Hello !!")

	a := 30
	b := 10

	sum := add(a, b)
	fmt.Printf("%d", sum)

	difference := sub(a, b)
	fmt.Printf("\n%d", difference)

	product := mul(a, b)
	fmt.Printf("\n%d", product)

	quotient := div(a, b)
	fmt.Printf("\n%d", quotient)
}