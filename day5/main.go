package day5

import (
	"advent2018/utils"
	"fmt"
	"log"
	"strconv"
)

func getInput() string {
	lines := utils.ReadFile("../day5/day5_input.txt")
	log.Print(len(lines))
	return lines[0]
}

func Day5() {
	input := getInput()
	//input := "dabAcCaCBAcCcaDA"
	asciis := []int{}
	for _, letter := range input {
		intval, _ := strconv.Atoi(fmt.Sprintf("%d", letter))
		if (intval >= 65 && intval <= 90) || (intval >= 97 && intval <= 122) {
			asciis = append(asciis, intval)
		} else {
			log.Print("ascii not accepted", intval)
		}
	}
	// A-Z -> 65-90
	// a-z -> 97-122
	keepRunning := true
	for keepRunning {
		keepAsciis := []int{}
		changeCounter := 0
		for x := 0; x < (len(asciis) - 1); x++ {
			var lookFor int
			if asciis[x] <= 90 {
				lookFor = asciis[x] + 32
			} else {
				lookFor = asciis[x] - 32
			}
			if lookFor == asciis[(x+1)] {
				//log.Printf("remove %d %d", asciis[x], asciis[(x+1)])
				changeCounter++
				x++
				if x == (len(asciis) - 2) {
					keepAsciis = append(keepAsciis, asciis[(x + 1)])
				}
				continue
			}

			keepAsciis = append(keepAsciis, asciis[x])
			if x == (len(asciis) - 2) {
				keepAsciis = append(keepAsciis, asciis[(x + 1)])
			}
		}
		//log.Print("Changes:", changeCounter)
		asciis = keepAsciis
		if changeCounter == 0 {
			keepRunning = false
		}
	}
	log.Print(len(asciis))
}

func Day5Part2() {
	input := getInput()
	//input := "dabAcCaCBAcCcaDA"
	asciis := []int{}
	for _, letter := range input {
		intval, _ := strconv.Atoi(fmt.Sprintf("%d", letter))
		if (intval >= 65 && intval <= 90) || (intval >= 97 && intval <= 122) {
			asciis = append(asciis, intval)
		} else {
			log.Print("ascii not accepted", intval)
		}
	}
	// A-Z -> 65-90
	// a-z -> 97-122
	for y := 65; y <= 90; y++ {
		filteredAsciis := []int{}
		for _, val := range asciis {
			if val != y && val != (y + 32) {
				filteredAsciis = append(filteredAsciis, val)
			}
		}

		keepRunning := true
		for keepRunning {
			keepAsciis := []int{}
			changeCounter := 0
			for x := 0; x < (len(filteredAsciis) - 1); x++ {
				var lookFor int
				if filteredAsciis[x] <= 90 {
					lookFor = filteredAsciis[x] + 32
				} else {
					lookFor = filteredAsciis[x] - 32
				}
				if lookFor == filteredAsciis[(x + 1)] {
					//log.Printf("remove %d %d", asciis[x], asciis[(x+1)])
					changeCounter++
					x++
					if x == (len(filteredAsciis) - 2) {
						keepAsciis = append(keepAsciis, filteredAsciis[(x + 1)])
					}
					continue
				}

				keepAsciis = append(keepAsciis, filteredAsciis[x])
				if x == (len(filteredAsciis) - 2) {
					keepAsciis = append(keepAsciis, filteredAsciis[(x + 1)])
				}
			}
			//log.Print("Changes:", changeCounter)
			filteredAsciis = keepAsciis
			if changeCounter == 0 {
				keepRunning = false
			}
		}
		log.Print(len(filteredAsciis))
	}
}

