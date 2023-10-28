package main

import "fmt"

func main() {
	fmt.Println("Welcome to Array in Go")

	var fruitList [5] string
	//fmt.Println(fruitList)
	fruitList[0]="Apple"
	fruitList[1]="Orange"
	fruitList[2]="Banana"
	fruitList[4]="Peach"

	fmt.Println(fruitList)

	fmt.Println("Lenght of the List is: ",len(fruitList))

	var vegList =[3]string{"potato","beans","mushroom"}
	fmt.Println("Veg List ",vegList)
	
}