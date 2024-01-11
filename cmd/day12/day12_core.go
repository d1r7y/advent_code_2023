package day12

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type SpringState byte

const (
	Operational SpringState = iota
	Broken
	Unknown
)

func (ss SpringState) Describe() string {
	switch ss {
	case Operational:
		return "."
	case Broken:
		return "#"
	case Unknown:
		return "?"
	}

	log.Panicf("unknown spring state: %d\n", ss)
	return "#"
}

type RejectState byte

const (
	NoReject RejectState = iota
	RejectOperationalState
	RejectBrokenState
)

type RejectStateList []RejectState

type SpringStateList []SpringState

func (ssl SpringStateList) Describe() string {
	str := ""

	for _, s := range ssl {
		str += s.Describe()
	}

	return str
}

type SpringGroup struct {
	Unfolded          int
	States            SpringStateList
	DamagedSpringRuns []int
}

func NewSpringGroup(unfolded int, states SpringStateList, damagedSpringRun []int) *SpringGroup {
	group := &SpringGroup{}

	group.Unfolded = unfolded

	group.States = make(SpringStateList, len(states))
	copy(group.States, states)

	group.DamagedSpringRuns = make([]int, len(damagedSpringRun))
	copy(group.DamagedSpringRuns, damagedSpringRun)

	return group
}

func (sg *SpringGroup) Describe() string {
	str := ""
	for _, state := range sg.States {
		str += state.Describe()
	}

	str += " "

	damagedRunStringsList := make([]string, 0)

	for _, dr := range sg.DamagedSpringRuns {
		damagedRunStringsList = append(damagedRunStringsList, fmt.Sprintf("%d", dr))
	}

	str += strings.Join(damagedRunStringsList, ",")

	return str
}

func (sg *SpringGroup) Solve() []*SpringGroup {
	solutions := make([]*SpringGroup, 0)

	// Get all the solutions given states
	for _, alternative := range GenerateAlternatives(sg.States, sg.DamagedSpringRuns, sg.Unfolded) {
		if !IsAlternativeValid(alternative, sg.DamagedSpringRuns) {
			continue
		}

		// Yes!  Add it to our list.
		group := NewSpringGroup(1, alternative, sg.DamagedSpringRuns)
		solutions = append(solutions, group)
	}

	return solutions
}

func IsAlternativeValid(states SpringStateList, requirements []int) bool {
	simplified := SimplifySpringStateRun(states, Broken)

	// Now see if this alternative matches our damaged spring run requirements.
	if len(simplified) != len(requirements) {
		return false
	}

	for i := 0; i < len(requirements); i++ {
		if simplified[i] != requirements[i] {
			return false
		}
	}

	return true
}

func (sg *SpringGroup) StateCount(InterestedState SpringState) int {
	count := 0

	for _, state := range sg.States {
		if state == InterestedState {
			count++
		}
	}

	return count
}

func (sg *SpringGroup) Unfold(factor int) *SpringGroup {
	unfoldedStates := make(SpringStateList, 0)

	for i := 0; i < factor; i++ {
		unfoldedStates = append(unfoldedStates, sg.States...)

		if i < factor-1 {
			unfoldedStates = append(unfoldedStates, Unknown)
		}
	}

	unfoldedDamagedSpringRuns := make([]int, 0)

	for i := 0; i < factor; i++ {
		unfoldedDamagedSpringRuns = append(unfoldedDamagedSpringRuns, sg.DamagedSpringRuns...)
	}

	return NewSpringGroup(factor, unfoldedStates, unfoldedDamagedSpringRuns)
}

func SimplifySpringStateRun(states SpringStateList, interestedState SpringState) []int {
	simplified := make([]int, 0)

	runCount := 0

	for _, s := range states {
		if s == Unknown {
			break
		}

		if s != interestedState {
			if runCount > 0 {
				simplified = append(simplified, runCount)
			}
			runCount = 0
		} else {
			runCount++
		}
	}

	if runCount > 0 {
		simplified = append(simplified, runCount)
	}

	return simplified
}

