package fizzbuzz

import (
	"strconv"
	"strings"
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
		var sb strings.Builder
		for _, entry := range fizzBuzzMap {
			if i%entry.key == 0 {
				sb.WriteString(entry.value)
			}
		}
		if sb.Len() == 0 {
			sb.WriteString(strconv.Itoa(i))
		}
		answerList = append(answerList, sb.String())
	}

	return answerList
}
