package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Welcome to our app")
	fmt.Println("Please rate our Pizza between 1 to g")
	reader :=bufio.NewReader(os.Stdin)

	input,_:= reader.ReadString('\n')
	fmt.Println("Thanks for rating",input)
	numRating,err:=strconv.ParseInt(strings.TrimSpace(input),32,64)

	if err !=nil{
		fmt.Println(err)

	}else{
		fmt.Println("Added 1 to your rating",numRating+1)
	}

}