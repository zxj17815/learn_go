package operator_test

import "testing"

func TestCompareArry(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 4, 5}
	//c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b)
	//t.Log(a == c)
	t.Log(a == d)
}

const (
	Executable = 1 << iota
	Writable
	Readable
)

func TestBitClear(t *testing.T) {
	t.Logf("%b,%b,%b", Readable, Writable, Executable)
	a := 7 //0111
	t.Log(a &^ Readable)
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
