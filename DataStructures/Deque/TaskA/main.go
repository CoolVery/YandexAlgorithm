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
//Удаляем голову
func deleteHead(deque []int) []int {
	return deque[1:]
}
//Печатаем голову
func printHead(deque []int) {
	if len(deque) == 0 {
		fmt.Println(-1)
		return
	}
	fmt.Println(deque[0])
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
			deque = addInEnd(deque, num)
		case "2":
			deque = deleteHead(deque)
		}
		printHead(deque)
	}
}
