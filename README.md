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

# format specifiers -- Use fmt.Printf
- %v - formats value in a default format
- %d - formats decimal integers
- %s - plain string
- %q - quoted characters/string
- %f - floating numbers
- $2f - floating numbers up to 2 decimal places
etc..

