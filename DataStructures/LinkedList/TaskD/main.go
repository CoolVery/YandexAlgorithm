package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var countValue int
	//Создаем ридер для чтения из консоли
	reader := bufio.NewReader(os.Stdin)
	//Функция Fscan позволяет пропускать все пробелы, переносы и т.п. - только символы
	fmt.Fscan(reader, &countValue)
	sliceInt := make([]int, countValue)
	//Читаем из консоли значения в слайс
	for i := 0; i < countValue; i++ {
		fmt.Fscan(reader, &sliceInt[i])
	}
	//Создаем максимальный Минумум
	//Логика в чем: у нас есть 1 2 3 1, вместо того, чтобы каждый раз искать минимум в подслайсе, мы сверяем максимальный Минимум и следующий элемент, т.к. он является последним 
	// 1 - 1 (макс минимум 1)
	// 1 2 - 1 (сравнили макс минимум с 2)
	// 1 2 3 - 1 (сравнил макс минимум с 3. т.е. 1 и 3) - 1 и т.п.
	//Всегда макс минум - это первый элемент в начале
	maxMinValue := sliceInt[0]
	sliceResult := make([]int, 0, countValue)
	//Добавили этот первый элемент
	sliceResult = append(sliceResult, maxMinValue)
	for i := 1; i < countValue; i++ {
		//Если текущий элемент меньше этого максимума, то меняем макс минимум и добавляем в результат
		if sliceInt[i] < maxMinValue {
			maxMinValue = sliceInt[i]
			sliceResult = append(sliceResult, maxMinValue)
			continue
		}
		//Иначе просто добавляем наш максимальный минимум
		sliceResult = append(sliceResult, maxMinValue)
	}
	//Печатаем
	for _, value := range sliceResult {
		fmt.Printf("%d ", value)
	}
}
