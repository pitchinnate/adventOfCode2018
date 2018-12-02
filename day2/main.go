package day2

import (
	"advent2018/utils"
	"log"
	"strings"
)

func getInput() []string {
	lines := utils.ReadFile("../day2/day2_input.txt")
	return lines
}

func Day2() {
	countTwo := 0
	countThree := 0
	inputs := getInput()

	for _, line := range inputs {
		counts := map[string]int{}
		letters := strings.Split(line, "")
		for _, letter := range letters {
			counts[letter] += 1
		}
		for _, count := range counts {
			if count == 2 {
				countTwo += 1
				break
			}
		}
		for _, count := range counts {
			if count == 3 {
				countThree += 1
				break
			}
		}
	}
	log.Printf("Twos: %d Threes: %d checksum: %d", countTwo, countThree, (countTwo * countThree))
}

type combos struct {
	StringA string
	StringB string
	IndexA int
	IndexB int
	Shared int
}

func Day2Part2() {
	inputs := getInput()
	var allCombos []combos

	for x := 0; x < (len(inputs) - 1); x++ {
		for y := (x+1); y < len(inputs); y++ {
			allCombos = append(allCombos, combos{
				inputs[x],
				inputs[y],
				x,
				y,
				0,
			})
		}
	}

	var best combos
	maxShared := 0
	for _, combo := range allCombos {
		sameLetter := 0
		lettersA := strings.Split(combo.StringA, "")
		lettersB := strings.Split(combo.StringB, "")
		for i, letter := range lettersA {
			if letter == lettersB[i] {
				sameLetter += 1
			}
		}
		combo.Shared = sameLetter
		if combo.Shared > maxShared {
			maxShared = combo.Shared
			best = combo
		}
	}

	log.Print(best)
}

