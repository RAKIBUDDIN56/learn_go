package main

import "fmt"

func main() {
	fmt.Println("Welome to pointers")

	var pointer *int

	fmt.Println(pointer)

	myNumber :=24
// & refereces to memory address 
	var memoryAddress =&myNumber
	fmt.Println(memoryAddress)
	fmt.Println(*memoryAddress)

	*memoryAddress=*memoryAddress*2
	fmt.Println(myNumber)

}