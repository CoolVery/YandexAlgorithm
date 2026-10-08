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
	var top, countRoutes int
	fmt.Fscan(in, &top)
	fmt.Fscan(in, &countRoutes)
	//1 матрица, где ребра идут последовательно
	matrixAdjacent := make([][]int, 0, top + 1)
	//2 матрица, где ребра пересекаются друг с другом
	matrixAchievements := make([][]int, 0, top + 1)
	for i := 0; i < top + 1; i++ {
		matrixAdjacent = append(matrixAdjacent, make([]int, top + 1))
		matrixAchievements = append(matrixAchievements, make([]int, top + 1))

	}
	//Этот слайс хранит информацию о всех маршрутах, количества остановок, и сами точки ребра
	sliceInfo := make([][]int, 0, countRoutes)
	for i := 0; i < countRoutes; i++ {
		var countStop int
		fmt.Fscan(in, &countStop)
		rout := make([]int, 0, countStop)
		for i := 0; i < countStop; i++ {
			var val int
			fmt.Fscan(in, &val)
			rout = append(rout, val)
		}
		sliceInfo = append(sliceInfo, rout)
	}
	//1 матрица
	//----
	for _, rout := range sliceInfo {
		for i := 0; i < len(rout) - 1; i++ {
			matrixAdjacent[rout[i]][rout[i + 1]] = 1
			matrixAdjacent[rout[i + 1]][rout[i]] = 1
		}
	}
	//печать
	for i := 1; i < top + 1; i++ {
		for j := 1; j < top + 1; j++ {
			out.WriteString(strconv.Itoa(matrixAdjacent[i][j]))
			if j != top {
				out.WriteByte(' ')
			}
		}
		out.WriteByte('\n')
	}
	//----
	//2 матрица
	//Здесь суть в том, что rout - это слайс 1,2,3
	//И мы проходим по ней 1 |2,3| - 1 2, 1 3
	// 2 | 3 | - 2 3
	for _, rout := range sliceInfo {
		for i := 0; i < len(rout) - 1; i++ {
			window := rout[i + 1:]
			for _, num := range window {
				matrixAchievements[rout[i]][num] = 1
				matrixAchievements[num][rout[i]] = 1
			}
		}
	}
	for i := 1; i < top + 1; i++ {
		for j := 1; j < top + 1; j++ {
			out.WriteString(strconv.Itoa(matrixAchievements[i][j]))
			if j != top {
				out.WriteByte(' ')
			}
		}
		out.WriteByte('\n')
	}
	//----
}