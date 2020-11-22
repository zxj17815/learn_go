package main

import (
	"fmt" // You should replace this with your username

	golang "../golang"
	/*
		go 使用 package 来管理源文件。
		package 必须在一个文件夹内，且一个文件夹内也只能有一个package，但是一个文件夹可以有多个文件。
	*/)

func main() {
	fmt.Println(golang.Version())
}
