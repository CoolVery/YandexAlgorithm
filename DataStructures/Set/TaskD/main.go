package main

import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
	var countSet int
	setAll := make([]map[int]struct{}, 0)
    setValues := make(map[int]int)
	fmt.Fscan(reader, &countSet)
	//Считываем наше количество множеств
    for i := 0; i < countSet; i++ {
		var countValues int
		currentMap := make(map[int]struct{})
		fmt.Fscan(reader, &countValues)
		//Считываем склолько элементов в множестве
        for i := 0; i < countValues; i++ {
            var value int
			fmt.Fscan(reader, &value)
			//Читаем значение и записываем в множество 
            currentMap[value] = struct{}{}
        }
		//Добавляем множество в слайс всех множеств
		setAll = append(setAll, currentMap)
		//Проходя множества мы обрабатываем счетчик ВСЕХ уникальных значений
		for value := range setAll[i] {
            setValues[value]++
        }
    }
	//Длина ядра
	lenKron := 0
	//Проходим по множеству ВСЕХ уникальных значений
    for key, value := range setValues {
		//Если элемент встречался НЕ 1 раз (лепесток) или НЕ количество всех множеств (ядро), то уже подсолнуха быть не может
		if value != 1 && value != countSet {
			fmt.Println("NO")
			return
		}
		//Если элемент Ядро, то
		if value == countSet {
			//Увеличиваем длину ядра
			lenKron++
			//В каждом множестве удаляем этот ключ
			for _, set := range setAll {
				delete(set, key)
			}
			continue
		}
	}
	//Выписываем длину ядра и длину всех множеств после удаления
	fmt.Printf("YES\n%d\n", lenKron)
	for _, set := range setAll {
		fmt.Printf("%d ", len(set))
	}
}