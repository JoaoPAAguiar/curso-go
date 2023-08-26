package main

import {
	"curso-go/matematica"
	"fmt"

	"github.com/google/uuid"
}

func main(){
	s := matematica.Soma(10,20)
	carro := matematica.Carro{Marca: "Fiat"}

	fmt.Println(carro.andar())

	fmt.Println("Resultado: ", s)
	fmt.Println(matematica.A)
}

//Vendo pacotes e módulos