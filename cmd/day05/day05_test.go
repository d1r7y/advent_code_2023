package day05

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	type mapTest struct {
		mapRange            []Range
		source              int
		expectedDestination int
	}

	tests := []mapTest{
		{mapRange: []Range{{98, 50, 2}, {50, 52, 48}}, source: 0, expectedDestination: 0},
		{mapRange: []Range{{98, 50, 2}, {50, 52, 48}}, source: 1, expectedDestination: 1},
		{mapRange: []Range{{98, 50, 2}, {50, 52, 48}}, source: 49, expectedDestination: 49},
		{mapRange: []Range{{98, 50, 2}, {50, 52, 48}}, source: 50, expectedDestination: 52},
		{mapRange: []Range{{98, 50, 2}, {50, 52, 48}}, source: 51, expectedDestination: 53},
		{mapRange: []Range{{98, 50, 2}, {50, 52, 48}}, source: 60, expectedDestination: 62},
		{mapRange: []Range{{98, 50, 2}, {50, 52, 48}}, source: 96, expectedDestination: 98},
		{mapRange: []Range{{98, 50, 2}, {50, 52, 48}}, source: 97, expectedDestination: 99},
		{mapRange: []Range{{98, 50, 2}, {50, 52, 48}}, source: 98, expectedDestination: 50},
		{mapRange: []Range{{98, 50, 2}, {50, 52, 48}}, source: 99, expectedDestination: 51},
		{mapRange: []Range{{98, 50, 2}, {50, 52, 48}}, source: 79, expectedDestination: 81},
		{mapRange: []Range{{98, 50, 2}, {50, 52, 48}}, source: 14, expectedDestination: 14},
		{mapRange: []Range{{98, 50, 2}, {50, 52, 48}}, source: 55, expectedDestination: 57},
		{mapRange: []Range{{98, 50, 2}, {50, 52, 48}}, source: 13, expectedDestination: 13},
	}

	for _, test := range tests {
		m := NewMap()
		for _, r := range test.mapRange {
			m.AddRange(r)
		}

		assert.Equal(t, test.expectedDestination, m.Lookup(test.source))
	}
}

func TestParseRange(t *testing.T) {
	type parseRangeTest struct {
		line          string
		expectedRange Range
	}

	tests := []parseRangeTest{
		{line: "50 98 2", expectedRange: Range{98, 50, 2}},
		{line: "52 50 48", expectedRange: Range{50, 52, 48}},
	}

	for _, test := range tests {
		assert.Equal(t, test.expectedRange, ParseRange(test.line))
	}
}

func TestParseSeedsPartOne(t *testing.T) {
	type parseSeedsTest struct {
		line          string
		expectedSeeds []int
	}

	tests := []parseSeedsTest{
		{line: "seeds: 79 14 55 13", expectedSeeds: []int{79, 14, 55, 13}},
	}

	for _, test := range tests {
		a := NewAlmanac(true)

		ParseSeedsPartOne(a, test.line)

		seeds := make([]int, 0)

		for i := 0; i < a.GetSeedCount(); i++ {
			seeds = append(seeds, a.GetNextSeed())
		}

		assert.Equal(t, test.expectedSeeds, seeds)
	}
}

func TestParseSeedsPartTwo(t *testing.T) {
	line := "seeds: 222541566 218404460 670428364 432472902"

	a := NewAlmanac(true)

	ParseSeedsPartTwo(a, line)

	expectedSeeds := make([]int, 0)

	for i := 222541566; i < 222541566+218404460; i++ {
		expectedSeeds = append(expectedSeeds, i)
	}

	for i := 670428364; i < 670428364+432472902; i++ {
		expectedSeeds = append(expectedSeeds, i)
	}

	seeds := make([]int, 0)

	for i := 0; i < a.GetSeedCount(); i++ {
		seeds = append(seeds, a.GetNextSeed())
	}

	assert.Equal(t, expectedSeeds, seeds)
}

