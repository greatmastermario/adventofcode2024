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
	file, err := os.Open("day2/day2.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	data := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		report := []int{}
		for _, val := range split {
			reading, err := strconv.Atoi(val)
			if err != nil {
				panic(err)
			}
			report = append(report, reading)
		}
		data = append(data, report)
	}

	part1(data, true)  //Part 1 Skip dampening
	part1(data, false) //Part 2 include dampening
}

func part1(data [][]int, dampenedDefault bool) {
	good := 0
	for _, report := range data {
		if runReport(report, dampenedDefault) {
			good++
		} else if !dampenedDefault {
			slices.Reverse(report)
			if runReport(report, dampenedDefault) {
				good++
			}
		}
	}
	fmt.Println(good)
}

func runReport(report []int, dampened bool) bool {
	lastDiff := 0
	lastReading := 0
	for idx, reading := range report {
		if idx != 0 {
			diff := reading - lastReading
			if diff > 0 && lastDiff < 0 || diff < 0 && lastDiff > 0 || diff == 0 || diff > 3 || diff < -3 {
				if dampened {
					break
				}
				dampened = true
				if idx == len(report)-1 {
					return true
				}
				continue
			}
			lastDiff = diff
		}
		lastReading = reading
		if idx == len(report)-1 {
			return true
		}
	}
	return false
}
