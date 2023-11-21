package main

import "fmt"

func main() {
	user:= User{"Rakib","ff@gmail.com",true,11}
	user.GetStatus()
	user.NewMail()

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)

}
func (u User) NewMail(){
	u.Email="rr@gmail.com"
	fmt.Println("Email of this user is :",u.Email)
}