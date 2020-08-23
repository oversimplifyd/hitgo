package main

import (
	"fmt"
	"learningGo/mylib"
)

func main(){
	mylib.PrintVariables()
	mylib.PrintTypes()
	mylib.PrintConstants()

	a, _ := mylib.FuncWithNamedParameters(3, 4.0)
	// second return value ignored using a blank identifier 
	fmt.Println(a)

	mylib.Conditionals()
	mylib.Loops()
	mylib.PrintArrays()
	mylib.PrintSlices()
}