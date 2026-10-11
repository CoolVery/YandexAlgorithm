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
	const MOD = 1_000_000_007

	var countMoney, money int
	fmt.Fscan(in, &countMoney, &money)
	sliceMoney := make([]int, countMoney)
	dp := make([]int, money + 1)
	for i := 0; i < countMoney; i++ {
		fmt.Fscan(in, &sliceMoney[i])
	}
	dp[0] = 1
	for _, m := range sliceMoney {
		for i := m; i <= money; i++ {
			dp[i] = (dp[i] + dp[i - m]) % MOD
		}
	}
	out.WriteString(strconv.Itoa(dp[money]))
	out.WriteByte('\n')
}