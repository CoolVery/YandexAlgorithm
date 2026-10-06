package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type pair struct {
	val   int
	count int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var countValue int
	fmt.Fscan(reader, &countValue)
	slicePair := make([]pair, 0)
	sliceInt := make([]int, countValue)
	dictionaryValues := make(map[int]int)
	for i := 0; i < countValue; i++ {
		fmt.Fscan(reader, &sliceInt[i])
	}
    //делаем словарь, где считается количество повторений уникального элемента
	for _, val := range sliceInt {
		dictionaryValues[val]++
	}
    //Делаем из мапы слайс, чтобы от сортировать
	for key, val := range dictionaryValues {
		slicePair = append(slicePair, pair{key, val})
	}
	slices.SortFunc(slicePair, func(a, b pair) int {
        //Если частоты не равны, то сортируем по убыванию
		if a.count != b.count {
			if a.count > b.count {
				return -1
			}
			return 1
		} 
        //Иначе сравниваем по значие по возрастанию
		if a.val > b.val {
			return 1
		}
		return -1
	})
    //Сортируем по возрастанию 3 первых элемента
	result := []int{slicePair[0].val, slicePair[1].val, slicePair[2].val}
	slices.Sort(result)
	fmt.Println(result[0], result[1], result[2] )
}