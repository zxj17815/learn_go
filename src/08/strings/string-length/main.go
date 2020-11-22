package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	name1 := "123456789İ你好?!"

	// strings are made up of bytes
	// len function counts the bytes in a string value
	fmt.Println(len(name1))
	name1 = "İ你好"
	fmt.Println(len(name1))

	name1 = "?!"
	fmt.Println(len(name1))
	// strings are made up of bytes

	// len function counts the bytes in a string value.
	//
	// This string literal contains unicode characters.
	//
	// And, unicode characters can be 1-4 bytes.
	// So, "İnanç" is 7 bytes long, not 5.
	//
	// İ = 2 bytes
	// n = 1 byte
	// a = 1 byte
	// n = 1 byte
	// ç = 2 bytes
	// TOTAL = 7 bytes
	name := "İnanç"
	fmt.Printf("%q is %d bytes\n", name, len(name))

	// To get the actual characters (or runes) inside
	// a utf-8 encoded string value, you should do this:
	fmt.Printf("%q is %d characters\n",
		name, utf8.RuneCountInString(name))

	msg := os.Args[1]

	// it's important to calculate things only once
	// so, do not call the repeat function twice
	// calling it once is enough
	marks := strings.Repeat("!", len(msg))
	s := marks + msg + marks
	s = strings.ToUpper(s)

	// you can also type this program more concisely
	// like this:
	//
	// msg := strings.ToUpper(os.Args[1])
	// marks := strings.Repeat("!", len(msg))
	fmt.Println(marks + msg + marks)
}
