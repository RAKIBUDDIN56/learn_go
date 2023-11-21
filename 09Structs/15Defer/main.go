package main

import "fmt"

func main() {
	//defer means put back/ delay . defer delays  the execution
//last in first come
defer	fmt.Println("Hello Defer")
fmt.Println("First comes")
myDefer()
}

func myDefer(){
	for i:=0; i<5;i++{
	defer	fmt.Println(i)
	}
}