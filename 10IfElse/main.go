package main

import "fmt"

func main() {
	fmt.Println("If else in golang")

	loginCount := 20
	var result string

	if loginCount < 10 {
		result = "Regular user"

	}else if loginCount> 10{
		result ="Watch out"
	}else{
		result="Exactly 10 login count"
	}
if num:=3; num<10{
	fmt.Println("Num is less than 10")
}else{
	fmt.Println("Num is not less than 10")
}

	fmt.Println(result)

}
