package main

import (
	"fmt"
	
)

func main() {
	fmt.Println("Maps in golang")
	//create map
	var m map[string] int
	//ERROR
	//m["one"]=1
	fmt.Println(m)


	languages :=make(map[string]string)
	languages["one"]="Java"
	languages["two"]="JS"
	languages["three"]="Go"
	languages["four"]="python"
	
	fmt.Println(languages)
	//delete an item form map using key
	delete(languages,"one")
	fmt.Println(languages["two"])

	for key, value :=range languages{
		fmt.Printf("Key %v :Value %v\n",key,value)
	}
}