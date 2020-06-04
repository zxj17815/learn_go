package main

import (
	"fmt"
	"os"
)

// STEPS:
//
// Compile it by typing:
//   go build -o myprogram
//
// Then run it by typing:
//   ./myprogram
//
// If you're on Windows, then type:
//   myprogram
// 编译为windows的exe程序：CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o myprogram.exe
func main() {
	fmt.Println("This file path:", os.Args[0])
}
