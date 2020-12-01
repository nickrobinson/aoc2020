package accounting

import (
	"errors"
	"sort"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func FindTwoMatchingExpenses(expenses []int, sum int) (matchingExpenses []int, err error) {
	sort.Ints(expenses)
	i := 0
	j := len(expenses) - 1

	for i < j {
		if expenses[i]+expenses[j] > sum {
			j--
		} else if expenses[i]+expenses[j] < sum {
			i++
		} else {
			return []int{expenses[i], expenses[j]}, nil
		}
	}

	return []int{}, errors.New("Match not found")
}

func FindThreeMatchingExpenses(expenses []int, sum int) (matchingExpenses []int, err error) {
	sort.Ints(expenses)

	for i := 0; i < len(expenses); i++ {
		for j := 0; j < len(expenses); j++ {
			for k := 0; k < len(expenses); k++ {
				if (expenses[i] + expenses[j] + expenses[k]) == sum {
					return []int{expenses[i], expenses[j], expenses[k]}, nil
				}
			}
		}
	}
	return []int{}, errors.New("Match not found")
}
