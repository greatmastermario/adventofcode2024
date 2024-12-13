package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	file, err := os.Open("day3/day3.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	data := ""
	for scanner.Scan() {
		data += scanner.Text()
	}
	//Part 1
	total := parseMuls(data, false)
	fmt.Println(total)
	//Part 2
	total = parseMuls(data, true)
	fmt.Println(total)
}

func parseMuls(data string, do bool) int {
	result := 0
	dataBytes := []byte(data)
	skip := 0
	for idx := bytes.Index(dataBytes, []byte("mul(")); idx > -1; idx = bytes.Index(dataBytes, []byte("mul(")) {
		idx += 4
		var args []byte
		for ; idx < len(dataBytes) && (unicode.IsDigit(rune(dataBytes[idx])) || ',' == dataBytes[idx]); idx++ {
			args = append(args, dataBytes[idx])
		}
		if idx < len(dataBytes) && ')' == dataBytes[idx] {
			argString := string(args)
			splitArgs := strings.Split(argString, ",")
			arg1, err := strconv.Atoi(splitArgs[0])
			if err != nil {
				panic(err)
			}
			arg2, err := strconv.Atoi(splitArgs[1])
			if err != nil {
				panic(err)
			}
			if do {
				doIdx := strings.LastIndex(data[:idx+skip], "do()")
				dontIdx := strings.LastIndex(data[:idx+skip], "don't()")
				if doIdx > dontIdx || (dontIdx == -1 && doIdx == -1) {
					result += arg1 * arg2
				}
			} else {
				result += arg1 * arg2
			}
		}
		if idx < len(dataBytes) {
			dataBytes = dataBytes[idx+1:]
			skip += idx + 1
		} else {
			dataBytes = []byte{}
		}
	}
	return result
}
