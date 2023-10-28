package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time tudy")

	currenttime :=time.Now()

	fmt.Println(currenttime)
//data format
	fmt.Println(currenttime.Format("02-01-2006"))
	//day format
	fmt.Println(currenttime.Format("Monday"))
	//time format
	fmt.Println(currenttime.Format("15:04:05"))

	createDate:=time.Date(2020,time.April,2,4,1,2,3,time.Now().Local().Location())
	fmt.Println(createDate)
	fmt.Println(createDate.Format("02-01-2006 Monday"))

	

}