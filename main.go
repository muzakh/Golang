// My first Go code

// Here main package represent that the code is executable instead of a shared library
// Go is statically typed language. Variables are explicitly and 
package main

import "fmt" 
// OR import ("fmt")

func main()  {
	name:= "Lisa"  // This is dynamically typing
	var greeting string = "Hello" // This is statically typing
	var subject string = "Mathematics"
	// OR var greeting, name, subject string = "Hello", "Lisa", "Mathematics"

	var marks int = 82 
	/*
OR
var marks int
marks = 82
	*/


/*
var (
	greeting string = "Lisa"
    subject string = "Mathematics"
	marks int = 82
)

*/

	fmt.Println(greeting) 	// OR fmt.Print(greeting, "\n")
	fmt.Println(name)
	fmt.Printf("Congrats! You have scored %d marks in %v", marks, subject) // fmt.Printf is to enable format specifie else it will mot work



	
}

