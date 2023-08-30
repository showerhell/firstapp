// Note: "package" declaration below
package main

// Note: "import" block below
import (
	"fmt"
)

// Note: "var" block below
// Declare Package Level value for variable i (cannot use ":=" syntax at package level)
var (
	i int = 10
)

func main() {
	fmt.Println("Hello to the interesting world of Golang, var i is set to 10")
	fmt.Println(i)
	// Override Package Level value of variable i by shadowing declaration of new value
	i := 20
	fmt.Println("var i now set to 20")
	fmt.Println(i)
	// Format string representation of v alue and T ype of variable i
	fmt.Printf("%v, %T", i, i)
}
