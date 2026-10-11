package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	//
	var money int
	fmt.Fscan(in, &money)
	nums := []int{1, 3, 4}
	setDp := make([]int, money + 1)
	setDp[0] = 0
	for i := 1; i < money + 1; i++ {
		setDp[i] = math.MaxInt
		for _, num := range nums {
			if i >= num {
				candidate := setDp[i-num] + 1
				if candidate < setDp[i] {
    				setDp[i] = candidate
				}
			}
		}
	}
	out.WriteString(strconv.Itoa(setDp[money]))
}
