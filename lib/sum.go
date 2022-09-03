package lib

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// В первой строке входных данных содержится целое число t (1≤𝑡≤10^4) — количество наборов входных данных в тесте.
// Далее следуют описания t наборов входных данных, один набор в строке.
// В  каждой строке набора записаны два целых числа a и b (−1000≤𝑎,𝑏≤1000)
// Для каждого набора входных данных выведите сумму двух заданных чисел, то есть a+b
func Sum() {
	var t int
	var list []int

	// bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(os.Stdin)
	i := 0

	for scanner.Scan() {
		text := scanner.Text()

		if i == 0 {
			intValue, err := strconv.Atoi(text)
			if err == nil {
				t = intValue
				list = make([]int, t)
			}
		} else {
			sum := 0
			digits := strings.Split(text, " ")

			for _, value := range digits {
				intValue, err := strconv.Atoi(value)
				if err == nil {
					sum += intValue
				}
			}

			list[i-1] = sum
		}

		i++

		if i == t+1 {
			break
		}
	}

	for _, sum := range list {
		fmt.Println(sum)
	}
}
