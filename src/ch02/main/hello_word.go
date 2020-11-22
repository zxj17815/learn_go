package main

import (
	"fmt"
	"os"
)

func main() {
	test := "\"Hello\"" + ` \"World\"`
	fmt.Println(test)
	os.Exit(-1)
}
