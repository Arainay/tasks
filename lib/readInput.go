package lib

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadInput() []string {
	limit := 2
	lines := []string{}
	line := ""
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		text := string(scanner.Text())
		isNewLine := text == "\n"

		if !isNewLine {
			if text != "\r" {
				line += text
			}
		} else {
			if len(lines) == 0 {
				limit = StringToInteger(line)*2 + 1
			}
			lines = append(lines, line)
			line = ""
		}

		if len(lines) >= limit {
			break
		}
	}

	return lines
}

func StringToInteger(str string) int {
	intValue, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}

	return intValue
}

func StringToIntegers(str string) []int {
	digits := strings.Split(str, " ")
	list := make([]int, len(digits))

	for idx, value := range digits {
		list[idx] = StringToInteger(value)
	}

	return list
}
