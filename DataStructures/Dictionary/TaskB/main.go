package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var countValue int
	fmt.Fscan(reader, &countValue)
	sliceInt := make([]int, countValue)
	dictionaryValues := make(map[int]int)
	for i := 0; i < countValue; i++ {
		fmt.Fscan(reader, &sliceInt[i])
	}
	//Создаем словарь счетчик всех уникальных значений
	for _, value := range sliceInt {
		dictionaryValues[value]++
	}
	maxValue, maxKey := 0, 0
	for key, value := range dictionaryValues {
		//Если встречается больше, то меняем и запоминаем ключ
		if value > maxValue {
			maxValue = value
			maxKey = key
		}
		//Если встречается одинаково, то сравниваем ключи и запоминаем наименьший
		if value == maxValue {
			if key < maxKey {
				maxKey = key
			}
		}
	}
	fmt.Println(maxKey)
}