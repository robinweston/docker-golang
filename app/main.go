package main

import (
	"github.com/fatih/color"
)

func Add(x int, y int) int {
	return x + y
}

func main() {
	result := Add(3,4)
	color.Blue("Hello world. The answer is %v", result)
}