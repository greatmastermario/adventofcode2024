package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("day4/day4.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	var data [][]byte
	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, []byte(line))
	}
	fmt.Println(findXmas(data))
	fmt.Println(findXShapedMas(data))
}

func findXmas(data [][]byte) int {
	cnt := 0
	for row, line := range data {
		for col, char := range line {
			if char == 'X' {
				if row > 2 {
					if col > 2 && isXmas(data[row-1][col-1], data[row-2][col-2], data[row-3][col-3]) {
						cnt++
					}
					if isXmas(data[row-1][col], data[row-2][col], data[row-3][col]) {
						cnt++
					}
					if col < len(line)-3 && isXmas(data[row-1][col+1], data[row-2][col+2], data[row-3][col+3]) {
						cnt++
					}
				}
				if col > 2 && isXmas(line[col-1], line[col-2], line[col-3]) {
					cnt++
				}
				if col < len(line)-3 && isXmas(line[col+1], line[col+2], line[col+3]) {
					cnt++
				}
				if row < len(data)-3 {
					if col > 2 && isXmas(data[row+1][col-1], data[row+2][col-2], data[row+3][col-3]) {
						cnt++
					}
					if isXmas(data[row+1][col], data[row+2][col], data[row+3][col]) {
						cnt++
					}
					if col < len(line)-3 && isXmas(data[row+1][col+1], data[row+2][col+2], data[row+3][col+3]) {
						cnt++
					}
				}
			}
		}
	}
	return cnt
}

func isXmas(m byte, a byte, s byte) bool {
	return m == 'M' && a == 'A' && s == 'S'
}

func findXShapedMas(data [][]byte) int {
	cnt := 0
	for row, line := range data {
		if row == 0 {
			continue
		}
		if row == len(data)-1 {
			break
		}
		for col, char := range line {
			if col == 0 {
				continue
			}
			if col == len(line)-1 {
				break
			}
			if char == 'A' && isXShapedMas(data[row-1][col-1], data[row-1][col+1], data[row+1][col-1], data[row+1][col+1]) {
				cnt++
			}
		}
	}
	return cnt
}

func isXShapedMas(topleft byte, topright byte, botleft byte, botright byte) bool {
	if topleft == 'A' || topleft == 'X' || topright == 'A' || topright == 'X' ||
		botleft == 'A' || botleft == 'X' || botright == 'A' || botright == 'X' ||
		topleft == botright || topright == botleft {
		return false
	}
	return (topleft == botleft && topright == botright) || (topleft == topright && botleft == botright)
}
