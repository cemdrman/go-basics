package variables

import "fmt"

func main() {

	var x int
	x = 10
	fmt.Print(x)

	a := 20
	fmt.Print(a)
	a = 30
	fmt.Print(a)

	// string
	name := "bilisimio"
	fmt.Println(name)

	// float64
	pi := 3.14
	fmt.Printf(pi)

	// bool
	verdade := true
	mentira := false

	fmt.Println(verdade == mentira)
}
