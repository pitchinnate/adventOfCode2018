package day1

import (
	"advent2018/utils"
	"log"
)

func getInput() []int {
	strings := utils.ReadFile("../day1/day1_input.txt")
	ints := utils.ArrayOfStringsToInts(strings)
	return ints
}

func Day1() {
	currentFreq := 0
	inputs := getInput()
	for _, input := range inputs {
		currentFreq += input
	}
	log.Printf("Final Freq: %d", currentFreq)
}

func Day1Part2() {
	var usedFreqs []int
	currentFreq := 0
	usedFreqs = append(usedFreqs, 0)
	inputs := getInput()
	found := false
	for !found {
		found = part2(&usedFreqs, &currentFreq, inputs)
	}
}

func part2(usedFreqs *[]int, currentFreq *int, inputs []int) bool {
	log.Print("Running through a cycle")
	for _, input := range inputs {
		*currentFreq = *currentFreq + input
		searchIndex := utils.SearchIntsForInt(*usedFreqs, *currentFreq)
		if searchIndex > -1 {
			log.Print("Found this frequency twice ", *currentFreq)
			return true
		}
		*usedFreqs = append(*usedFreqs, *currentFreq)
	}
	return false
}
