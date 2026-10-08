package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
Задача по поиску НОД - наибольшему общему делителю
суть алгоритма
Есть два числа (48 и 18) и с ним мы делаем данный процесс
48 % 18 = 12
18 % 12 = 6
12 % 6 = 0
Когда остаток равен 0, то ответ 6
Если а и 0, то ответ a
Если 0 и b, то ответ b
Если 0 и 0, то ответ 0
*/
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var num1, num2 int
	fmt.Fscan(in, &num1)
	fmt.Fscan(in, &num2)
	if num1 != 0 && num2 == 0 {
		out.WriteString(strconv.Itoa(num1))
		return
	}
	if num1 == 0 && num2 != 0 {
		out.WriteString(strconv.Itoa(num2))
		return
	}
	if num1 == 0 && num2 == 0 {
		out.WriteString(strconv.Itoa(0))
		return
	}
	for {
		mod := num1 % num2
		if mod == 0 {
			out.WriteString(strconv.Itoa(num2))
			break
		}
		num1 = num2
		num2 = mod
	}
}
