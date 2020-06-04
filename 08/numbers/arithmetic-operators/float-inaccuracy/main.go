package main

import "fmt"

func main() {
	ratio := 1.0 / 10.0

	// after 10 operations
	// the inaccuracy is clear
	//
	// BTW, don't mind about this loop syntax for now
	// I'm going to explain it afterwards
	for range [...]int{10: 0} {
		ratio += 1.0 / 10.0
	}

	fmt.Printf("%.60f\n", ratio)

	// Go compiler sees these numbers as integers,
	//   since, there are no fractional parts in
	//   integer values,
	// So, the result becomes 1 instead of 1.5

	// So, ratio variable here is an int variable,
	//   it's because, 3 divided by 2 results
	//   in an integer.

	ratio2 := 3 / 2

	fmt.Printf("%d\n", ratio2)

	// When you use a float value with an integer
	// in a calculation,
	// the result always becomes a float.

	ratio3 := 3.0 / 2

	// OR:
	// ratio = 3 / 2.0

	// OR - if 3 is inside an int variable:
	// n := 3
	// ratio = float64(n) / 2

	fmt.Printf("%f\n", ratio3)
}
