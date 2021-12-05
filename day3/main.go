package main

import (
	"Adventofcode2021/utils/files"
	"fmt"
	"strconv"
)

func main() {
	source := files.ReadFileString("input.txt")

	var ones [12]int
	source = source[:len(source)-1]
	for _, v := range source {

		for k, x := range v {
			if x == '1' {
				ones[k]++
			}
		}
	}

	var gamma, eps string
	for _, v := range ones {
		if v > len(source)/2 {
			gamma += "1"
			eps += "0"

		} else {
			gamma += "0"
			eps += "1"

		}
	}

	nGamma, _ := strconv.ParseInt(gamma, 2, 64)
	nEps, _ := strconv.ParseInt(eps, 2, 64)
	fmt.Println(nGamma * nEps)

	x := search(source, oxygenFilter) * search(source, co2Filter)
	fmt.Println(x)
}

func oxygenFilter(most, _ rune) rune {
	return most
}

func co2Filter(_, least rune) rune {
	return least
}

func search(input []string, filter func(rune, rune) rune) int {
	remaining := make([]string, len(input))
	copy(remaining, input)

	offset := 0

	for len(remaining) > 1 {
		left := make([]string, 0, len(remaining))
		target := filter(mostLeast(remaining, offset))

		for _, line := range remaining {
			if rune(line[offset]) == target {
				left = append(left, line)
			}
		}

		remaining = left
		offset++
	}

	v, _ := strconv.ParseInt(remaining[0], 2, 32)
	return int(v)

}

func mostLeast(input []string, idx int) (rune, rune) {
	zeroes := 0
	ones := 0

	for _, line := range input {
		if line[idx] == '0' {
			zeroes++
		} else {
			ones++
		}
	}

	if zeroes > ones {
		return '0', '1'
	}

	return '1', '0'
}
