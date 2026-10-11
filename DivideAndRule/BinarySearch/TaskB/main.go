package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	
	var count, countSearch int
	fmt.Fscan(in, &count)
	nums := make([]int, 0, count)
	for i := 0; i < count; i++ {
		var num int
		fmt.Fscan(in, &num)
		nums = append(nums, num)
	}
	fmt.Fscan(in, &countSearch)
	for i := 0; i < countSearch; i++ {
		var search int
		fmt.Fscan(in, &search)
		minIndex := 0
		maxIndex := count - 1
		resultIndex := -1
		for maxIndex >= minIndex {
			midIndex := (minIndex + maxIndex) / 2
			if nums[midIndex] == search {
				resultIndex = midIndex
				break
			}
			if nums[midIndex] < search {
				minIndex = midIndex + 1
				continue
			}
			maxIndex = midIndex - 1
		}
		out.WriteString(strconv.Itoa(resultIndex))
		out.WriteByte('\n')
	}
}