package main

import (
	"fmt"
)

func main() {
	total := calculator.internalSum(5)
	fmt.Println(total)
	fmt.Println("Version: ", calculator.logMessage)
}
