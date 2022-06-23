package main

import (
	"fmt"
	"strconv"
)

type variables string

func main() {
	var size, color, char variables
	size = "3"
	color = "\033[31m"
	char = "X"
	sandglass([]variables{size, color, char})
}

func sandglass(variable []variables) {
	size, _ := strconv.Atoi(string(variable[0]))
	color := variable[1]
	char := variable[2]
	for i := 1; i <= size; i++ {
		switch {
		case i == 1 || i == size:
			for j := 1; j <= size; j++ {
				fmt.Print(color, char)
			}
			fmt.Print("\n")
		case i >= 2 && i <= (size/2+size%2) && size%2 == 0:
			begEnd := i - 1
			middle := size - begEnd*2 - 2
			printSandglass(begEnd, middle, string(char))
		case i >= 2 && i < (size/2+size%2) && size%2 != 0:
			begEnd := i - 1
			middle := size - begEnd*2 - 2
			printSandglass(begEnd, middle, string(char))
		case i > (size/2+size%2) && i <= size-1:
			begEnd := size - i
			middle := size - begEnd*2 - 2
			printSandglass(begEnd, middle, string(char))
		default:
			begEnd := size / 2
			printSandglassMiddle(begEnd, string(char))
		}
	}
}

func printSandglass(begEnd int, middle int, char string) {
	for j := 1; j <= begEnd; j++ {
		fmt.Print(" ")
	}
	fmt.Print(char)
	for j := 1; j <= middle; j++ {
		fmt.Print(" ")
	}
	fmt.Print(char)
	fmt.Print("\n")
}

func printSandglassMiddle(begEnd int, char string) {
	for j := 1; j <= begEnd; j++ {
		fmt.Print(" ")
	}
	fmt.Print(char)
	fmt.Print("\n")
}
