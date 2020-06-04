package main

import "fmt"

func main() {
	fmt.Println(
		(2+2)*4/2,
		(2+2)*(4/2), // same as above
	)

	n, m := 1, 5

	fmt.Println(2 + 1*m/n)         //7
	fmt.Println(2 + ((1 * m) / n)) // same as above

	// let's change the precedence using parentheses
	fmt.Println(((2 + 1) * m) / n)

	celsius := 35.

	// Wrong formula  :  9*celsius + 160  / 5
	// Correct formula: (9*celsius + 160) / 5
	fahrenheit := (9*celsius + 160) / 5

	fmt.Printf("%g ºC is %g ºF\n", celsius, fahrenheit)
}
