package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var countValue int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &countValue)
	sliceInt := make([]int, countValue)
	
	for i := 0; i < countValue; i++ {
		fmt.Fscan(reader, &sliceInt[i])
	}
	maxMinValue := sliceInt[0]
	sliceResult := make([]int, 0, countValue)
	sliceResult = append(sliceResult, maxMinValue)
	for i := 1; i < countValue; i++ {
		if sliceInt[i] < maxMinValue {
			maxMinValue = sliceInt[i]
			sliceResult = append(sliceResult, maxMinValue)
			continue
		}
		sliceResult = append(sliceResult, maxMinValue)
	}
	for _, value := range sliceResult {
		fmt.Printf("%d ", value)
	}
}
