package day08

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

type Node struct {
	Name  string
	Left  *Node
	Right *Node
}

func (n *Node) Describe() string {
	return fmt.Sprintf("%s = (%s, %s)", n.Name, n.Left.Name, n.Right.Name)
}

type NodeDescription struct {
	Name      string
	LeftName  string
	RightName string
}

type Network struct {
	NameLookup map[string]*Node
}

func NewNetwork() *Network {
	return &Network{NameLookup: make(map[string]*Node)}
}

func (n *Network) ForEach(yield func(node *Node) bool) {
	for _, v := range n.NameLookup {
		if !yield(v) {
			break
		}
	}
}

func (n *Network) Find(name string) *Node {
	return n.NameLookup[name]
}

func (n *Network) Walk(start *Node, directions []Direction, end *Node) int {
	currentNode := start

	steps := 0

	for {
		for _, direction := range directions {
			if currentNode.Name == end.Name {
				return steps
			}

			steps++

			if direction == Left {
				currentNode = currentNode.Left
			} else {
				currentNode = currentNode.Right
			}
		}
	}
}

func (n *Network) GhostWalk(start []*Node, directions []Direction, shouldEnd func(nodes []*Node) bool) int {
	currentNodes := make([]*Node, len(start))
	copy(currentNodes, start)

	steps := 0

	for {
		for _, direction := range directions {
			if shouldEnd(currentNodes) {
				return steps
			}

			steps++

			for i := range currentNodes {
				if direction == Left {
					currentNodes[i] = currentNodes[i].Left
				} else {
					currentNodes[i] = currentNodes[i].Right
				}
			}
		}
	}
}

func ParseNetwork(lines []string) *Network {
	network := NewNetwork()

	nodeDescriptions := make([]NodeDescription, 0)

	for _, line := range lines {
		nodeDescriptions = append(nodeDescriptions, ParseNodeDescription(line))
	}

	// Go through the descriptons twice: first to gather all the nodes, then to connect them up.
	for _, nd := range nodeDescriptions {
		node := &Node{Name: nd.Name}
		network.NameLookup[nd.Name] = node
	}

	for _, nd := range nodeDescriptions {
		node := network.NameLookup[nd.Name]
		node.Left = network.NameLookup[nd.LeftName]
		node.Right = network.NameLookup[nd.RightName]
	}

	return network
}

type Direction byte

const (
	Left Direction = iota
	Right
)

func ParseDirections(line string) []Direction {
	directionsRE := regexp.MustCompile(`L|R`)
	directionsMatches := directionsRE.FindAllString(line, -1)

	directions := make([]Direction, 0)

	for _, d := range directionsMatches {
		var direction Direction

		switch d {
		case "L":
			direction = Left
		case "R":
			direction = Right
		default:
			log.Panicf("unknown direction '%s'\n", d)
		}

		directions = append(directions, direction)
	}

	return directions
}

func ParseNodeDescription(line string) NodeDescription {
	nodeNameRE := regexp.MustCompile(`[0-9A-Z]{3}`)
	nodeNameMatches := nodeNameRE.FindAllString(line, -1)

	if len(nodeNameMatches) != 3 {
		log.Panicf("unexpected node line '%s'\n", line)
	}

	return NodeDescription{
		Name:      nodeNameMatches[0],
		LeftName:  nodeNameMatches[1],
		RightName: nodeNameMatches[2],
	}
}

func PrimeFactors(n int) (pfs []int) {
	// Get the number of 2s that divide n
	for n%2 == 0 {
		pfs = append(pfs, 2)
		n = n / 2
	}

	// n must be odd at this point. so we can skip one element
	// (note i = i + 2)
	for i := 3; i*i <= n; i = i + 2 {
		// while i divides n, append i and divide n
		for n%i == 0 {
			pfs = append(pfs, i)
			n = n / i
		}
	}

	// This condition is to handle the case when n is a prime number
	// greater than 2
	if n > 2 {
		pfs = append(pfs, n)
	}

	return
}

func day08(fileContents string) error {
	lines := strings.Split(string(fileContents), "\n")

	directions := ParseDirections(lines[0])

	if lines[1] != "" {
		log.Panicf("unexpected non blank line in input: '%s'\n", lines[1])
	}

	n := ParseNetwork(lines[2:])

	// Part 1: Starting at AAA, follow the left/right instructions. How many steps are required to reach ZZZ?
	steps := n.Walk(n.Find("AAA"), directions, n.Find("ZZZ"))

	log.Printf("Total steps: %d\n", steps)

	// Part 2: Simultaneously start on every node that ends with A. How many steps does it take before you're
	// only on nodes that end with Z?
	// Really what we should be doing is finding all the unique factors for all the steps and multiply them together.

	nodesEndingInA := make([]*Node, 0)

	n.ForEach(func(node *Node) bool {
		if strings.HasSuffix(node.Name, "A") {
			nodesEndingInA = append(nodesEndingInA, node)
		}

		return true
	})

	log.Println("Nodes ending in A")
	for _, n := range nodesEndingInA {
		fmt.Println(n.Describe())
	}

	allGhostSteps := make([]int, 0)

	for _, node := range nodesEndingInA {
		ghostSteps := n.GhostWalk([]*Node{node}, directions, func(nodes []*Node) bool {
			for _, n := range nodes {
				if !strings.HasSuffix(n.Name, "Z") {
					return false
				}
			}

			return true
		})

		allGhostSteps = append(allGhostSteps, ghostSteps)
	}

	uniqueFactors := make(map[int]bool)

	for _, n := range allGhostSteps {

		log.Printf("Prime factors for %d: %v\n", n, PrimeFactors(n))
		for _, p := range PrimeFactors(n) {
			uniqueFactors[p] = true
		}
	}

	ghostSteps := 1

	for k := range uniqueFactors {
		ghostSteps *= k
	}

	log.Printf("Total ghost steps %d\n", ghostSteps)

	return nil
}
