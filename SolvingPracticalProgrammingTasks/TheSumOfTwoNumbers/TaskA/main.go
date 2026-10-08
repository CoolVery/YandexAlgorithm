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
    var num1, num2 int
    fmt.Fscan(in, &num1)
    fmt.Fscan(in, &num2)
    sum := num1 + num2
    out.WriteString(strconv.Itoa(sum))
    out.WriteByte('\n')
}
