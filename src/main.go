package main

import (
	"fmt"
)

func main() {
	//decalaracion de constante
	// const pi float64 = 3.14
	// const pip2 = 31.0
	// fmt.Println(pi, pip2)

	//declaracion de varible
	base := 12

	var altura string = "hola"
	var numero int = 14
	fmt.Println(altura, numero, base)

	//area cuadrada 
	const baseCuadrada = 10
	areaCuadrada := baseCuadrada * baseCuadrada
	fmt.Println("Area cuadrado:", areaCuadrada)

	x := 10
	y := 50 

	//suma
	result := x*x + y 
	fmt.Println(result)

	helloMessage := "Hello"
	worldMessage := "world"

	// Println: Salto de Linea Automatico
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	// Printf
	nombre := "Platzi"
	cursos := 500
	// Con valores seguros
	fmt.Printf("%s tiene más de %d cursos\n", nombre, cursos)
	// Con valores inseguros
	fmt.Printf("%v tiene más de %v cursos\n", nombre, cursos)

	// Sprintf
	message := fmt.Sprintf("%v tiene más de %v cursos\n", nombre, cursos)
	fmt.Println(message)

	// Tipo de datos:
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T\n", cursos)
}