package files

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadFileInt(input string) []int {

	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	sc := bufio.NewScanner(file)

	var values []int

	for sc.Scan() {
		num, _ := strconv.Atoi(sc.Text())
		values = append(values, num)

	}
	return values

}

func ReadFileString(input string) []string {
	bytes, err := ioutil.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}

	content := string(bytes)

	split := strings.Split(content, "\n")
	split = split[:len(split)-1]

	return split

}
