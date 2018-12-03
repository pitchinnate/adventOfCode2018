package day3

import (
	"advent2018/utils"
	"log"
	"strconv"
	"strings"
)

type blockCloth struct {
	Id int
	Left int
	Top int
	Width int
	Height int
	Overlaps bool
}

func getInput() []blockCloth {
	var blocks []blockCloth
	lines := utils.ReadFile("../day3/day3_input.txt")
	for _, line := range lines {
		pieces := strings.Split(line, " ")
		idString := pieces[0][1:len(pieces[0])]
		id, _ := strconv.Atoi(idString)

		leftRight := pieces[2][:(len(pieces[2]) - 1)]
		cords := strings.Split(leftRight, ",")
		left, _ := strconv.Atoi(cords[0])
		top, _ := strconv.Atoi(cords[1])

		square := strings.Split(pieces[3], "x")
		width, _ := strconv.Atoi(square[0])
		height, _ := strconv.Atoi(square[1])

		blocks = append(blocks, blockCloth{
			id,
			left,
			top,
			width,
			height,
			false,
		})
	}
	return blocks
}

func Day3() {
	blocks := getInput()
	grid := map[int]map[int]int{}

	maxX := 0
	maxY := 0

	for _, block := range blocks {
		for y := block.Top; y < (block.Height + block.Top); y++ {
			if grid[y] == nil {
				grid[y] = map[int]int{}
			}
			for x := block.Left; x < (block.Width + block.Left); x++ {
				grid[y][x] += 1
				if x > maxX {
					maxX = x
				}
			}
			if y > maxY {
				maxY = y
			}
		}
	}

	countRedudent := 0
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if grid[y][x] >= 2 {
				countRedudent++
			}
		}
	}
	log.Print("Redudent areas ", countRedudent)
}

func Day3Part2() {
	blocks := getInput()
	grid := map[int]map[int][]int{}

	maxX := 0
	maxY := 0

	for _, block := range blocks {
		for y := block.Top; y < (block.Height + block.Top); y++ {
			if grid[y] == nil {
				grid[y] = map[int][]int{}
			}
			for x := block.Left; x < (block.Width + block.Left); x++ {
				grid[y][x] = append(grid[y][x], block.Id)
				if x > maxX {
					maxX = x
				}
			}
			if y > maxY {
				maxY = y
			}
		}
	}

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if len(grid[y][x]) >= 2 {
				for _, id := range grid[y][x] {
					blocks[(id - 1)].Overlaps = true
				}
			}
		}
	}

	for _, block := range blocks {
		if !block.Overlaps {
			log.Print("Block doesn't overlap ", block)
		}
	}
}
