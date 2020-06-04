package main

import (
	"fmt"
	"path"
)

func main() {
	// var dir, file string
	// 路径分割
	// dir, file = path.Split("abc/css/main.css")
	// 简写
	dir, file := path.Split("css/main.css")
	fmt.Println("dir :", dir)
	fmt.Println("file:", file)
}
