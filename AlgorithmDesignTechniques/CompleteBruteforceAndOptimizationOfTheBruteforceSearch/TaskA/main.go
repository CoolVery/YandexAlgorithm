package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
//Считаем факториал
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var num int
	fmt.Fscan(in, &num)
	p := 1
	multy := 0
	for i := 0; i < num; i++ {
		multy++
		p *= multy
	}
	out.WriteString(strconv.Itoa(p))
}
