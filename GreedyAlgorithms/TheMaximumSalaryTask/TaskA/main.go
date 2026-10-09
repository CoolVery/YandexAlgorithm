package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func IsBetter(first, second string) bool {
	tempFirst, _ := strconv.Atoi(first + second)
	tempSecond, _ := strconv.Atoi(second + first)

	if tempFirst > tempSecond {
		return true
	} else if tempFirst == tempSecond{
		tempF, _ := strconv.Atoi(first)
		tempS, _ := strconv.Atoi(second)
		if tempF > tempS {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var countNum int
	fmt.Fscan(in, &countNum)
	//Слайс ввода с консоли
	sliceNum := make([]string, 0, countNum)
	//Слайс по которому будет сделана зарплата, она заполнена -
	sliceSalary := make([]string, 0, countNum)
	for i := 0; i < countNum; i++ {
		sliceSalary = append(sliceSalary, "-")
	}
	for i := 0; i < countNum; i++ {
		var num string
		fmt.Fscan(in, &num)
		sliceNum = append(sliceNum, num)
	}
	for index, num := range sliceNum {
		//Если мы только начали, то записываем число 
		if index == 0 {
			sliceSalary[index] = num
			continue
		}
		//теперь каждое число мы сравниваем с теми, что уже есть в зп
		for indexS, sal := range sliceSalary {
			//Это проверка, что если мы число перекидывали направо (оно маленькое) и сравнить не с чем
			//То записываем по индексу num в зарплату (они совпадают)
			if sal == "-" {
				sliceSalary[index] = num
				break
			}
			//Функция IsBetter сравнивает число из зарплаты и num
			//если окажется, что в зарплате число лучше, то num надо в конец, поэтому продолжаем
			if IsBetter(sal, num) {
				continue
			//если окажется, что num лучше числа зарплаты, то нам надо num поставить перед ним в зарплате
			//поэтому копируем все после этого число из зарплаты со сдвигом один, а на его место пишем num
			} else {
				copy(sliceSalary[indexS+1:], sliceSalary[indexS:len(sliceSalary) - 1])
				sliceSalary[indexS] = num
				break
			}
		}
	}
	salary := strings.Join(sliceSalary, "")
	out.WriteString(salary)
}
