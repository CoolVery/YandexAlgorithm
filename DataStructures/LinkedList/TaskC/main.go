package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var countValue, countQuery int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	stringValue := strings.Fields(scanner.Text())
	//Значение в массиве
	countValue, _ = strconv.Atoi(stringValue[0])
	//Количество запросов
	countQuery, _ = strconv.Atoi(stringValue[1])

	sliceInt := make([]int, 0, countValue)
	scanner.Scan()
	sliceString := strings.Fields(scanner.Text())
	//Конвертируем строк и числа
	for _, elem := range sliceString {
		elemInt, _ := strconv.Atoi(elem) 
		sliceInt = append(sliceInt, elemInt)
	}
	//Создаем мапу, где ключ - это элемент из слайса, а значение - это его перый индекс
	mapInt := make(map[int]int, len(sliceInt))
	for index, elem := range sliceInt {
		//Если в мапе нет элемента слайса, то мы создаем запись, куда пишем его индекс
		if _, ok := mapInt[elem]; !ok {
			mapInt[elem] = index + 1
		}
	}
	for i := 0; i < countQuery; i++ {
		scanner.Scan()
		elemInSlice, _ := strconv.Atoi(scanner.Text())
		//Если значения есть, выводим индекс, если нет, то -1
		if value, ok := mapInt[elemInSlice]; ok {
			fmt.Println(value)
		} else {
			fmt.Println(-1)
		}
	}
}