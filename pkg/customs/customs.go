package customs

func GetCustomsCount(answers string, groupCount int) int {
	seenAnswers := map[rune]int{}
	for _, c := range answers {
		seenAnswers[c] += 1
	}

	if groupCount > 1 {
		answerCount := 0
		for _, val := range seenAnswers {
			if val == groupCount {
				answerCount++
			}
		}
		return answerCount
	} else {
		return len(seenAnswers)
	}
}
