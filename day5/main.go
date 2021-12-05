package main

import (
	"Adventofcode2021/utils/files"
	"fmt"
	"strconv"
	"strings"
)

type Maze struct {
	X, Y int
}

func main() {
	clouds := make(map[Maze]int)
	fmt.Println("vim-go")
	contents := files.ReadFileString("input.txt")

	for _, s := range contents {
		parts := strings.Split(s, " -> ")
		sRaw := strings.Split(parts[0], ",")
		eRaw := strings.Split(parts[1], ",")
		sX, _ := strconv.Atoi(sRaw[0])
		sY, _ := strconv.Atoi(sRaw[1])
		eX, _ := strconv.Atoi(eRaw[0])
		eY, _ := strconv.Atoi(eRaw[1])
		if sX == eX {
			if sY < eY {
				for y := sY; y <= eY; y++ {
					clouds[Maze{sX, y}]++
				}
			} else {
				for y := sY; y >= eY; y-- {
					clouds[Maze{sX, y}]++
				}

			}
		} else if sY == eY {

			if sX < eX {
				for x := sX; x <= eX; x++ {
					clouds[Maze{x, sY}]++
				}
			} else {
				for x := sX; x >= eX; x-- {
					clouds[Maze{x, sY}]++
				}

			}
		} else {
			// Nope
		}

	}
	count := 0
	for _, v := range clouds {
		if v >= 2 {
			count++
		}
	}
	fmt.Println(count)
}
