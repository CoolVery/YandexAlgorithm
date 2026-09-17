package main

import (
	"bufio"
	"fmt"
	"os"
)

type Sub struct {
	numerator   int
	denominator int
}

func CutBack(sub *Sub) {
	if sub.denominator%sub.numerator == 0 && sub.numerator <= sub.denominator {
		tempDenominator := sub.numerator
		sub.numerator /= tempDenominator
		sub.denominator /= tempDenominator
	}
	if sub.numerator > sub.denominator && sub.numerator%sub.denominator == 0 {
		tempDenominator := sub.denominator
		sub.numerator /= tempDenominator
		sub.denominator /= tempDenominator
	}
}

func FoundMinSubForMax() {
	
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
			numerator:   numerator,
			denominator: denominator,
		}
		sliceSub[i] = newSub
	}

	minSub := Sub{}
	isFoundMin := false
	for _, sub := range sliceSub {
		CutBack(&sub)
		subResult := float64(sub.numerator) / float64(sub.denominator)
		dictionarySubResult[sub] = subResult
		dictionaryResultCount[subResult]++
		if !isFoundMin {
			minSub = sub
			isFoundMin = true
		}
	}

	tempCount := 1
	equalCount := 0
	for _, count := range dictionaryResultCount {
		if count == tempCount {
			equalCount++
		}

		tempCount = count
	}

	if equalCount == len(dictionaryResultCount) {
		for sub, _ := range dictionarySubResult {
			
				if minSub.numerator * sub.denominator > sub.numerator * minSub.denominator {
					minSub = sub
				}
			
		}
	} else {
		maxResult := 0.0
		maxCount := dictionaryResultCount[dictionarySubResult[minSub]]
		for result, count := range dictionaryResultCount {
			if maxCount < count {
				maxCount = count
				maxResult = result
			}
		}


		for sub, result := range dictionarySubResult {
			if result == maxResult {
				if minSub.numerator * sub.denominator > sub.numerator * minSub.denominator {
					minSub = sub
				}
			}
		}
	}


	

	fmt.Printf("%d %d", minSub.numerator, minSub.denominator)
}
