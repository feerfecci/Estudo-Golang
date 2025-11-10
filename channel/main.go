package main

import (
	"fmt"
	"time"
)

func main() {

	// var list []int
	// go setList(&list)
	// for _, v := range list {
	// 	fmt.Println(v)
	// }

	// channel bufferizado, serve para primeiro escrita tudo que foi definido no segundo parametro
	channel := make(chan int, 100)
	go setList(channel)

	for v := range channel {
		fmt.Println("recebendo: ", v)
		time.Sleep(time.Second)
	}

}

// func setList(list *[]int) {
//
//	for i := 0; i < 100; i++ {
//		(*list) = append((*list), i)
//	}
//
// }

// com <-chan sinaliza que o channel é leitura ou ou channel<- escrita
func setList(channel chan<- int) {
	for i := 0; i < 100; i++ {
		//sempre que escrevo o channel, travo a gorotine, e alguem PRECISA ouvir esse canal
		channel <- i
		fmt.Println("enviando: ", i)
	}

	//sempre fecha o channel
	close(channel)
}
