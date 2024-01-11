package day12

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseLine(t *testing.T) {
	type testCase struct {
		line          string
		expectedGroup *SpringGroup
	}
	testCases := []testCase{
		{
			line: "#.#.### 1,1,3",
			expectedGroup: &SpringGroup{
				Unfolded:          1,
				States:            []SpringState{Broken, Operational, Broken, Operational, Broken, Broken, Broken},
				DamagedSpringRuns: []int{1, 1, 3},
			},
		},
		{
			line: ".#...#....###. 1,1,3",
			expectedGroup: &SpringGroup{
				Unfolded:          1,
				States:            []SpringState{Operational, Broken, Operational, Operational, Operational, Broken, Operational, Operational, Operational, Operational, Broken, Broken, Broken, Operational},
				DamagedSpringRuns: []int{1, 1, 3},
			},
		},
		{
			line: "???.### 1,1,3",
			expectedGroup: &SpringGroup{
				Unfolded:          1,
				States:            []SpringState{Unknown, Unknown, Unknown, Operational, Broken, Broken, Broken},
				DamagedSpringRuns: []int{1, 1, 3},
			},
		},
		{
			line: "?###???????? 3,2,1",
			expectedGroup: &SpringGroup{
				Unfolded:          1,
				States:            []SpringState{Unknown, Broken, Broken, Broken, Unknown, Unknown, Unknown, Unknown, Unknown, Unknown, Unknown, Unknown},
				DamagedSpringRuns: []int{3, 2, 1},
			},
		},
	}

	for _, test := range testCases {
		group := ParseLine(test.line)
		assert.Equal(t, test.expectedGroup, group)
	}
}

func TestSimplifySpringStateRun(t *testing.T) {
	type testCase struct {
		state           SpringStateList
		interestedState SpringState
		expectedRun     []int
	}
	testCases := []testCase{
		{
			state:           SpringStateList{Broken, Operational, Broken, Operational, Broken, Broken, Broken},
			interestedState: Broken,
			expectedRun:     []int{1, 1, 3},
		},
		{
			state:           SpringStateList{Operational, Broken, Operational, Operational, Operational, Broken, Operational, Operational, Operational, Operational, Broken, Broken, Broken, Operational},
			interestedState: Broken,
			expectedRun:     []int{1, 1, 3},
		},
		{
			state:           SpringStateList{Broken, Broken, Broken, Operational, Broken, Broken, Broken},
			interestedState: Broken,
			expectedRun:     []int{3, 3},
		},
		{
			state:           SpringStateList{Operational, Broken, Broken, Broken, Operational, Operational, Operational, Broken, Broken, Operational, Operational, Operational, Operational, Broken, Operational, Operational, Operational, Operational, Operational},
			interestedState: Operational,
			expectedRun:     []int{1, 3, 4, 5},
		},
	}

	for _, test := range testCases {
		assert.Equal(t, test.expectedRun, SimplifySpringStateRun(test.state, test.interestedState))
	}

}

func TestSpringGroupSolve(t *testing.T) {
	type testCase struct {
		line                 string
		expectedAlternatives []*SpringGroup
	}
	testCases := []testCase{
		{
			line: "???.### 1,1,3",
			expectedAlternatives: []*SpringGroup{
				{Unfolded: 1, DamagedSpringRuns: []int{1, 1, 3}, States: SpringStateList{Broken, Operational, Broken, Operational, Broken, Broken, Broken}},
			},
		},
		{
			line: ".??..??...?##. 1,1,3",
			expectedAlternatives: []*SpringGroup{
				{Unfolded: 1, DamagedSpringRuns: []int{1, 1, 3}, States: SpringStateList{Operational, Broken, Operational, Operational, Operational, Broken, Operational, Operational, Operational, Operational, Broken, Broken, Broken, Operational}},
				{Unfolded: 1, DamagedSpringRuns: []int{1, 1, 3}, States: SpringStateList{Operational, Broken, Operational, Operational, Operational, Operational, Broken, Operational, Operational, Operational, Broken, Broken, Broken, Operational}},
				{Unfolded: 1, DamagedSpringRuns: []int{1, 1, 3}, States: SpringStateList{Operational, Operational, Broken, Operational, Operational, Broken, Operational, Operational, Operational, Operational, Broken, Broken, Broken, Operational}},
				{Unfolded: 1, DamagedSpringRuns: []int{1, 1, 3}, States: SpringStateList{Operational, Operational, Broken, Operational, Operational, Operational, Broken, Operational, Operational, Operational, Broken, Broken, Broken, Operational}},
			},
		},
	}

	for _, test := range testCases {
		group := ParseLine(test.line)
		assert.Equal(t, test.expectedAlternatives, group.Solve())
	}
}

func TestSpringGroupSolveUnfolded(t *testing.T) {
	line := "?###???????? 3,2,1"
	group := ParseLine(line)
	group.Unfold(5).Solve()
}

func TestSpringGroupSolveCount(t *testing.T) {
	type testCase struct {
		line                      string
		expectedAlternativesCount int
	}
	testCases := []testCase{
		{line: "???.### 1,1,3", expectedAlternativesCount: 1},
		{line: ".??..??...?##. 1,1,3", expectedAlternativesCount: 4},
		{line: "?###???????? 3,2,1", expectedAlternativesCount: 10},
	}

	for _, test := range testCases {
		group := ParseLine(test.line)
		assert.Equal(t, test.expectedAlternativesCount, len(group.Solve()))
	}
}

func TestSpringGroupUnfold(t *testing.T) {
	type testCase struct {
		line         string
		expectedLine string
	}
	testCases := []testCase{
		{line: ".# 1", expectedLine: ".#?.#?.#?.#?.# 1,1,1,1,1"},
		{line: "???.### 1,1,3", expectedLine: "???.###????.###????.###????.###????.### 1,1,3,1,1,3,1,1,3,1,1,3,1,1,3"},
	}

	for _, test := range testCases {
		group := ParseLine(test.line)
		unfoldedGroup := group.Unfold(5)
		assert.Equal(t, test.expectedLine, unfoldedGroup.Describe())
	}
}
