package fizzbuzz

import (
	"strconv"
)

func fizzBuzz(n int) []string {
	fizzBuzzMap := []struct {
		key   int
		value string
	}{
		{3, "Fizz"},
		{5, "Buzz"},
	}

	var answerList []string
	for i := 1; i <= n; i++ {
		var result string
		for _, entry := range fizzBuzzMap {
			if i%entry.key == 0 {
				result += entry.value
			}
		}
		if result == "" {
			result = strconv.Itoa(i)
		}
		answerList = append(answerList, result)
	}

	return answerList
}
