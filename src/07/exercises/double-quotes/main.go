package main

import "fmt"

func main() {
	fmt.Printf("%v\n", "hello world")
	fmt.Printf("%s\n", "hello world")
	// %q会加上双引号,对于string
	fmt.Printf("%q\n", "123")
}
