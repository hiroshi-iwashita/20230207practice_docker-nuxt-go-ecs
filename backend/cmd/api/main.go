package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Hello Golang world.")
	fmt.Println("Go version :", runtime.Version(), ", OS :", runtime.GOOS, ", ARCH :", runtime.GOARCH)
}