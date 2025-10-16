package main

import (
	"fmt"
	"math"
)

type Geometria interface {
	area() float64
}

type Retangulo struct {
	largura, altura float64
}

func (r Retangulo) area() float64 {
	return r.altura * r.largura
}

type Circulo struct {
	radius float64
}

func (c Circulo) area() float64 {
	return math.Pi * c.radius * c.radius
}

func ExibiAreaGeometria(g Geometria) {
	fmt.Println("area ret:  ", g.area())
}

// func ExibiAreaRetangulo(retangulo Retangulo) {
// 	area := retangulo.altura * retangulo.largura
// 	fmt.Println("area ret:  ", area)
// }

// func ExibiAreaCirculo(circulo Circulo) {
// 	area := math.Pi * circulo.radius * circulo.radius
// 	fmt.Println("area circ:  ", area)
// }

func main() {
	fmt.Println("iniciando...")

	// AULA 50
	// retangulo := Retangulo{
	// 	largura: 1,
	// 	altura:  2,
	// }

	// circulo := Circulo{
	// 	radius: 3,
	// }

	// ExibiAreaGeometria(retangulo)
	// ExibiAreaGeometria(circulo)

	// var err error
	// err = errors.New("teste")
	// fmt.Println(err)
	// ExibeError(errors.New("teste"))

	// p := ProblemaNet{
	// 	rede:     true,
	// 	hardware: false,
	// }

	// ExibeError(p)

	//AULA 52 interface vázia
	// var a interface{}
	// a = 5
	// a = Circulo{
	// fmt.Println(a)
	// 	radius: 10,
	// }
	var lista []interface{}
	lista = append(lista, 10)
	lista = append(lista, 7.5)
	lista = append(lista, true)
	lista = append(lista, "teste")
	for _, valor := range lista {
		if v, ok := valor.(string); ok {
			fmt.Println(v + " string")

		} else {
			fmt.Println(valor)

		}
	}

}

type ProblemaNet struct {
	rede     bool
	hardware bool
}

func (p ProblemaNet) Error() string {
	if p.rede {
		return "Problema de rede"
	} else if p.hardware {
		return "Problema de hardware"
	} else {
		return "Outro Problema"
	}
}

func ExibeError(err error) {
	fmt.Println(err.Error())
}
