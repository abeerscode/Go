package main

import "fmt"

func main(){
	var arr1[5] int
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	arr1[3] = 4
	arr1[4] = 5

	fmt.Println("Array 1 = ", arr1)

	var arr2 = [5] int {6,7,8,9,10}
	fmt.Println("Array 2 = ", arr2)

	arr3 := [4] string {"Md.","Abeer", "Hasan", "Ratul"}
	fmt.Printf("Array 3 = %q\n", arr3)

}
