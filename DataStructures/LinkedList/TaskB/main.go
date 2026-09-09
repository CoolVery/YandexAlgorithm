package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type DateResult struct {
	subtraction int 
	elemSub []int
	indexElemSub []int
}

func FindMinSub(dr []DateResult) {
	resultIndex := dr[0].indexElemSub
	minDR := dr[0]
	for _, elem := range dr {
		if elem.subtraction < minDR.subtraction {
			resultIndex = elem.indexElemSub
			minDR = elem
		} else if elem.subtraction == minDR.subtraction {
			if elem.indexElemSub[0] < minDR.indexElemSub[0] {
				resultIndex = elem.indexElemSub
			} 
			if elem.indexElemSub[0] == minDR.indexElemSub[0] {
				if elem.indexElemSub[1] < minDR.indexElemSub[1] {
					resultIndex = elem.indexElemSub
				} else {
					resultIndex = elem.indexElemSub
				}
			}
		}
	}

	fmt.Println(resultIndex[0] + 1, resultIndex[1] + 1)
}

func FindMaxSub(dr []DateResult) {
	resultIndex := dr[0].indexElemSub
	maxSub := dr[0]
	for _, elem := range dr {
		if elem.subtraction > maxSub.subtraction {
			resultIndex = elem.indexElemSub
			maxSub = elem
		} else if elem.subtraction == maxSub.subtraction {
			if elem.indexElemSub[0] < maxSub.indexElemSub[0] {
				resultIndex = elem.indexElemSub
			} 
			if elem.indexElemSub[0] == maxSub.indexElemSub[0] {
				if elem.indexElemSub[1] < maxSub.indexElemSub[1] {
					resultIndex = elem.indexElemSub
				} else {
					resultIndex = elem.indexElemSub
				}
			}
		}
	}
	fmt.Println(resultIndex[0] + 1, resultIndex[1] + 1)

}


func SearchMaxInSlice(sliceInt []int) int {
	var result int
	for _, elem := range sliceInt {
		if elem > result {
			result = elem
		}
	}
	return result
}

func SearcMinSub(sliceInt []int) {
	sliceSubtraction := make([]int, 0)
	sliceSubtractionDateResult := make([]DateResult, 0)
	for index, elem := range sliceInt {
		if index == len(sliceInt) - 1 {
			break
		}
		subslice := sliceInt[index + 1:]
		maxInSubSlice := SearchMaxInSlice(subslice)		
		subtraction := elem - maxInSubSlice
		sliceSubtraction = append(sliceSubtraction, subtraction)
		sliceSubtractionDateResult = append(sliceSubtractionDateResult, DateResult{
			subtraction: subtraction,
			elemSub: []int{elem, maxInSubSlice},
			indexElemSub: []int{index, slices.Index(sliceInt, maxInSubSlice)},
		})
	}
	for _, el := range sliceSubtractionDateResult {
		fmt.Println(el)
	}

	FindMinSub(sliceSubtractionDateResult)
}

func SearcMaxSub(sliceInt []int) {
	sliceSubtraction := make([]int, 0)
	sliceSubtractionDateResult := make([]DateResult, 0)
	lenSlice := len(sliceInt)
	for i := lenSlice - 1; i > 0; i-- {
		// if i == 0 {
		// 	break
		// }
		subslice := sliceInt[0:i]
		maxInSubSlice := SearchMaxInSlice(subslice)		
		subtraction := maxInSubSlice - sliceInt[i]
		sliceSubtraction = append(sliceSubtraction, subtraction)
		sliceSubtractionDateResult = append(sliceSubtractionDateResult, DateResult{
			subtraction: subtraction,
			elemSub: []int{sliceInt[i], maxInSubSlice},
			indexElemSub: []int{slices.Index(sliceInt, maxInSubSlice), i},
		})
	}
	
	for _, el := range sliceSubtractionDateResult {
		fmt.Println(el)
	}

	FindMaxSub(sliceSubtractionDateResult)
}

func main() {
	var countElem int
	var stringElems string
	elemIntSlice := make([]int, 0)
	fmt.Scanln(&countElem)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	stringElems = scanner.Text()
	elemStrSlice := strings.Split(stringElems, " ")
	for _, elem := range elemStrSlice {
		elemInt, _ := strconv.Atoi(elem)
		elemIntSlice = append(elemIntSlice, elemInt)
	}


	SearcMinSub(elemIntSlice)
	SearcMaxSub(elemIntSlice)
}