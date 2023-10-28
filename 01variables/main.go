package main

import "fmt"
var jwttoken int =3000
//first letter capital means public variable
const LoginToken string ="ffff"

func main(){
	var userName string ="Rakib"
	fmt.Println(userName)

	fmt.Printf("Varibale is %T",userName)
var isTrue bool =false;
fmt.Printf("%v",isTrue)

var smallVal  uint8=255
fmt.Printf("\n%v\n",smallVal)
var mediumVal  uint16=2556
fmt.Printf("\n%v\n",mediumVal)
var floatVal  float64=2556.1234342199999999999
fmt.Printf("%v",floatVal)
fmt.Printf("\n%T\n",floatVal)
//default vlaue and some aliases
var anothervar int
fmt.Println(anothervar)
fmt.Printf("Varibale is %T\n",anothervar)
//implicit type
var website="my website"
fmt.Println(website)
fmt.Printf("Varibale is %T\n",website)
//no var tyle
numberUser :=3000
fmt.Println(numberUser)
fmt.Printf("Varibale is %T\n",numberUser)
fmt.Println(LoginToken)
fmt.Printf("Varibale is %T\n",LoginToken)

}