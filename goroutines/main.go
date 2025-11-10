package main

import (
	"fmt"
	"strconv"
)

func main() {

	for i := 0; i < 10000; i++ {
		// tudo no golang é uma goroutine, se a principal for finalizada, nenhuma outra será executada depois
		go /*assíncono*/ showMessage(strconv.Itoa(i))
	}
}

func showMessage(message string) {
	fmt.Println(message)
}
