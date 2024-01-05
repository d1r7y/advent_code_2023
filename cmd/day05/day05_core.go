package day05

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type Range struct {
	SourceStart      int
	DestinationStart int
	Count            int
}

func ParseRange(line string) Range {
	rangeRE := regexp.MustCompile(`[0-9]+`)

	rangeMatches := rangeRE.FindAllString(line, -1)
	if len(rangeMatches) != 3 {
		log.Fatal(fmt.Errorf("invalid line '%s'", line))
	}

	destinationStart, err := strconv.Atoi(rangeMatches[0])
	if err != nil {
		log.Fatal(err)
	}

	sourceStart, err := strconv.Atoi(rangeMatches[1])
	if err != nil {
		log.Fatal(err)
	}

	count, err := strconv.Atoi(rangeMatches[2])
	if err != nil {
		log.Fatal(err)
	}

	return Range{SourceStart: sourceStart, DestinationStart: destinationStart, Count: count}
}

func ParseSeedsPartOne(almanac *Almanac, line string) {
	seedRE := regexp.MustCompile(`[0-9]+`)

	seedMatches := seedRE.FindAllString(strings.TrimPrefix(line, "seeds: "), -1)

	for _, s := range seedMatches {
		seed, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}

		almanac.AddSeedRange(seed, 1)
	}
}

func ParseSeedsPartTwo(almanac *Almanac, line string) {
	seedRE := regexp.MustCompile(`[0-9]+`)

	seedMatches := seedRE.FindAllString(strings.TrimPrefix(line, "seeds: "), -1)

	for i := 0; i < len(seedMatches); {
		start, err := strconv.Atoi(seedMatches[i])
		if err != nil {
			log.Fatal(err)
		}

		count, err := strconv.Atoi(seedMatches[i+1])
		if err != nil {
			log.Fatal(err)
		}

		almanac.AddSeedRange(start, count)
		i += 2
	}
}

type Map struct {
	Ranges []Range
}

func NewMap() *Map {
	return &Map{}
}

func (m *Map) AddRange(r Range) {
	m.Ranges = append(m.Ranges, r)

	sortRanges(m.Ranges)
}

func sortRanges(ranges []Range) []Range {
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].SourceStart < ranges[j].SourceStart
	})
	return ranges
}

func (m *Map) Lookup(source int) int {
	for _, r := range m.Ranges {
		if source >= r.SourceStart && source < r.SourceStart+r.Count {
			return r.DestinationStart + (source - r.SourceStart)
		}
	}

	return source
}

type SeedsPartOne struct {
	SeedList     []int
	CurrentIndex int
}

func NewSeedsPartOne() *SeedsPartOne {
	return &SeedsPartOne{SeedList: make([]int, 0)}
}

func (s *SeedsPartOne) ResetSeeds() {
	s.CurrentIndex = 0
}

func (s *SeedsPartOne) AddSeedRange(start int, count int) {
	for i := start; i < start+count; i++ {
		s.SeedList = append(s.SeedList, i)
	}
}

func (s *SeedsPartOne) GetSeedCount() int {
	return len(s.SeedList)
}

func (s *SeedsPartOne) GetNextSeed() int {
	nextSeed := s.SeedList[s.CurrentIndex]
	s.CurrentIndex++

	return nextSeed
}

type SeedRange struct {
	Start int
	Count int
}

type SeedsPartTwo struct {
	SeedList              []SeedRange
	SeedCount             int
	CurrentSeedRangeIndex int
	CurrentSeedIndex      int
}

func NewSeedsPartTwo() *SeedsPartTwo {
	return &SeedsPartTwo{SeedList: make([]SeedRange, 0)}
}

func (s *SeedsPartTwo) ResetSeeds() {
	s.CurrentSeedRangeIndex = 0
	s.CurrentSeedIndex = 0
}

func (s *SeedsPartTwo) AddSeedRange(start int, count int) {
	s.SeedList = append(s.SeedList, SeedRange{start, count})
	s.SeedCount += count
}

func (s *SeedsPartTwo) GetSeedCount() int {
	return s.SeedCount
}

func (s *SeedsPartTwo) GetNextSeed() int {
	if len(s.SeedList) == 0 {
		log.Panic("GetNextSeed() called before adding any seeds")
	}

	for {
		currentRange := s.SeedList[s.CurrentSeedRangeIndex]

		if s.CurrentSeedIndex < currentRange.Count {
			seed := currentRange.Start + s.CurrentSeedIndex
			s.CurrentSeedIndex++

			return seed
		} else {
			s.CurrentSeedRangeIndex++
			log.Printf("GetNextSeed() range index: %d\n", s.CurrentSeedRangeIndex)
			s.CurrentSeedIndex = 0
		}
	}
}

type GetSeeds interface {
	ResetSeeds()
	AddSeedRange(start int, count int)
	GetSeedCount() int
	GetNextSeed() int
}

