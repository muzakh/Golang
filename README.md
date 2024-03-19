# Golang

# reflect.TypeOf 
import (
    "fmt"
    "reflect"
    "strconv" 
)

- fmt - Used for formating
- reflect.TypeOf(<variable>) returns the data type of string, integer, boolean etc.
- strconv - For String type casting
strconv.Itoa(<variable>) - for integer to string conversion.

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

# Zero values:
Happens when none of the value is assigned to a variable in declartion. It picks 0, false, null or nil depending upon the data type as follows:

int or %d = 0
float or %f = 0.0
%.2f = 0.00
string or %s = ""
bool = false
pointers, functions, interfaces, maps = nil

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
- %.2f - floating numbers up to 2 decimal places
- %t - boolean values - true or false
- %T - retuns DataType of a variable. Can also use reflect.TypeOf(<variable>)
etc..


# fmt.Scanf - read input from the user
- fmt.Scanf("%s %t", &name, &IsStudent)
Here "&" sign is mandatory. 

- fmt.Scanf returns two values "count" and "err":
- count - number of values that fmt.Scanf function writes on. In case of incorrect data-typed values it returns an error and doesn't performs the increment to the count value. 
- err - this is to log the error whenever wrong datatype values is provided for a user input and in that case a "Zero Value" is then assigned to that variable depending upon its type.

var stringValue string  
var integerValue int 

count, err := fmt.Scanf("%s %d", &stringValue, &integerValue)

fmt.Println("count: ", count)
fmt.Println("error: ", err)
fmt.Println("stringValue: ", stringValue)
fmt.Println("integerValue: ", integerValue)

###Output:
count: 1
error: expected integer
stringValue: Zohaib
integerValue: 0


# Type Casting:
simply assign a variable enclosed in parenthesis. 
e.g 
int(floatVariable) - Converting float to variable.
float64(int) - Converting integer into float.

