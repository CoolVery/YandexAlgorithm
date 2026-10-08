package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
/*
Задача по поиску НОК - наименьшему общему кратному
суть алгоритма
НОК - это произведение двух чисел деленное на НОД этих двух чисел
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
	if num1 == 0 || num2 == 0 {
		out.WriteString(strconv.Itoa(0))
		return
	}
	gcd := 0
	//Ищем НОД
	temp1, temp2 := num1, num2
	for {
		mod := temp1 % temp2
		if mod == 0 {
			gcd = temp2
			break
		}
		temp1 = temp2
		temp2 = mod
	}
	//-----
	lcm := (num1 / gcd) * num2
	out.WriteString(strconv.Itoa(lcm))

}
