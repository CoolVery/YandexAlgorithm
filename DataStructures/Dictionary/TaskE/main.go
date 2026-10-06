package main

import (
	"bufio"
	"fmt"
	"os"
)

type agg struct {
	//сколько всего вхождений слов с этим шаблоном.
	total int
	//сколько пар можно составить внутри каждого слова (включая само с собой)
	sumSq int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var countValue int
	fmt.Fscan(reader, &countValue)
	sliceStr := make([]string, countValue)
	dictUnique := make(map[string]int)
	for i := 0; i < countValue; i++ {
		fmt.Fscan(reader, &sliceStr[i])
		dictUnique[sliceStr[i]]++
	}
	dictMask := make(map[string]*agg)
	for word, f := range dictUnique {
		for i := 0; i < len(word); i++ {
			//Делаем шаблон слово, меняя по символу на *
			pattern := word[:i] + "*" + word[i+1:]
			//Если в словаре нет такого шаблона, то добавляем
			if dictMask[pattern] == nil {
				dictMask[pattern] = &agg{}
			}
			// Сколько раз это слово есть в массиве — столько раз и шаблон от него.
			dictMask[pattern].total += f
			// А тут прибавляем квадрат частоты — это понадобится,
			// чтобы потом не считать пары одинаковых слов.
			dictMask[pattern].sumSq += f * f
		}
	}
	sum := 0
	for _, a := range dictMask {
		//a.total*a.total все пары (включая ненужные)
		sum += (a.total*a.total - a.sumSq) / 2
	}
	fmt.Println(sum)
}