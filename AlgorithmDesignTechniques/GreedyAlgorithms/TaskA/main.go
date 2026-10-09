package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)
//Сделали структуру интервала
type Interval struct {
	start int
	end int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var countInterval int
	fmt.Fscan(in, &countInterval)
	sliceInterval := make([]Interval, 0, countInterval)
	var start, end int
	//Заполняем все интервалы
	for i := 0; i < countInterval; i++ {
		fmt.Fscan(in, &start, &end)
		sliceInterval = append(sliceInterval, Interval{
			start: start,
			end: end,
		})
	}
	//Сортиурем по концу каждого интервала по возрастанию
	slices.SortFunc(sliceInterval, func(a, b Interval) int {
		switch {
		case a.end > b.end:
			return 1
		case a.end < b.end:
			return -1
		default:
			return 0
		}
	})
	//Запоминаем первый конец первого интервала и счетчик
	latestEnd := sliceInterval[0].end
	count := 1
	//Если начала интервала больше максимального конца, то перезаписываем
	//и увеличиваем счетчик
	for _, i := range sliceInterval {
		if i.start > latestEnd {
			count++
			latestEnd = i.end
		}
	}
	out.WriteString(strconv.Itoa(count))
	out.WriteByte('\n')
}