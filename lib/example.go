package lib

import (
	"bufio"
	"fmt"
	"os"
)

func Example() {
	limit := 0
	i := 0
	num := ""
	fullCostMap := map[string]int{}
	fullCost := 0

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		text := string(scanner.Text())
		isNewLine := text == "\n"
		isSpace := text == " "

		if isSpace {
			fullCostMap[num]++
			num = ""
		} else if isNewLine {
			if limit == 0 {
				limit = StringToInteger(num)*2 + 1
			} else if i%2 == 0 {
				fullCostMap[num]++
				for cost, count := range fullCostMap {
					fullCost += StringToInteger(cost) * (count - count/3)
				}

				fmt.Println(fullCost)
			}

			fullCostMap = map[string]int{}
			fullCost = 0
			num = ""
			i++
		} else if text != "\r" {
			num += text
		}

		if i >= limit && limit > 0 {
			break
		}
	}
}
