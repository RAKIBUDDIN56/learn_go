package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	welcome :="Welcme to user input"
	fmt.Println(welcome)

	reader :=bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for out Pizza")

	//comma ok || err ok syntaxt
	imput,_ :=reader.ReadString('\n')
	fmt.Println("Thanks for rating",imput)
	fmt.Printf("Type of this is %T",imput)

}