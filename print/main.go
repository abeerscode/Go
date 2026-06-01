package main

import "fmt"

func main(){
	fmt.Print("Helllo world \n")                     // normal printing
	fmt.Println("Hello world")                       // normal printing + newline
	fmt.Printf("Hello World\n")                      // requires format verbs (%v, %s, %d, etc.)


	num := []int{0}
	num = append(num, 1, 2, 3, 4, 5)
	fmt.Printf("Numbers = %v\n", num)
	fmt.Printf("Position - 1: %d\n", num[0])         // %d → only for single integers

	name := []string{"Md.", "Abeer"}
	name = append(name, "Hasan")
	name = append(name, "Ratul")    
	fmt.Printf("Name = %v \n", name)                 // %v → print anything of a slice
	fmt.Printf("Position - 1: %s\n", name[0])        // %s → only for single strings


	cars := []string{"BMW", "MERCEDES"}
	cars = append(cars, "PORSCHE")
	fmt.Printf("Name = %q \n", cars)                  // %q  prints quoted string --> ["BMW" "MERCEDES" "PORSCHE"] <---output

	arr := [3]string{"DOG", "COW", "CAT"}
	fmt.Printf("Animals : %v\n", arr)

}
