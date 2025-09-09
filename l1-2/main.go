package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	//одна из штук, которая ждет выполнения горутин
	var wg sync.WaitGroup

	for _, number := range numbers {
		wg.Add(1)
		//анонимная горутина
		go func(num int) {
			defer wg.Done()
			sqr := num * num
			fmt.Printf("%d * %d = %d\n", num, sqr, sqr)
		}(number)
	}
	//дожидается выполнения всех горутин
	wg.Wait()
}
