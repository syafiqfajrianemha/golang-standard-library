package main

import "fmt"

func main() {
	firstName := "Syafiq Fajrian"
	lastName := "Emha"

	fmt.Println("Hello '", firstName, lastName, "'")

	fmt.Printf("Hello '%s %s'\n", firstName, lastName)
}
