package day15

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	type testCase struct {
		str          string
		expectedHash byte
	}
	testCases := []testCase{
		{str: "HASH", expectedHash: 52},
		{str: "rn=1", expectedHash: 30},
		{str: "cm-", expectedHash: 253},
		{str: "qp=3", expectedHash: 97},
		{str: "cm=2", expectedHash: 47},
		{str: "qp-", expectedHash: 14},
		{str: "pc=4", expectedHash: 180},
		{str: "ot=9", expectedHash: 9},
		{str: "ab=5", expectedHash: 197},
		{str: "pc-", expectedHash: 48},
		{str: "pc=6", expectedHash: 214},
		{str: "ot=7", expectedHash: 231},
		{str: "rn", expectedHash: 0},
		{str: "cm", expectedHash: 0},
		{str: "qp", expectedHash: 1},
		{str: "pc", expectedHash: 3},
	}

	for _, test := range testCases {
		assert.Equal(t, test.expectedHash, Hash(test.str))
	}
}

func TestSumInitializationSequence(t *testing.T) {
	assert.Equal(t, 1320, SumInitializationSequence("rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7"))
}

func TestBoxFocusingPower(t *testing.T) {
	type testCase struct {
		box                   Box
		expectedFocusingPower int
	}
	testCases := []testCase{
		{box: Box{Number: 0, Lenses: []Lens{{Name: "rn", FocalLength: 1}, {Name: "cm", FocalLength: 2}}}, expectedFocusingPower: 5},
		{box: Box{Number: 3, Lenses: []Lens{{Name: "ot", FocalLength: 7}, {Name: "ab", FocalLength: 5}, {Name: "pc", FocalLength: 6}}}, expectedFocusingPower: 140},
	}

	for _, test := range testCases {
		assert.Equal(t, test.expectedFocusingPower, test.box.FocusingPower())
	}
}

func TestBoxLineSetLens(t *testing.T) {
	boxLine := NewBoxLine()
	expectedBoxLine := NewBoxLine()

	boxLine.SetLens("rn", 1)
	expectedBoxLine[0].Lenses = []Lens{{Name: "rn", FocalLength: 1}}
	assert.Equal(t, expectedBoxLine, boxLine)

	boxLine.SetLens("qp", 3)
	expectedBoxLine[1].Lenses = []Lens{{Name: "qp", FocalLength: 3}}
	assert.Equal(t, expectedBoxLine, boxLine)

	boxLine.SetLens("cm", 2)
	expectedBoxLine[0].Lenses = append(expectedBoxLine[0].Lenses, Lens{Name: "cm", FocalLength: 2})
	assert.Equal(t, expectedBoxLine, boxLine)
}

func TestBoxLineRemoveLens(t *testing.T) {
	boxLine := NewBoxLine()
	expectedBoxLine := NewBoxLine()

	boxLine.SetLens("rn", 1)
	expectedBoxLine[0].Lenses = []Lens{{Name: "rn", FocalLength: 1}}
	assert.Equal(t, expectedBoxLine, boxLine)

	boxLine.RemoveLens("cm")
	assert.Equal(t, expectedBoxLine, boxLine)

	boxLine.SetLens("qp", 3)
	expectedBoxLine[1].Lenses = []Lens{{Name: "qp", FocalLength: 3}}
	assert.Equal(t, expectedBoxLine, boxLine)

	boxLine.RemoveLens("qp")
	expectedBoxLine[1].Lenses = make([]Lens, 0)
	assert.Equal(t, expectedBoxLine, boxLine)
}

func TestSumFocusingPowerFromInitializationSequence(t *testing.T) {
	assert.Equal(t, 145, SumFocusingPowerFromInitializationSequence("rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7"))
}
