package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	//Читаем данные
	reader := bufio.NewReader(os.Stdin) 
	var countValue int
	fmt.Fscan(reader, &countValue)
	sliceInt := make([]int, countValue)
	for i := 0; i < countValue; i++ {
		fmt.Fscan(reader, &sliceInt[i])
	}
	//Создаем мапу, где каждый уникальный элемент хранит в себе последний индекс в слайсе
	mapValues := make(map[int]int)
	for index, value := range sliceInt {
		mapValues[value] = index
	}
	//Находим максимальное значения
	maxValue := slices.Max(sliceInt)
	//Проходим по слайсу, если индексы совпали из мапы, то пропускаем в печати
	for i := 0; i < countValue; i++ {
		if mapValues[maxValue] != i {
			fmt.Printf("%d ", sliceInt[i])
		}
	}
}
