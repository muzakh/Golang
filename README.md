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
strconv.Itoa(<variable>) - for integer to string conversion. It returns only one value i.e converted string.
strconv.Atoi(<variable>) - for string to integer conversion. It returns two values: converted integer and error(if any).

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


# Constants:
- These are constant variables whose values never change. Reassignment is not possible
- No concept of "Zero values" or "Default values" in case of Constants.
- declared with a keyword as const <constantVariableName> = value
- Short Variable assignment := is not supported in Constants.
- For instance, assigning a constant value to Pii = 3.142, etc.

### Types of Constants:
- Typed: When we specify the data type specifically in Constant declaration for e.g, Const VarName <DataType> = value
- Untyped: When we dont specify the Data Type in declaration. Data Type in Constants in optional. They are flexible as compared to typed Constants.

- Constants variables give error if declared as follows:
### Wrong:
It only works in case of Variables declared by **var** keyword
Const variableName String
variableName = "Zohaib"
### Correct way:
Const variableName String = "Zohaib"

# Types of Operators:
- Comparison Operators: ==, !=, <, <=, >, >=
- Arithmetc Operators: +, -, *, /, %, ++, --
- Assignment Operators: =, +=, -=, *=, %=, /=
- Bitwise Operators: &, |, << (left shift operator), >> (right shift operator), (carrot sign - XOR). They operate on every bit level of the given operands.
- Logical Operators: && , ||, !


# If-else statement:
- The else keyword must start right after if ends otherwise the code throws syntaxt error:
"syntax error: unexpected else, expecting }"

