package main

import (
	"bufio"
	"fmt"
	"os"
)

type Sub struct {
	numerator int
	denominator int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var countValue int
	fmt.Fscan(reader, &countValue)
	sliceSub := make([]Sub, countValue)
	dictionarySubResult := make(map[Sub]float64)
	dictionaryResultCount := make(map[float64]int)

	for i := 0; i < countValue; i++ {
		var numerator, denominator int
		fmt.Fscan(reader, &numerator, &denominator)
		newSub := Sub{
			numerator: numerator,
			denominator: denominator,
		}
		sliceSub[i] = newSub 
	}

	for _, sub := range sliceSub {
		dictionarySubResult[sub] = float64(sub.numerator) / float64(sub.denominator)
	}

	for _, result := range dictionarySubResult {
		dictionaryResultCount[result]++
	}
	maxValue := 0
	maxResult := 0.0
	for key, value := range dictionaryResultCount {
		if value > maxValue {
			maxResult = key
			maxValue = value
		}
	}
	minSub := sliceSub[0]
	for sub, result := range dictionarySubResult {
		if result == maxResult {
			if minSub.numerator * sub.denominator > sub.numerator * minSub.denominator {
				minSub = sub
			}
		}
	}
	fmt.Printf("%d %d", minSub.numerator, minSub.denominator)
}