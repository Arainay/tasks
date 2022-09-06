package lib

import (
	"regexp"
	"sort"
	"strings"
)

// Вам задан набор отрезков времени. Каждый отрезок задан в формате HH:MM:SS-HH:MM:SS,
// то есть сначала заданы часы, минуты и секунды левой границы отрезка,
// а затем часы, минуты и секунды правой границы.
// Вам необходимо выполнить валидацию заданного набора отрезков времени.
// Иными словами, вам нужно проверить следующие условия:
//  1. часы, минуты и секунды заданы корректно (то есть часы находятся в промежутке от 0
//     0 до 23, а минуты и секунды — в промежутке от 0 до 59);
//  2. левая граница отрезка находится не позже его правой границы (но границы могут быть равны);
//  3. никакая пара отрезков не пересекается (даже в граничных моментах времени).
func ValidateTimePeriods(list []string) bool {
	pattern := `^[0-2]{1}[0-9]{1}:[0-5]{1}[0-9]{1}:[0-5]{1}[0-9]{1}-[0-2]{1}[0-9]{1}:[0-5]{1}[0-9]{1}:[0-5]{1}[0-9]{1}$`
	periods := [][]int{}

	for i := 0; i < len(list); i++ {
		matched, _ := regexp.MatchString(pattern, list[i])

		if !matched {
			return false
		}

		item := strings.ReplaceAll(list[i], ":", "")
		times := strings.Split(item, "-")
		start := StringToInteger(times[0])
		end := StringToInteger(times[1])

		if start > end {
			return false
		}

		periods = append(periods, []int{start, end})
	}

	sort.Slice(periods, func(i, j int) bool {
		return periods[i][0] < periods[j][0]
	})

	for i := 1; i < len(periods); i++ {
		if periods[i][0] >= periods[i-1][0] && periods[i][0] <= periods[i-1][1] || periods[i][1] >= periods[i-1][0] && periods[i][1] <= periods[i-1][1] {
			return false
		}
	}

	return true
}
