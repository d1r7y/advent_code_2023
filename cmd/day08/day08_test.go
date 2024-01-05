package day08

import (
	"log"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseDirection(t *testing.T) {
	type testCase struct {
		str                string
		expectedDirections []Direction
	}

	testCases := []testCase{
		{"RL", []Direction{Right, Left}},
		{"LLR", []Direction{Left, Left, Right}},
		{"L", []Direction{Left}},
		{"R", []Direction{Right}},
		{"LLLL", []Direction{Left, Left, Left, Left}},
		{"RRRR", []Direction{Right, Right, Right, Right}},
	}

	for _, test := range testCases {
		directions := ParseDirections(test.str)
		assert.Equal(t, test.expectedDirections, directions, test.str)
	}
}

func TestParseNodeDescription(t *testing.T) {
	type testCase struct {
		str                     string
		expectedNodeDescription NodeDescription
	}

	testCases := []testCase{
		{"FTD = (QRN, JJC)", NodeDescription{"FTD", "QRN", "JJC"}},
		{"AAA = (BBB, CCC)", NodeDescription{"AAA", "BBB", "CCC"}},
		{"CKF = (XCC, SGZ)", NodeDescription{"CKF", "XCC", "SGZ"}},
		{"GGG = (GGG, GGG)", NodeDescription{"GGG", "GGG", "GGG"}},
		{"ZZZ = (ZZZ, ZZZ)", NodeDescription{"ZZZ", "ZZZ", "ZZZ"}},
	}

	for _, test := range testCases {
		description := ParseNodeDescription(test.str)
		assert.Equal(t, test.expectedNodeDescription, description, test.str)
	}
}

func TestWalk(t *testing.T) {
	type testCase struct {
		content       string
		expectedSteps int
	}

	testCases := []testCase{
		{`RL

	AAA = (BBB, CCC)
	BBB = (DDD, EEE)
	CCC = (ZZZ, GGG)
	DDD = (DDD, DDD)
	EEE = (EEE, EEE)
	GGG = (GGG, GGG)
	ZZZ = (ZZZ, ZZZ)`, 2},
		{`LLR

	AAA = (BBB, BBB)
	BBB = (AAA, ZZZ)
	ZZZ = (ZZZ, ZZZ))`, 6},
	}

	for _, test := range testCases {
		lines := strings.Split(test.content, "\n")

		directions := ParseDirections(lines[0])

		if lines[1] != "" {
			log.Panicf("unexpected non blank line in input: '%s'\n", lines[1])
		}

		n := ParseNetwork(lines[2:])

		steps := n.Walk(n.Find("AAA"), directions, n.Find("ZZZ"))

		assert.Equal(t, test.expectedSteps, steps)
	}

}

func TestGhostWalk(t *testing.T) {
	type testCase struct {
		content       string
		expectedSteps int
	}

	testCases := []testCase{
		{`LR

		11A = (11B, XXX)
		11B = (XXX, 11Z)
		11Z = (11B, XXX)
		22A = (22B, XXX)
		22B = (22C, 22C)
		22C = (22Z, 22Z)
		22Z = (22B, 22B)
		XXX = (XXX, XXX)`, 6},
	}

	for _, test := range testCases {
		lines := strings.Split(test.content, "\n")

		directions := ParseDirections(lines[0])

		if lines[1] != "" {
			log.Panicf("unexpected non blank line in input: '%s'\n", lines[1])
		}

		n := ParseNetwork(lines[2:])

		nodesEndingInA := make([]*Node, 0)

		n.ForEach(func(node *Node) bool {
			if strings.HasSuffix(node.Name, "A") {
				nodesEndingInA = append(nodesEndingInA, node)
			}

			return true
		})

		steps := n.GhostWalk(nodesEndingInA, directions, func(nodes []*Node) bool {
			for _, n := range nodes {
				if !strings.HasSuffix(n.Name, "Z") {
					return false
				}
			}

			return true
		})

		assert.Equal(t, test.expectedSteps, steps)
	}
}
