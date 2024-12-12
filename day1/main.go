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
	file, err := os.Open("day1/day1.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	data := [2][]int{{}, {}}
	totalDistance := 0
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "   ")
		split0, err := strconv.Atoi(split[0])
		if err != nil {
			panic(err)
		}
		data[0] = append(data[0], split0)
		split1, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}
		data[1] = append(data[1], split1)
	}

	for _, v := range data {
		slices.Sort(v)
	}

	for idx := 0; idx < len(data[0]); idx++ {
		distance := data[1][idx] - data[0][idx]
		if distance < 0 {
			distance = -distance
		}
		totalDistance += distance
	}
	fmt.Println(totalDistance)
}
