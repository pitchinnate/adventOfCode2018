package day4

import (
	"advent2018/utils"
	"log"
	"sort"
	"strconv"
	"strings"
)

type guardAction struct {
	Type string
	Minute  int
	GuardId string
}

func getInput() map[string][]guardAction {
	allKeys := []string{}
	allActions := map[string]guardAction{}
	dateActions := map[string][]guardAction{}
	lines := utils.ReadFile("../day4/day4_input.txt")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		date := line[1:17]
		time, _ := strconv.Atoi(line[15:17])
		actionType := line[19:24]
		gId := ""
		if actionType == "Guard" {
			pieces := strings.Split(line, " ")
			gId = pieces[3]
		}
		allKeys = append(allKeys, date)
		allActions[date] = guardAction{
			actionType,
			time,
			gId,
		}
	}
	sort.Strings(allKeys)

	guardId := ""
	for _, key := range allKeys {
		if allActions[key].Type == "Guard" {
			guardId = allActions[key].GuardId
			continue
		}
		date := key[0:11]
		action := allActions[key]
		action.GuardId = guardId
		dateActions[date] = append(dateActions[date], action)
	}

	return dateActions
}

type guardDay struct {
	Date string
	GuardId string
	MinutesAsleep []int
}

func Day4() {
	days := []guardDay{}
	sleepers := getInput()
	guardSleepCount := map[string]int{}
	for date, actions := range sleepers {
		minuteActions := map[int]string{}
		for _, action := range actions {
			minuteActions[action.Minute] = action.Type
		}

		sleepingMinutes := []int{}
		sleeping := false
		for x := 0; x <= 59; x++ {
			if minuteActions[x] == "falls" {
				sleeping = true
			}
			if minuteActions[x] == "wakes" {
				sleeping = false
			}
			if sleeping {
				sleepingMinutes = append(sleepingMinutes, x)
				guardSleepCount[actions[0].GuardId] += 1
			}
		}

		days = append(days, guardDay{
			date,
			actions[0].GuardId,
			sleepingMinutes,
		})
	}

	guardSleepMinutesCount := map[string]map[int]int{}
	bestGuardId := ""
	bestMinute := 0
	bestMinuteCount := 0

	for _, day := range days {
		// this is for part 1 had to find guard that slept the most and limit it to them
		//if day.GuardId != "#1049" {
		//	continue
		//}
		if guardSleepMinutesCount[day.GuardId] == nil {
			guardSleepMinutesCount[day.GuardId] = map[int]int{}
		}
		for _, minute := range day.MinutesAsleep {
			guardSleepMinutesCount[day.GuardId][minute] += 1
			if guardSleepMinutesCount[day.GuardId][minute] > bestMinuteCount {
				bestMinuteCount = guardSleepMinutesCount[day.GuardId][minute]
				bestMinute = minute
				bestGuardId = day.GuardId
			}
		}
	}

	log.Print("Best guard ", bestGuardId, " best minute: ", bestMinute, " was sleeping ", bestMinuteCount)
}

