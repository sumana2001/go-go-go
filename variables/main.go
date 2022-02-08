package main

import "fmt"

func main() {
	//strings
	var nameOne string = "tina"

	//go figures out the datatype
	var nameTwo = "john"

	//add value later
	var nameThree string
	fmt.Println(nameOne, nameTwo, nameThree)

	//update string value
	nameOne = "peach"
	//add string value
	nameThree = "hehe"
	fmt.Println(nameOne, nameTwo, nameThree)

	//shorthand version but can't use it outside a function
	nameFour := "blah"
	fmt.Println(nameFour)

	//integers
	var ageOne int = 25
	var ageTwo = 30
	ageThree := 40
	fmt.Println(ageOne, ageTwo, ageThree)

	//bits & memory (use different ints to get different ranges)
	//ranges: https://go.dev/ref/spec#Numeric_types
	var numOne int8 = 25
	var numTwo int16 = 250

	//only positive numbers
	var numThree uint = 30
	//since negative numbers are out, range doubles on positive side
	var numFour uint8 = 255
	fmt.Println(numOne, numTwo, numThree, numFour)

	//float
	var scoreOne float32 = 25.98
	var scoreTwo float64 = 37873862876386625.72677576585

	//will be infered as float64
	scoreThree := 1.5
	fmt.Println(scoreOne, scoreTwo, scoreThree)
}
