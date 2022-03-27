package main

import (
	"fmt"

	"github.com/abishpius/calculator"
)

func main() {
	total := calculator.Sum(3, 5)
	fmt.Println(total)
	fmt.Println("Version: ", calculator.logMessage)
	fmt.Println("Hello World!")
}
