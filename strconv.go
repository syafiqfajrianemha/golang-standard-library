package main

import (
	"fmt"
	"strconv"
)

func main() {
	// strconv (string conversion)
	result, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println("Error", err.Error())
	}

	resultInt, err := strconv.Atoi("1000")
	if err == nil {
		fmt.Println(resultInt)
	} else {
		fmt.Println("Error", err.Error())
	}

	binary := strconv.FormatInt(999, 2)
	fmt.Println(binary)

	intString := strconv.Itoa(999)
	fmt.Println(intString)
}
