package main

import (
	"fmt"
	"os"
)

//PANIC DEFER
// func ShowText() {
// 	fmt.Println("Finalizando")
// }

// func main() {
// 	file, err := os.Create("./settings.txt")

// 	if err != nil {
// 		panic(err)
// 	}

// 	defer file.Close() // defer é usado para colocar a função no final do método /
// 	defer ShowText()

// 	_, err = file.Write([]byte("teste"))

// 	if err != nil {
// 		panic(err)
// 	}

// }

// RECOVER
func ReadFile() {

	//caso queiramos que o código não pare no Panic e pare aqui
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recuperado")
		}
	}() //função anonima
	// colocando dentro do método faz com que o código vá até o fim

	_, err := os.Open("./settings.txt")
	if err != nil {
		panic(err)
	}
}

func main() {

	ReadFile()

	fmt.Println("Fim")
}