func TestParseAlmanac(t *testing.T) {
	content := `seeds: 79 14 55 13

	seed-to-soil map:
	50 98 2
	52 50 48
	
	soil-to-fertilizer map:
	0 15 37
	37 52 2
	39 0 15
	
	fertilizer-to-water map:
	49 53 8
	0 11 42
	42 0 7
	57 7 4
	
	water-to-light map:
	88 18 7
	18 25 70
	
	light-to-temperature map:
	45 77 23
	81 45 19
	68 64 13
	
	temperature-to-humidity map:
	0 69 1
	1 0 69
	
	humidity-to-location map:
	60 56 37
	56 93 4
	`

	almanac1 := ParseAlmanac(content, true)

	seeds1 := make([]int, 0)

	for i := 0; i < almanac1.GetSeedCount(); i++ {
		seeds1 = append(seeds1, almanac1.GetNextSeed())
	}

	assert.Equal(t, []int{79, 14, 55, 13}, seeds1)
	assert.Equal(t, sortRanges([]Range{{98, 50, 2}, {50, 52, 48}}), almanac1.SeedSoilMap.Ranges)
	assert.Equal(t, sortRanges([]Range{{15, 0, 37}, {52, 37, 2}, {0, 39, 15}}), almanac1.SoilFertilizerMap.Ranges)
	assert.Equal(t, sortRanges([]Range{{53, 49, 8}, {11, 0, 42}, {0, 42, 7}, {7, 57, 4}}), almanac1.FertilizerWaterMap.Ranges)
	assert.Equal(t, sortRanges([]Range{{18, 88, 7}, {25, 18, 70}}), almanac1.WaterLightMap.Ranges)
	assert.Equal(t, sortRanges([]Range{{77, 45, 23}, {45, 81, 19}, {64, 68, 13}}), almanac1.LightTemperatureMap.Ranges)
	assert.Equal(t, sortRanges([]Range{{69, 0, 1}, {0, 1, 69}}), almanac1.TemperatureHumidityMap.Ranges)
	assert.Equal(t, sortRanges([]Range{{56, 60, 37}, {93, 56, 4}}), almanac1.HumidityLocationMap.Ranges)

	almanac2 := ParseAlmanac(content, false)

	expectedSeeds2 := make([]int, 0)

	for i := 79; i < 79+14; i++ {
		expectedSeeds2 = append(expectedSeeds2, i)
	}

	for i := 55; i < 55+13; i++ {
		expectedSeeds2 = append(expectedSeeds2, i)
	}

	seeds2 := make([]int, 0)

	for i := 0; i < almanac2.GetSeedCount(); i++ {
		seeds2 = append(seeds2, almanac2.GetNextSeed())
	}

	assert.Equal(t, expectedSeeds2, seeds2)
	assert.Equal(t, sortRanges([]Range{{98, 50, 2}, {50, 52, 48}}), almanac2.SeedSoilMap.Ranges)
	assert.Equal(t, sortRanges([]Range{{15, 0, 37}, {52, 37, 2}, {0, 39, 15}}), almanac2.SoilFertilizerMap.Ranges)
	assert.Equal(t, sortRanges([]Range{{53, 49, 8}, {11, 0, 42}, {0, 42, 7}, {7, 57, 4}}), almanac2.FertilizerWaterMap.Ranges)
	assert.Equal(t, sortRanges([]Range{{18, 88, 7}, {25, 18, 70}}), almanac2.WaterLightMap.Ranges)
	assert.Equal(t, sortRanges([]Range{{77, 45, 23}, {45, 81, 19}, {64, 68, 13}}), almanac2.LightTemperatureMap.Ranges)
	assert.Equal(t, sortRanges([]Range{{69, 0, 1}, {0, 1, 69}}), almanac2.TemperatureHumidityMap.Ranges)
	assert.Equal(t, sortRanges([]Range{{56, 60, 37}, {93, 56, 4}}), almanac2.HumidityLocationMap.Ranges)
}

func TestGetLocation(t *testing.T) {
	content := `seeds: 79 14 55 13

	seed-to-soil map:
	50 98 2
	52 50 48
	
	soil-to-fertilizer map:
	0 15 37
	37 52 2
	39 0 15
	
	fertilizer-to-water map:
	49 53 8
	0 11 42
	42 0 7
	57 7 4
	
	water-to-light map:
	88 18 7
	18 25 70
	
	light-to-temperature map:
	45 77 23
	81 45 19
	68 64 13
	
	temperature-to-humidity map:
	0 69 1
	1 0 69
	
	humidity-to-location map:
	60 56 37
	56 93 4
	`

	almanac := ParseAlmanac(content, true)

	type getLocationTest struct {
		seed             int
		expectedLocation int
	}

	tests := []getLocationTest{
		{seed: 79, expectedLocation: 82},
		{seed: 14, expectedLocation: 43},
		{seed: 55, expectedLocation: 86},
		{seed: 13, expectedLocation: 35},
	}

	for _, test := range tests {
		assert.Equal(t, test.expectedLocation, almanac.GetLocation(test.seed))
	}
}
