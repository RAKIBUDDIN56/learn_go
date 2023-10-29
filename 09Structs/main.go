package main

import (
	
	"fmt"
)

func main() {
	fmt.Println("Structs in go")
	// no inheritance in golang; No super or parent

	user:=User{"Rakib","rrr@gmail.com",true,23}
	fmt.Println(user)
	// %+v reprensents the o value with field names ( for struct)
	fmt.Printf("User : %+v",user)
	fmt.Printf("Name is %v and email is %v",user.Name,user.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
