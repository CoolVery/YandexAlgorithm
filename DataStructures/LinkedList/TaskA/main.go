package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Количество запросов
	var countQuery int
	//Строка запроса
	var query string
	//Количество прочитанных строк в консоли
	var countQueryStringInConsole int
	//Основной лист
	var linkedList *LinkedList
	//Сканер для построчного ввода
	scanner := bufio.NewScanner(os.Stdin)
	//Считаем количество запросов
	fmt.Scanln(&countQuery)
	for scanner.Scan() {
		//Читаем строку
		query = scanner.Text()
		//Сплитим по пробелам
		querySplit := strings.Split(query, " ")
		switch querySplit[0] {
			case "1":
				//Конвертируем х и у
				xInt, _ := strconv.Atoi(querySplit[1])
				yInt, _ := strconv.Atoi(querySplit[2])
				//Вызываем функцию
				linkedList = linkedList.AddYAfterX(xInt, yInt)
			case "2":
				posithionInList, _ := strconv.Atoi(querySplit[1])
				fmt.Printf("%d\n", linkedList.SearchPosithionInList(posithionInList))
			case "3":
				posithionInList, _ := strconv.Atoi(querySplit[1])
				linkedList = linkedList.DeleteForPosithion(posithionInList)
		}

		countQueryStringInConsole++
		//Проверка на количество прочитанных строк
		if countQueryStringInConsole == countQuery {
			break
		}
	}
}