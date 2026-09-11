package main

import (
	"bufio"
	"fmt"
	"os"
)



func main() {
	set := make(map[int]struct{})
	reader := bufio.NewReader(os.Stdin)
	var typeQuery, valueQuery int
	var countQuery int
	fmt.Fscan(reader, &countQuery)
	for i := 0; i < countQuery; i++ {
		fmt.Fscan(reader, &typeQuery, &valueQuery)
		switch typeQuery{
		case 1:
			//Добавляем элемент в множество
			set[valueQuery] = struct{}{}
		case 2:
			//Проверяем элемент в множестве
			if _, ok := set[valueQuery]; ok {
				fmt.Println(1)
			} else {
				fmt.Println(0)
			}
		}
	}
}
