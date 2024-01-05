package day06

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Race struct {
	Time     int
	Distance int
}

func ParseIntList(line string, prefix string) []int {
	intRE := regexp.MustCompile(`[0-9]+`)
	intMatches := intRE.FindAllString(strings.TrimPrefix(line, prefix), -1)

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

func ParseIntListRemovingAllWhitespace(line string, prefix string) []int {
	intRE := regexp.MustCompile(`[0-9]+`)
	intMatches := intRE.FindAllString(strings.Join(strings.Fields(strings.TrimPrefix(line, prefix)), ""), -1)

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

func SpeedAtTime(pressTime int, raceTime int) int {
	if pressTime < raceTime {
		return pressTime
	} else {
		return 0
	}
}

func DistanceAtTime(speed int, time int) int {
	return speed * time
}

func DoesBeatRecord(pressTime int, raceTime int, recordDistance int) bool {
	return DistanceAtTime(SpeedAtTime(pressTime, raceTime), raceTime-pressTime) > recordDistance
}

func ParseRaces(fileContents string, part1 bool) []Race {
	races := make([]Race, 0)

	for _, line := range strings.Split(fileContents, "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}

		switch {
		case strings.HasPrefix(strings.TrimSpace(line), "Time: "):
			var timeList []int
			if part1 {
				timeList = ParseIntList(line, "Time: ")
			} else {
				timeList = ParseIntListRemovingAllWhitespace(line, "Time: ")
			}

			for i, t := range timeList {
				if i >= len(races) {
					races = append(races, Race{})
				}

				races[i].Time = t
			}
		case strings.HasPrefix(strings.TrimSpace(line), "Distance: "):
			var distanceList []int
			if part1 {
				distanceList = ParseIntList(line, "Distance: ")
			} else {
				distanceList = ParseIntListRemovingAllWhitespace(line, "Distance: ")
			}

			for i, d := range distanceList {
				if i >= len(races) {
					races = append(races, Race{})
				}

				races[i].Distance = d
			}
		default:
			// Range line.
			log.Panicf("unexpected line '%s'\n", line)
		}
	}

	return races
}

func day06(fileContents string) error {
	races1 := ParseRaces(fileContents, true)

	// Part 1: Determine the number of ways you could beat the record in each race.
	// What do you get if you multiply these numbers together?
	totalWinningWays := 1

	for _, r := range races1 {
		beatRecordCount := 0
		for pt := 0; pt < r.Time; pt++ {
			if DoesBeatRecord(pt, r.Time, r.Distance) {
				beatRecordCount++
			}
		}

		totalWinningWays *= beatRecordCount
	}

	log.Printf("Total winning ways: %d\n", totalWinningWays)

	races2 := ParseRaces(fileContents, false)

	// Part 1: How many ways can you beat the record in this one much longer race?
	totalWinningWays = 1

	for _, r := range races2 {
		beatRecordCount := 0
		for pt := 0; pt < r.Time; pt++ {
			if DoesBeatRecord(pt, r.Time, r.Distance) {
				beatRecordCount++
			}
		}

		totalWinningWays *= beatRecordCount
	}

	log.Printf("Total winning ways: %d\n", totalWinningWays)

	return nil
}
