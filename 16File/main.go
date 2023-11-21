package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content:="This needs to go in a file";
	file,error:= os.Create("./myLogFile.txt")//os package provides the ways to create a file

	if error !=nil{
		panic(error)
	}
	length,err:= io.WriteString(file,content)
	checkNilError(err)

	fmt.Println("Length is :",length)
	file.Close()
	readFile("./myLogFile.txt")
}

func readFile (fileName string){

databyte,err:=	os.ReadFile(fileName)
checkNilError(err)
fmt.Println("Test data inside the file is: \n",string(databyte))
}

func checkNilError(err error){

	if err !=nil{
		panic(err)
	}
}