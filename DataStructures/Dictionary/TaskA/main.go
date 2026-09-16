package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var countQuery int
	fmt.Fscan(reader, &countQuery)
	dictionary := make(map[int]int)
	for i := 0; i < countQuery; i++ {
		var typeQuery int 
		fmt.Fscan(reader, &typeQuery)
		switch typeQuery {
		//Добавляем запись в словарь
		case 1:
			var key, value int
			fmt.Fscan(reader, &key, &value)
			dictionary[key] = value
		//Читаем значение по ключу
		case 2:
			var key int
			fmt.Fscan(reader, &key)
			if value, ok := dictionary[key]; ok {
				fmt.Println(value)
			} else {
				fmt.Println(-1)
			}
		}
	}
}