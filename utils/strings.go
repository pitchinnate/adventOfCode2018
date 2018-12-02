package utils

import (
	"log"
	"strconv"
)

func ArrayOfStringsToInts(strings []string) []int {
	var ints []int
	for _,str := range strings {
		newInt, err := strconv.Atoi(str)
		if err != nil {
			log.Fatal(err)
		}
		ints = append(ints, newInt)
	}
	return ints
}
