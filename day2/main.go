package main

import (
	"Adventofcode2021/utils/files"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	//content := files.ReadFileString("input_test.txt")
	content := files.ReadFileString("input.txt")

	hVector := 0
	vVector := 0
	aim := 0
	foo := 0

	for _, x := range content {
		parts := strings.Split(x, " ")
		partInt, err := strconv.Atoi(parts[1])

		if err != nil {
			log.Fatal(err)
		}

		if parts[0] == "forward" {
			hVector += partInt
			foo = aim*partInt*-1 + foo

		} else if parts[0] == "down" {
			vVector -= partInt
		} else {
			vVector += partInt
		}
		aim = vVector
		fmt.Println("Aim is :", aim)
		fmt.Println("Multiople by aim", foo)
		fmt.Println("Forward is:", hVector)
		fmt.Println("Down/UP is:", vVector)

	}

	fmt.Println(foo * hVector)

}
