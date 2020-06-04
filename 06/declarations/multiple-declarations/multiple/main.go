package main

import "fmt"

func main() {
	var (
		speed int
		heat  float64

		off   bool
		brand string
	)

	var speed2, velocity int
	// this is equal to:
	//
	//   var (
	//     speed int
	//     velocity int
	//   )
	//
	// or:
	//
	//   var speed int
	//   var velocity int
	//

	fmt.Println(speed)
	fmt.Println(heat)
	fmt.Println(off)
	fmt.Printf("%q\n", brand)
	fmt.Println(speed2)
	fmt.Println(velocity)
}
