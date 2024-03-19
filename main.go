// My first Go code

// Here main package represent that the code is executable instead of a shared library
// Go is statically typed language. Variables are explicitly and 
package main

import "fmt" 
// OR import ("fmt")

func main()  {
	name:= "Lisa"  // This is dynamically typing
	var greeting string = "Hello" // This is statically typing
	fmt.Println(greeting)
	fmt.Println(name)

	
}

