package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	canal := make(chan bool)
	go NombreLento("Hola", canal)

	var result bool
	result = <-canal
	if result {
		fmt.Println("Ejecucion finalizada")
	}
}

//------------------------------------------------------------>

//NombreLento imprime un texto letra por letra esperando un segundo
func NombreLento(nombre string, canal chan bool) {
	letras := strings.Split(nombre, "")

	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}

	canal <- true
}
