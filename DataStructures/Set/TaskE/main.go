package main

import (
    "fmt"
    "bufio"
    "os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(reader, &n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	prev := map[int]struct{}{}          // OR подотрезков, заканчивающихся на прошлом шаге
	all  := map[int]struct{}{}          // все уникальные OR за всё время

	for _, x := range a {               // идём по массиву слева направо

	    cur := map[int]struct{}{}       // новые OR
	    cur[x] = struct{}{}             //   ↑ сам x как подотрезк длины 1

    	for v := range prev {           // для каждого старого OR
        	cur[v|x] = struct{}{}       //   ↑ приклеиваем x
    	}

    	for v := range cur {            // всё, что получилось,
        	all[v] = struct{}{}         //   ↑ добавляем в общий ответ
    	}
    	prev = cur                      // текущее становится предыдущим
	}
	fmt.Println(len(all))
}