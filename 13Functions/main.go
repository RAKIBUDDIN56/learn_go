package main

import "fmt"

func main() {
	fmt.Println("Go functions")
	result:=adder(2,4)
	fmt.Println(result)
	proRes,msg:=proAdder(2,3,4,)
	fmt.Println(proRes)
	fmt.Println(msg)
	animal_cow :=Animal("Cow")
	animal_cow.PrintAnimal()

}
func adder(val1 int, val2 int) int{
	return val1 + val2;

}
func proAdder(values ... int)(int,string){//... variadic syntaxt
	total :=0
for _,val := range values{
	total +=val
}
	return total,"Added"
}

type Animal string
func (An Animal) PrintAnimal(){
	fmt.Println(An)
}
