package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Syafiq Fajrian Emha", "Syafiq"))
	fmt.Println(strings.Split("Syafiq, Fajrian, Emha", ","))
	fmt.Println(strings.ToLower("Syafiq Fajrian Emha"))
	fmt.Println(strings.ToUpper("Syafiq Fajrian Emha"))
	fmt.Println(strings.Trim("    Syafiq    ", " "))
	fmt.Println(strings.ReplaceAll("Syafiq Fajrian Emha", "a", "i"))
}