func generateAlternativesCore(alternatives []SpringStateList, states SpringStateList, offset int, requirements []int, unfold int, reject RejectStateList) []SpringStateList {

	createAlternate := func(states SpringStateList, offset int, newState SpringState) SpringStateList {
		alternateState := make(SpringStateList, len(states))
		copy(alternateState, states)
		alternateState[offset] = newState

		return alternateState
	}

	invalidAlternate := func(states SpringStateList, requirements []int) bool {
		simplified := SimplifySpringStateRun(states, Broken)

		// len(simplified) <= len(requirements) is possibly valid.
		// len(simplified) > len(requirements) is definitely NOT valid.
		if len(simplified) > len(requirements) {
			return true
		}

		// Except for the last one, each run in simplified must be equal to the
		// corresponding run in requirements.  The last one can be less as there
		// might be unmutated unknowns.
		for i := 0; i < len(simplified); i++ {
			sr := simplified[i]
			rr := requirements[i]

			if i < len(simplified)-1 {
				// Not last run, so they must be equal.
				if sr != rr {
					return true
				}
			} else {
				// Last run, sr can be less than or equal to rr.
				if sr > rr {
					return true
				}
			}
		}

		return false
	}

	updateRejectList := func(rejectState RejectState, reject RejectStateList, offset int, unfold int) {
		// nextOffset := ((len(reject) - (unfold - 1)) / unfold) + 1
		// for i := offset + nextOffset; i < len(reject); i += nextOffset {
		// 	reject[i] |= rejectState
		// }
	}

	if offset == len(states) {
		// We've reached the end.  No more states to mutate.
		alternatives = append(alternatives, states)

		return alternatives
	}

	if states[offset] == Unknown {
		if (reject[offset] & RejectBrokenState) == 0 {
			// Generate Broken alternative.
			mutateBroken := createAlternate(states, offset, Broken)

			// Before going down this path, sanity check that mutateBroken is valid.  If it
			// isn't, then any further mutations won't be valid.
			if invalidAlternate(mutateBroken, requirements) {
				updateRejectList(RejectBrokenState, reject, offset, unfold)
			} else {
				alternatives = generateAlternativesCore(alternatives, mutateBroken, offset+1, requirements, unfold, reject)
			}
		}

		if (reject[offset] & RejectOperationalState) == 0 {
			// Generate Operational alternative.
			mutateOperational := createAlternate(states, offset, Operational)

			// Before going down this path, sanity check that mutateOperational is valid.  If it
			// isn't, then any further mutations won't be valid.
			if invalidAlternate(mutateOperational, requirements) {
				updateRejectList(RejectOperationalState, reject, offset, unfold)
			} else {
				alternatives = generateAlternativesCore(alternatives, mutateOperational, offset+1, requirements, unfold, reject)
			}
		}

	} else {
		// Nothing to mutate at this state, pass the remainder on.
		alternatives = generateAlternativesCore(alternatives, states, offset+1, requirements, unfold, reject)
	}

	return alternatives
}

func GenerateAlternatives(states SpringStateList, requirements []int, unfold int) []SpringStateList {
	alternatives := make([]SpringStateList, 0)
	reject := make(RejectStateList, len(states))
	for i := 0; i < len(reject); i++ {
		reject[i] = NoReject
	}

	return generateAlternativesCore(alternatives, states, 0, requirements, unfold, reject)
}

func ParseLine(line string) *SpringGroup {
	conditionAndList := strings.Split(line, " ")

	if len(conditionAndList) != 2 {
		log.Panicf("unexpected line '%s'\n", line)
	}

	states := make(SpringStateList, 0)
	damagedSpringRuns := make([]int, 0)

	for _, s := range conditionAndList[0] {
		switch s {
		case '.':
			states = append(states, Operational)
		case '#':
			states = append(states, Broken)
		case '?':
			states = append(states, Unknown)
		default:
			log.Panicf("unexpected spring state: %d\n", s)
		}
	}

	for _, n := range strings.Split(conditionAndList[1], ",") {
		number, err := strconv.Atoi(n)
		if err != nil {
			log.Panic(err)
		}

		damagedSpringRuns = append(damagedSpringRuns, number)
	}

	group := NewSpringGroup(1, states, damagedSpringRuns)
	return group
}

func day12(fileContents string) error {
	// Part 1: For each row, count all of the different arrangements of operational and broken
	// springs that meet the given criteria. What is the sum of those counts?

	totalArrangements := 0

	for _, line := range strings.Split(strings.TrimSpace(fileContents), "\n") {
		springGroup := ParseLine(line)

		totalArrangements += len(springGroup.Solve())
	}

	log.Printf("Sum of possible arrangements: %d\n", totalArrangements)

	// Part 2: Unfold your condition records; what is the new sum of possible arrangement counts?

	totalUnfoldedArrangements := 0

	for _, line := range strings.Split(strings.TrimSpace(fileContents), "\n") {
		springGroup := ParseLine(line)
		unfolded := springGroup.Unfold(5)
		log.Println(unfolded.Describe())

		totalUnfoldedArrangements += len(unfolded.Solve())
	}

	log.Printf("Sum of unfolded possible arrangements: %d\n", totalUnfoldedArrangements)

	return nil
}
