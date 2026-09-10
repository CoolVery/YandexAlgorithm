package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)
//Структура хранения вычитаний, элементов, их индексов в слайсе
type DateResult struct {
	//Разница
	subtraction int
	//Слайс с элементами разницы 
	elemSub []int
	//Слайс с индексами элементов в слайсе
	indexElemSub []int
}

func FindMinSub(dr []DateResult) {
	//Сохраняем просто первые индексы
	resultIndex := dr[0].indexElemSub
	//Берем просто первый элемент структуры как минимум
	minDR := dr[0]
	//Проходим по всей структуре
	for _, elem := range dr {
		//Если разница меньше, то записываем индексы и меняем минимум
		if elem.subtraction < minDR.subtraction {
			resultIndex = elem.indexElemSub
			minDR = elem
		//Далее проверка, что если разницы равны, т.е. у нас (0 - 2 и 2, 0 - 5 и 5)
		} else if elem.subtraction == minDR.subtraction {
			//По условии задачи мы смотрим сначала на Уменьшаемое (откуда вычитываем), мы должны выбрать тот, где он меньше
			if elem.indexElemSub[0] < minDR.indexElemSub[0] {
				resultIndex = elem.indexElemSub
			} 
			//По условии задачи мы смотрим на Вычитаемое (что вычитаем), мы должны выбрать тот, где он меньше
			if elem.indexElemSub[0] == minDR.indexElemSub[0] {
				if elem.indexElemSub[1] < minDR.indexElemSub[1] {
					resultIndex = elem.indexElemSub
				}
			}
		}
	}
	//Печатаем ответ (по условию задачи нам надо не просто индекс, а его порядковый номер в слайсе)
	fmt.Println(resultIndex[0] + 1, resultIndex[1] + 1)
}
//Все также, как в  FindMinSub, но наоборот
func FindMaxSub(dr []DateResult) {
	resultIndex := dr[0].indexElemSub
	maxSub := dr[0]
	for _, elem := range dr {
		if elem.subtraction > maxSub.subtraction {
			resultIndex = elem.indexElemSub
			maxSub = elem
		} else if elem.subtraction == maxSub.subtraction {
			if elem.indexElemSub[0] < maxSub.indexElemSub[0] {
				resultIndex = elem.indexElemSub
			} 
			if elem.indexElemSub[0] == maxSub.indexElemSub[0] {
				if elem.indexElemSub[1] < maxSub.indexElemSub[1] {
					resultIndex = elem.indexElemSub
				} else {
					resultIndex = elem.indexElemSub
				}
			}
		}
	}
	fmt.Println(resultIndex[0] + 1, resultIndex[1] + 1)

}
//Функция поиска максимум
func SearchMaxInSlice(sliceInt []int) int {
	result := sliceInt[0]
	for _, elem := range sliceInt {
		if elem > result {
			result = elem
		}
	}
	return result
}

//Функция поиска минимальной разницы
//Суть алгоритма, мы идем слева направо и из левого элемента вычитаем максимальный из подслайса
// 1 2 3 4 -> 1 |2 3 4| (1 - 4) -> 1 2 |3 4 | (2 - 4) ...
//после чего формиурем нашу структуру со всеми разницами
//Почему мы вычитаем Максимум, потому что нет смысла вычитать каждый элемент из подслайса, т.к. большее значение даст разница с Максимумом
func SearcMinSub(sliceInt []int) {
	//Создаем слайс нашей структуры ответа, чтобы иметь все данные
	sliceSubtractionDateResult := make([]DateResult, 0)
	//Проходимся по каждому элементу
	for index, elem := range sliceInt {
		//Если мы пришли к последнему элементу, то ломаем цикл
		if index == len(sliceInt) - 1 {
			break
		}
		//Берем подслайс, пропуская текущий элемент и до конца
		subslice := sliceInt[index + 1:]
		//Находим максимум в подслайсе
		maxInSubSlice := SearchMaxInSlice(subslice)		
		//Находим разницу элемента и максимум
		subtraction := elem - maxInSubSlice
		//В слайс со структурой для ответа добавляем новый элемент
		sliceSubtractionDateResult = append(sliceSubtractionDateResult, DateResult{
			subtraction: subtraction,
			elemSub: []int{elem, maxInSubSlice},
			//Здесь механика такая, нам нужен максимальный элемент из подслайса, но индекс из Основного слайса, но не забываем про дубликаты, поэтому
			//Индекс максимума подслайса в слайсе - это позиция текущего элемента + 1 + индекс его в подслайсе
			indexElemSub: []int{index, index + 1 + slices.Index(subslice, maxInSubSlice)},
		})

	}
	FindMinSub(sliceSubtractionDateResult)
}
//Функция поиска максимальной разницы
//Суть алгоритма, мы идем справа налево и из левого элемента вычитаем максимальный из подслайса
// 1 2 3 4 -> |1 2 3| 4 (3 - 4) -> |1 2 |3 4  (2 - 3) ...
//после чего формиурем нашу структуру со всеми разницами
//Почему мы вычитаем Максимум, потому что нет смысла вычитать каждый элемент из подслайса, т.к. большее значение даст разница с Максимумом
func SearcMaxSub(sliceInt []int) {
	sliceSubtractionDateResult := make([]DateResult, 0)
	lenSlice := len(sliceInt)
	for i := lenSlice - 1; i > 0; i-- {
		subslice := sliceInt[0:i]
		maxInSubSlice := SearchMaxInSlice(subslice)		
		subtraction := maxInSubSlice - sliceInt[i]
		sliceSubtractionDateResult = append(sliceSubtractionDateResult, DateResult{
			subtraction: subtraction,
			elemSub: []int{sliceInt[i], maxInSubSlice},
			indexElemSub: []int{slices.Index(subslice, maxInSubSlice), i},
		})
	}
	FindMaxSub(sliceSubtractionDateResult)
}

func main() {
	//Сколько элементов
	var countElem int
	//Строка с элементами 
	var stringElems string
	elemIntSlice := make([]int, 0)
	fmt.Scanln(&countElem)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	stringElems = scanner.Text()
	elemStrSlice := strings.Fields(stringElems)
	//Конвертируем строки в число и заполняем слайс
	for _, elem := range elemStrSlice {
		elemInt, _ := strconv.Atoi(elem)
		elemIntSlice = append(elemIntSlice, elemInt)
	}

	//Ищем минимальную разницу
	SearcMinSub(elemIntSlice)
	//Ищем максимальную разницу 
	SearcMaxSub(elemIntSlice)
}