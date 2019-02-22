package helloworld

import "fmt"

const (
	packageVersion = "1.0.0"
)

// HelloWorld - print greeting
func HelloWorld() {
	fmt.Printf("%s (%s)!\n", "Hello world", packageVersion)
}
