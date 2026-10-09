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
	var countGroup, countPeople int
	fmt.Fscan(in, &countGroup) 
	//Этот слайс хранит информацию о всех группах, количества людей в группах, и самих людях
	sliceInfo := make([][]int, 0, countGroup)
	for i := 0; i < countGroup; i++ {
		var countPeopleInGroup int
		fmt.Fscan(in, &countPeopleInGroup)
		group := make([]int, 0, countPeopleInGroup)
		for i := 0; i < countPeopleInGroup; i++ {
			var val int
			fmt.Fscan(in, &val)
			group = append(group, val)
			if val > countPeople {
				countPeople = val
			}
		}
		sliceInfo = append(sliceInfo, group)
	}
	//-----
	workCountPeopleInMatrix := countPeople + 1
	matrixAchievements := make([][]int, 0, workCountPeopleInMatrix)
	for i := 0; i < workCountPeopleInMatrix; i++ {
		matrixAchievements = append(matrixAchievements, make([]int, workCountPeopleInMatrix))
	}
	for _, group := range sliceInfo {
		for i := 0; i < len(group) - 1; i++ {
			window := group[i + 1:]
			for _, num := range window {
				matrixAchievements[group[i]][num] = 1
				matrixAchievements[num][group[i]] = 1
			}
		}
	}
	//печать
	for i := 1; i < workCountPeopleInMatrix; i++ {
		for j := 1; j < workCountPeopleInMatrix; j++ {
			out.WriteString(strconv.Itoa(matrixAchievements[i][j]))
			if j != countPeople {
				out.WriteByte(' ')
			}
		}
		out.WriteByte('\n')
	}
}