package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day5/day5.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	afters := map[int][]int{}
	var updates [][]int
	startPages := false
	for scanner.Scan() {
		line := scanner.Text()
		if startPages {
			pageNums := strings.Split(line, ",")
			var update []int
			for _, page := range pageNums {
				pageNum, err := strconv.Atoi(page)
				if err != nil {
					panic(err)
				}
				update = append(update, pageNum)
			}
			updates = append(updates, update)
		} else if line == "" {
			startPages = true
		} else {
			pageComps := strings.Split(line, "|")
			before, err := strconv.Atoi(pageComps[0])
			if err != nil {
				panic(err)
			}
			after, err := strconv.Atoi(pageComps[1])
			if err != nil {
				panic(err)
			}
			afterGroup, found := afters[after]
			if !found {
				afterGroup = []int{}
			}
			afterGroup = append(afterGroup, before)
			afters[after] = afterGroup
		}
	}

	midSum := 0
	for _, update := range updates {
		midSum += getUpdateMid(update, afters, false)
	}
	fmt.Println(midSum)

	midSum = 0
	for _, update := range updates {
		midSum += getUpdateMid(update, afters, true)
	}
	fmt.Println(midSum)
}

func getUpdateMid(update []int, afters map[int][]int, fix bool) int {
	valid := true
	for idx := range update {
		if idx == len(update)-1 {
			break
		}
		for idxAfter := range update[idx+1:] {
			idxAfter += idx + 1
			group, found := afters[update[idx]]
			if !found {
				continue
			} else if slices.Contains(group, update[idxAfter]) {
				if fix {
					tmp := update[idx]
					update[idx] = update[idxAfter]
					update[idxAfter] = tmp
					valid = false
				} else {
					return 0
				}
			}
		}
	}
	if !fix && valid || fix && !valid {
		return update[len(update)/2]
	}
	return 0
}