type Almanac struct {
	Seeds                  []int
	Seeds2                 GetSeeds
	SeedSoilMap            *Map
	SoilFertilizerMap      *Map
	FertilizerWaterMap     *Map
	WaterLightMap          *Map
	LightTemperatureMap    *Map
	TemperatureHumidityMap *Map
	HumidityLocationMap    *Map
}

func NewAlmanac(partOne bool) *Almanac {
	var seeds GetSeeds

	if partOne {
		seeds = &SeedsPartOne{}
	} else {
		seeds = &SeedsPartTwo{}
	}

	return &Almanac{
		Seeds:                  make([]int, 0),
		Seeds2:                 seeds,
		SeedSoilMap:            NewMap(),
		SoilFertilizerMap:      NewMap(),
		FertilizerWaterMap:     NewMap(),
		WaterLightMap:          NewMap(),
		LightTemperatureMap:    NewMap(),
		TemperatureHumidityMap: NewMap(),
		HumidityLocationMap:    NewMap(),
	}
}

func (a *Almanac) GetLocation(seed int) int {
	soil := a.SeedSoilMap.Lookup(seed)
	fertilizer := a.SoilFertilizerMap.Lookup(soil)
	water := a.FertilizerWaterMap.Lookup(fertilizer)
	light := a.WaterLightMap.Lookup(water)
	temperature := a.LightTemperatureMap.Lookup(light)
	humidity := a.TemperatureHumidityMap.Lookup(temperature)
	location := a.HumidityLocationMap.Lookup(humidity)

	return location
}

func (a *Almanac) ResetSeed() {
	a.Seeds2.ResetSeeds()
}

func (a *Almanac) AddSeedRange(start int, count int) {
	p := message.NewPrinter(language.English)
	log.Print(p.Sprintf("Seed range: %d %d\n", start, count))
	a.Seeds2.AddSeedRange(start, count)
}

func (a *Almanac) GetSeedCount() int {
	return a.Seeds2.GetSeedCount()
}

func (a *Almanac) GetNextSeed() int {
	return a.Seeds2.GetNextSeed()
}

func ParseAlmanac(fileContents string, partOne bool) *Almanac {
	almanac := NewAlmanac(partOne)

	currentMap := almanac.SeedSoilMap

	for _, line := range strings.Split(fileContents, "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}

		switch {
		case strings.HasPrefix(strings.TrimSpace(line), "seeds: "):
			if partOne {
				ParseSeedsPartOne(almanac, line)
			} else {
				ParseSeedsPartTwo(almanac, line)
			}
		case strings.HasPrefix(strings.TrimSpace(line), "seed-to-soil map:"):
			currentMap = almanac.SeedSoilMap
		case strings.HasPrefix(strings.TrimSpace(line), "soil-to-fertilizer map:"):
			currentMap = almanac.SoilFertilizerMap
		case strings.HasPrefix(strings.TrimSpace(line), "fertilizer-to-water map:"):
			currentMap = almanac.FertilizerWaterMap
		case strings.HasPrefix(strings.TrimSpace(line), "water-to-light map:"):
			currentMap = almanac.WaterLightMap
		case strings.HasPrefix(strings.TrimSpace(line), "light-to-temperature map:"):
			currentMap = almanac.LightTemperatureMap
		case strings.HasPrefix(strings.TrimSpace(line), "temperature-to-humidity map:"):
			currentMap = almanac.TemperatureHumidityMap
		case strings.HasPrefix(strings.TrimSpace(line), "humidity-to-location map:"):
			currentMap = almanac.HumidityLocationMap
		default:
			// Range line.
			currentMap.AddRange(ParseRange(strings.TrimSpace(line)))
		}
	}

	return almanac
}

func day05(fileContents string) error {
	almanac1 := ParseAlmanac(fileContents, true)

	// Part 1: What is the lowest location number that corresponds to any of the initial seed numbers?
	lowestLocation := math.MaxInt

	for i := 0; i < almanac1.GetSeedCount(); i++ {
		location := almanac1.GetLocation(almanac1.GetNextSeed())

		if location < lowestLocation {
			lowestLocation = location
		}

	}

	log.Printf("Lowest location: %d\n", lowestLocation)

	// Part 2: Consider all of the initial seed numbers listed in the ranges on the first line of the almanac.
	// What is the lowest location number that corresponds to any of the initial seed numbers?
	almanac2 := ParseAlmanac(fileContents, false)

	lowestLocation2 := math.MaxInt

	for i := 0; i < almanac2.GetSeedCount(); i++ {
		location := almanac2.GetLocation(almanac2.GetNextSeed())

		if location < lowestLocation2 {
			lowestLocation2 = location
		}

	}

	log.Printf("Lowest location: %d\n", lowestLocation2)

	return nil
}
