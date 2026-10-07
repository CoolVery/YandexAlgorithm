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
	var countQuery int
	fmt.Fscan(in, &countQuery)
	//Делаем наш стэк
	stack := make([]int, 2 * countQuery + 5)
	//Делаем ссылку на head -1 значит стэк пустой
	head := -1
	for i := 0; i < countQuery; i++ {
		var cmd int
		fmt.Fscan(in, &cmd)
		switch cmd {
		case 1:
			//Добавляем в голову
			var num int
			fmt.Fscan(in, &num)
			head++
			stack[head] = num
		case 2:
			//Удаляем голову
			if head > -1 {
				head--
			}
		}
		if head == -1 {
			out.WriteString("-1")
			out.WriteByte('\n')
		} else {
			out.WriteString(strconv.Itoa(stack[head]))
			out.WriteByte('\n')
		}
	}
}
