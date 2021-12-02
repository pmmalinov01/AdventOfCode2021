package main

import (
	"Adventofcode2021/utils/files"
	"fmt"
	"strconv"
)

func main() {
	part1()
	part2()
}

func part1() {
	values := files.ReadFileInt("input.txt")

	inc := 0
	last := -1

	for _, v := range values {
		if last != -1 && v > last {
			inc++
		}
		last = v
	}

	fmt.Println(inc)

}

func part2() {
	values := files.ReadFileInt("input.txt")

	inc := 0

	for s := 0; s < len(values)-3; s++ {
		foo := values[s] + values[s+1] + values[s+2]
		bar := values[s+1] + values[s+2] + values[s+3]
		if foo < bar {
			inc++
		}
	}

	fmt.Println(inc)

	what := files.ReadFileString("input.txt")

	for i, x := range what {
		t, _ := strconv.Atoi(x)

		fmt.Println("Before Conversion")
		fmt.Printf("Type of %s is %T\n", what[i], x)
		fmt.Println("After Conversion")
		fmt.Printf("Type of %s is %T\n", what[i], t)
	}

}
