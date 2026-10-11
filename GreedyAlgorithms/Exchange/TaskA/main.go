package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var money int
	fmt.Fscan(in, &money)

	// сначала соберём все наборы
	sets := make([][]int, 0, 10)
	//Задача решается через уравнение 10a - 5b - c = money
	//Почему для 10 мы используем <= - потому что 10 - это наш максимальным наминал
	for cnt10 := 0; cnt10*10 <= money; cnt10++ {
		//rest10 - это остаток money без наших десяток
		rest10 := money - cnt10*10
		for cnt5 := 0; cnt5*5 <= rest10; cnt5++ {
			//количество единиц - это вычитание из остатка без десяток всех пятерок
			cnt1 := rest10 - cnt5*5
			//Теперь формируем каждый набор по числу монет
			set := make([]int, 0, cnt10+cnt5+cnt1)
			for i := 0; i < cnt10; i++ { set = append(set, 10) }
			for i := 0; i < cnt5; i++  { set = append(set, 5) }
			for i := 0; i < cnt1; i++  { set = append(set, 1) }
			sets = append(sets, set)
		}
	}

	out.WriteString(strconv.Itoa(len(sets)))
	out.WriteByte('\n')
	for _, set := range sets {
		out.WriteString(strconv.Itoa(len(set)))
		out.WriteByte(' ')
		for i, v := range set {
			out.WriteString(strconv.Itoa(v))
			if i == len(set)-1 {
				out.WriteByte('\n')
			} else {
				out.WriteByte(' ')
			}
		}
	}
}