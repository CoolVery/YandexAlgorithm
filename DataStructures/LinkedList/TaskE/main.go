package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//Считываем все данные
	reader := bufio.NewReader(os.Stdin)
	var countValue int
	fmt.Fscan(reader, &countValue)
	sliceInt := make([]int, countValue)
	resultSliceInt := make([]int, 0)
	for i := 0; i < countValue; i++ {
		fmt.Fscan(reader, &sliceInt[i])
	}
	//Первый и последний элемент всегда будет в результате
	resultSliceInt = append(resultSliceInt, sliceInt[0])
	//Начинаем со второго, не доходя до последнего
	for i := 1; i < countValue - 1; i++ {
		//Если текущий элемент подходит под локал минимум, то просто пропускаем, иначе добавляем в итоговый массив
		if sliceInt[i - 1] > sliceInt[i] {
			if sliceInt[i] < sliceInt[i + 1] {
				continue
			}
		}
		resultSliceInt = append(resultSliceInt, sliceInt[i])
	}
	//Добавляем последний элемент
	resultSliceInt = append(resultSliceInt, sliceInt[len(sliceInt) - 1])
	fmt.Printf("%d\n", len(resultSliceInt))
	for _, value := range resultSliceInt {
		fmt.Printf("%d ", value)
	}
}
