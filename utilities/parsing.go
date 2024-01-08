package utilities

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

func ParseIntList(line string) []int {
	intRE := regexp.MustCompile(`[-]?[0-9]+`)
	intMatches := intRE.FindAllString(line, -1)

	intList := make([]int, 0)

	for _, s := range intMatches {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}

		intList = append(intList, i)
	}

	return intList
}

func ParseIntListRemovingAllWhitespace(line string) []int {
	intRE := regexp.MustCompile(`[-]?[0-9]+`)
	intMatches := intRE.FindAllString(strings.Join(strings.Fields(line), ""), -1)

	intList := make([]int, 0)

	for _, s := range intMatches {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}

		intList = append(intList, i)
	}

	return intList
}
