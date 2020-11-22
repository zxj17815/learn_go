package main

import "fmt"

func main() {
	var apple int
	var orange int32

	// ERROR:
	// cannot assign orange to apple (different types)
	// apple = orange

	// you need to convert orange to apple

	// orange is convertable to int because,
	//   int and int32 are both numeric types

	apple = int(orange)

	// you can't convert a numeric type to a bool:
	// isDelicious := bool(orange)

	// but you can convert an int to a string
	// this only works with int types，这里注意和python不一样，转换为string是这个数字对应的char字符
	orange = 65 // 65 is A
	color := string(orange)
	fmt.Println(color)
	fmt.Println(orange)

	// this doesn't work. 65.0 is a float.
	// fmt.Println(string(65.0))

	// this works: there are two byte values
	// byte is also an int
	fmt.Println(string([]byte{104, 105}))

	var test = 6.5

	fmt.Println(int(test))

	_ = apple
}
