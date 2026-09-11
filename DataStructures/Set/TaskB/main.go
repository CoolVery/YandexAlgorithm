package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var countSet int
	set := make(map[int]struct{})
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
			if _, ok := set[value]; !ok {
				set[value] = struct{}{}
			}
		}		
	}
	fmt.Println(len(set))
}