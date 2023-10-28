package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices ")

	var fruitList =[] string{"Apple","Banana","Orange"}
	fmt.Printf("Tyoe of fruitList is %T\n",fruitList)
	var vegList =[3]string{"potato","beans","mushroom"}
	fmt.Printf("Veg List %T",vegList)
	fruitList=append(fruitList,"Pineapple" )
	fmt.Println(fruitList)
a:= [] string {"A"}
a= append(a,fruitList...)
fmt.Println(len(a))
//create slice form array
// slice := array[start_index : end_index]
arr1 :=[6] int {10, 11, 12, 13, 14,15}
mySlice := arr1[2:4]


fmt.Println(mySlice)

//creat slice using make() function
newSlice :=make([] int,4)
newSlice[0]=2;
newSlice[1]=45;
newSlice[2]=3;
newSlice[3]=10;
newSlice=append(newSlice,10)
fmt.Println(cap(newSlice))
fruitList=fruitList[1:2] 
fmt.Println(fruitList)
sort.Ints(newSlice)
fmt.Println(newSlice)
fmt.Println(sort.IntsAreSorted(newSlice))



	


}
