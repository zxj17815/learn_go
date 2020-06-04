// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package golang

import (
	"runtime"
)

// Version returns the current Go version，go语言的包里只有大写开头的值和方法可以导出
func Version() string {
	return runtime.Version()
}
