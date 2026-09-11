package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var countSet int
	set := make(map[int]int)
	fmt.Fscan(reader, &countSet)
	//Читаем и запускаем цикл - сколько множество
	for i := 0; i < countSet; i++ {
		var countValues int
		fmt.Fscan(reader, &countValues)
		//Читаем сколько элементов в множестве
		for i := 0; i < countValues; i++ {
			var value int
			fmt.Fscan(reader, &value)
			//Читаем элемент и проверяем в множестве
			set[value]++
		}
	}
	var result int
	for _, value := range set {
		if value == countSet {
			result++
		}
	}
	fmt.Println(result)
}