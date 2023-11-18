package main

import "fmt"

func main() {
	fmt.Println("Loops in golang")

	days := []string{"Sunday", "Tuseday", "Wednesday", "Friday"}
	fmt.Println(days)

	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	for i := range days {
		fmt.Println(days[i])
	}

	for index, day := range days {
		fmt.Printf("Index is %v and value is %v\n", index, day)
	}
lco:
	fmt.Println("Go to")
	rougueValue := 1
	for rougueValue < 10 {

		if rougueValue == 5 {
			rougueValue++
			continue
			
		}
		
		rougueValue++
		goto lco
	}


}
