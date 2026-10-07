package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//Печатаем голову
func printHeadAndEnd(deque []int, head, tail int) {
	if head > tail {
		fmt.Println(-1)
		return
	}
	fmt.Printf("%d %d\n", deque[head],  deque[tail])

}
func main() {
	//Основная идея - у нас есть готовый слайс с несколькими 0
    //Мы проходим и заполняем с помощью двух индексов - head и tail
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var countValue int
	fmt.Fscan(in, &countValue)
	deque := make([]int, 2*countValue + 5)
	//Изначально голова будет больше хвоста - значит наш дек пустой
	head := len(deque) / 2
	tail := head - 1
	for i := 0; i < countValue; i++ {
		var line int
		fmt.Fscan(in, &line)
		switch line {
		//Весь функционал - смещаем эти два индекса либо вперед или назад
		case 1:
			var num int
        	fmt.Fscan(in, &num)
			head--
			deque[head] = num
		case 2:
			var num int
        	fmt.Fscan(in, &num)
			tail++
			deque[tail] = num
		//При удалении смотрим, удаляем то или иное только при не пустом слайсе
		case 3:
			if head <= tail {
				head++
			}
		case 4:
			if tail >= head {
				tail--
			}
		}
		if head > tail {
			out.WriteString("-1\n")
		} else {
			out.WriteString(strconv.Itoa(deque[head]))
			out.WriteByte(' ')
			out.WriteString(strconv.Itoa(deque[tail]))
			out.WriteByte('\n')
		}
	}
}
