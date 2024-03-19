# Golang

- Go is a statically/strictly typed language that also often feels dynamically typed sometimes.

- Package main is for executable. Package main must contain declaration of a main func 
func main() {

}

# declaring a variable:
var <variableName> <dataType> = value
e.g:
var name string = "zohaib"
var i int = 100
var b bool = false
var f float64 = 77.90

- OR 
  var name string
  name = "zohaib"

- OR 
  var greeting, name, subject string = "Hello", "Lisa", "Mathematics"

- OR
  var (
      name string = "zohaib"
      i int = 100
      b bool = false
      f float64 = 77.99
  )

# Short Variable Declaration:
Here we use := a.k.a short variable assignment operator
e.g 
name:= "Lisa"
name = "Peter"
fmt.Println(name)

# Local and Global variables
- Local Variables: declared inside and belong to inner block or Function only.
- Global Variables: declared outside of any func or outside of any block. For instance, those variables can be declared outside of the main() fun i.e on the top of the program and can be used anywhere within the program.

package main
import "fmt"

var global_variable string = "accessible anywhere within the program"

func main(){
    local_variable:= "accessible locally within main func block"
}

# Scope of Variables
- Inner block can access variables in the outer Block
- Outer block CANNOT access the variables in the inner block because variables inside the inner block are local variable and belong to inner block only

# format specifiers -- Use fmt.Printf
- %v - formats value in a default format
- %d - formats decimal integers
- %s - plain string
- %q - quoted characters/string
- %f - floating numbers
- $2f - floating numbers up to 2 decimal places
etc..

