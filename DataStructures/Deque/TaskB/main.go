package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
//Добавляем в конец
func addInEnd(deque []int, num int) []int {
	return append(deque, num)
}
//Добавляем в начало
func addInHead(deque []int, num int) []int {
	temp := make([]int, 0, len(deque) + 1)
	temp = append(temp, num)
	temp = append(temp, deque...)
	return temp
}
//Удаляем голову
func deleteHead(deque []int) []int {
	if len(deque) != 0 {
		return deque[1:]
	}
	return deque
}
//Удаляем конец
func deleteEnd(deque []int) []int {
	if len(deque) != 0 {
		return deque[:len(deque) - 1]
	}
	return deque
}
//Печатаем голову
func printHeadAndEnd(deque []int) {
	if len(deque) == 0 {
		fmt.Println(-1)
		return
	} else if len(deque) == 1 {
		fmt.Printf("%d %d\n", deque[0], deque[0])
		return
	}
			fmt.Printf("%d %d\n", deque[0], deque[len(deque) - 1])

}
func main() {
	deque := make([]int, 0)
	reader := bufio.NewReader(os.Stdin)
	var countValue int
	fmt.Fscan(reader, &countValue)
	reader.ReadString('\n')
	for i := 0; i < countValue; i++ {
		var line string
		line, _ = reader.ReadString('\n')
		split := strings.Fields(line)
		switch split[0] {
		case "1":
			num, _ := strconv.Atoi(split[1])
			deque = addInHead(deque, num)
		case "2":
			num, _ := strconv.Atoi(split[1])
			deque = addInEnd(deque, num)
		case "3":
			deque = deleteHead(deque)
		case "4":
			deque = deleteEnd(deque)
		}
		printHeadAndEnd(deque)
	}
}
