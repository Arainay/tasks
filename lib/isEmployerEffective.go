package lib

import "fmt"

func IsEmployerEffective(list []int) bool {
	job := map[int]bool{}

	for i := 1; i < len(list); i++ {
		if list[i-1] != list[i] {
			job[list[i-1]] = true
			_, isExists := job[list[i]]
			if isExists {
				fmt.Println("!!!", false)
				return false
			}
		}
	}

	return true
}
