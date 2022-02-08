package main

import "fmt"

func main() {
	age := 35
	name := "john"

	//Print (no new line)
	fmt.Print("Hello ")
	fmt.Print("World! \n")

	//Println (Prints with new line)
	fmt.Println("Hey there!")
	fmt.Println("My name is", name, "and my age is", age)

	//Printf (formatted strings)
	fmt.Printf("my age is %v and my name is %v \n", age, name)
	//%q puts double qoutes around strings
	fmt.Printf("my age is %q and my name is %q \n", age, name)
	//%T gives type of variable
	fmt.Printf("age is of type %T \n", age)
	//print float with %f
	fmt.Printf("you scored %f points! \n", 225.26)
	//round to decimal
	fmt.Printf("you scored %0.1f points! \n", 225.26)

	//Sprintf (save formatted strings)
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	fmt.Println("the saved string is:", str)
}
