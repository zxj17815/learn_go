package first_test

import (
	"fmt"
	"testing"
)

//尝试单元测试
func TestFirstTry(t *testing.T) {
	var a = 1
	var b = 1
	for i := 0; i <= 5; i++ {
		fmt.Print(" ", b)
		tmp := a
		a = b
		b = tmp + a
		fmt.Println()
	}
}

//简单交换数值
func TestExchange(t *testing.T) {
	a := 1
	b := 2
	t.Log(a, b)
	a, b = b, a
	t.Log(a, b)
}

//快速赋值
const (
	Monday = 1 + iota
	Tuesday
	Wednesday
)
const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestFastAssignment(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday)
	a := 1
	t.Log(a & Writable)
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
