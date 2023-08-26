package main

import {
	"buffer"
	"fmt"
	"os"
}

func main(){
	f, err := os.Create("arquivo.txt")
	if err != nill{
		panic(err)
	}

	tamanho, err := f.Write([]byte("Escrevendo dados no arquivo"))
	//Tamanho, err := f.WriteString("Hello, World!")
	if err != nill{
		panic(err)
	}
	fmt.println("Arquivo criado com sucesso! Tamanho: %d bytes", tamanho)

	f.Close()

	///leitura
	arquivo. err := os.ReadFile("arquivo.txt")
	if err != nill{
		panic(err)
	}
	fmt.Println(string(arquivo))

	// leitura de pouco em pouco abrindo o arquivo
	arquivo2, err := os.Open("arquivo.txt")
	if err != nill{
		panic(err)
	}
	reader := bufio.NewReader(arquivo2)
	buffer := nake([]byte,10)
	for{
		n, err := reader.Read(buffer)
		if err != nill{
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	err := os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}
}
//manipulação de arquivos