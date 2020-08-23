package mylib

import "fmt"

func init(){
	// Can not be called explicityly 
	// Called before any other thing 
	// Used in doing some pre-work 
	var _ = 40  // _ blank identifier, can also be used as silencer. Go's compiler typically complains of unused vars except this. 

	fmt.Println("Setting up init..")
}

func PrintVariables(){

	var age int = 10 

	// Variables can be assigned to any value of its type 
	age = 20  

	// Types can be inferrred 
	var secondAge = 13 

	// Multiple variable declaration the same type 
	var thirdAge, foruthAge = 10, 15

	// Multiple variable declartion different types 
	var (
		fifthAge = 10
		stringValue = "Yope.."
		sixthAge = 100
	)

	//shorthand declaration without the 'var' keyword 
	firstAgeShortHand := 20 
	secondAgeShortHand, thirdAgeShortHand, firstAgeShortHand, stringValueShorthand := 10, 20, 30, "Yope.."

	fmt.Println(age)
	fmt.Println(secondAge)
	fmt.Println(thirdAge)
	fmt.Println(foruthAge)
	fmt.Println(fifthAge)
	fmt.Println(sixthAge)
	fmt.Println(stringValue)

	fmt.Println(secondAgeShortHand)
	fmt.Println(firstAgeShortHand)
	fmt.Println(stringValueShorthand)
	fmt.Println(thirdAgeShortHand)

}

func PrintTypes(){

	a,b := true,false

	var (
		c uint16 = 70
		d rune = 30
	)

	e := 14
	f := 17.4

	// g := e + f  -> Illegal, incompatible types 
	g := e + int(f)

	//NB: There are no implicit casting in Go. 
	// Coercion/Casting has to be done explicitly
	// string is a collection of chars

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(a)
	fmt.Println(f)
	fmt.Println(g)
}

func PrintConstants(){
	// Any string enclosed in double quotes is a constant 
	// Constants can not be re-assigned 
	// Types can be aliased 
	// const can not be assigned to the result of a function call, because func call happens at runtime while const is determined at compile time 
	// Real types and their alias cannot be used interchangeably

	var a = "Klakd" // String constant, by default 

	const b = "Jammy" //Not a string constant, just a constant, untyped
	const c string = "Lomilo"  // Typed string constant

	var d = 5.9/8 // type is inferred in the numeric expression

	//const b = math.Sqrt(4)//not allowed

	type customStringType string // type aliasing, customStringType is now a type

	fmt.Println(a, d, b, c)
}

func FuncWithMultipleParams(a, b int)(int){
	return a + b
}

func FuncWithMultipleReturnValues(a, b int)(int, int){
	return a + b, b + 3
}

func FuncWithNamedParameters(a int, b float32)(first float32, second int){
	first = b + 0.4
	second = a + int(b)

	return
}

func Conditionals(){

	num := 10
	var _ = 50

	// if.else-if..else
	if num > 2 {
		num = 15
	} else if num < 3 {
		num = 12
	} else {
		num = 60
	}

	// if..else-if...else with asssignments 
	if num2 := 43; num > 5 { // variables defined in block are only visible in the block
		num = num2
	}

	switch num {
	case 3:
		fmt.Println("Yah..")
		break
	case 10:
		fmt.Println("Case..10")
		fallthrough // examines the next case
	default:
		fmt.Println("Default case.")
	}
	// Go also supports expressionless switch, in this switch, all truthy values are evaulated

	fmt.Println(num)
}

func Loops(){
	i := 1

	//variant 1
	for ; i <= 10; i++ {
		fmt.Println(i)
	}

	// variant 2
	for i <= 10 {
		fmt.Println(i)
		i+=3 // inrementer
	}

	// variant 3
	for i := 0; i <= 3; i+=1 {
		fmt.Println(i)
	}
}

func PrintArrays(){

	// The other values are prefilled with the zero values of int
	// Len can not be modified after declared 
	// Array types cannot be mixed .: firstArray can not be mixed with firstArrayb

	var firstArray = [9]int{1, 2, 3}   // partial initialization
	var firstArrayb = [4]int{3}
	secondArray := [6]string{"first", "second", "third"}

	// for..range, gets but the key and value of arrays 
	for i, v := range firstArray {
		fmt.Println(i, v)
	}

	//MultiDimensional arrays 
	mutliDimensionalArray := [3][2]string {
		{"james", "Ada"},
		{"Love", "Ayam"},
		{"Peter", "Jonah"},  //this last comma is mandatory -> mutliDimensionalArray[0][1]
	}

	fmt.Println(len(secondArray), len(mutliDimensionalArray))
	fmt.Println(len(firstArrayb))
}

func PrintSlices(){

	/** 
	Notes: 
	- slices, affects the underlying array
	- A slice can be appended to. If while appending the slice goes out of cap,  the slice expands. 
	- This is done internally by creating new array with 2x the original aray  for the slice 
	- The make functions creates an array and treutrns a slice reference of it 
	- Slices are passed by value 
	- When passed to a function, changes to slices are reflected even out o the function unlke an array. 
	- To avoid issues with memory optimization, if slices are being obtained from a large array, remember the underlying array they were being obtained from ...
		are still in memory and can not be garbage collected because the slice, even though it's a tiny cut of the arrray still points to the array, 
		to avoid this problem, use the copy() to copy a slice into another, so that underlying array of the prev can be GCd.
	**/
	initialArray := [...]int{3, 4, 5, 6, 7, 8, 9, 10, 11, 12} // the ... operator is equivalent to number of lieterals in the array, it doesn't mean the array does not have a speiified capacity
	
	sliceA := initialArray[:]
	sliceB := sliceA[1:4]  // [0], [1], [2], [3] 
	sliceB[0] = 100

	
	sliceA = append(sliceA, 10)

	fmt.Println(cap(sliceA))   // Capacity is the cap of the underlying array 
	fmt.Println(len(sliceA))  // len is the current len of the slice 
	fmt.Println(sliceB[1])

	sliceC := make([]int, 3, 4)  // initilizze the slice with 0s
	sliceC[0] = 30

	sliceD := []int{8, 7, 6}  // This creates a slice, although it uses an implicit array internally. This is different from an array, notice the definition has no size 

	//Multidimensional slices 
	sliceE := [][]string{
		{"lomie", "jame"},
		{"omakd", "loalkdo"},
	}

	memoryArray  := [20]int{1, 2, 3, 4, 5, 6, 7}
	tinySlice := memoryArray[1:4]  // just a tiny slice but still points to the main array 
	tinyCopy := make([]int, 4, 4) // cap can be omitted as it defaults to len, second param

	copy(tinyCopy, tinySlice)

	fmt.Println(sliceC[0])
	fmt.Println(len(sliceD))
	fmt.Println(sliceE[0][1])

	for i,v := range tinyCopy{
		fmt.Println(i)
		fmt.Println(v)
	}
}
