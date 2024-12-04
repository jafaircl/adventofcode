package main

import "fmt"

func Hello(name string) string {
	result := "Hello " + name
	return result
}

func main() {
	fmt.Println(Hello("aoc2024-go"))
}
