package main

import (
	"fmt"
	"runtime"
)

// application is the type for all data we want to share with the
// various parts of our application. We will share this information
// in most cases by using this type as the receiver for functions.
// type applicationConfig struct {
// 	port         int
// 	infoLog      *log.Logger
// 	errorLog     *log.Logger
// 	// models       models.Models
// 	environment  string
// 	inProduction bool
// }


func main() {
	fmt.Println("Hello Golang world.")
	fmt.Println("Go version :", runtime.Version(), ", OS :", runtime.GOOS, ", ARCH :", runtime.GOARCH)
}