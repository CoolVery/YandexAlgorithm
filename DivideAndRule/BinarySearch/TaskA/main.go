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
	
	var count, search int
	fmt.Fscan(in, &count)
	minIndex := 0
	maxIndex := count - 1
	resultIndex := -1
	nums := make([]int, 0, count)
	for i := 0; i < count; i++ {
		var num int
		fmt.Fscan(in, &num)
		nums = append(nums, num)
	}
	fmt.Fscan(in, &search)
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

}