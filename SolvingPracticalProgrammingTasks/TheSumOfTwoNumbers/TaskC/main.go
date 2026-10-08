package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
    var lenStr int
    var string1, string2 string
    fmt.Fscan(in, &lenStr)
    fmt.Fscan(in, &string1)
    if len(string1) != lenStr {
        return
    }
    fmt.Fscan(in, &string2)
    if len(string2) != lenStr {
        return
    }
    string1Rune, string2Rune := []rune(string1), []rune(string2)
    sliceRune := make([]rune, 0, lenStr*2)
    for i := 0; i < lenStr; i++ {
        sliceRune = append(sliceRune, string1Rune[i])
        sliceRune = append(sliceRune, string2Rune[i])
    }
    result := string(sliceRune)
    out.WriteString(result)
}