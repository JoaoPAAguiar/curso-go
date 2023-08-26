package main

func main() {
	for i := 0; i< 10; i++{
		println(i)
	}

	numeros := []string{"um", "dois","três"}
	for k, v := range numeros {
		println(k, v)
	}

	i := 0
	for i < 10{
		println(i)
		i++
	}

	for {
		println("Hello, World")
	}
}

func main02(){
	a := 1 
	b := 2
	c := 3

	switch a {
	case 1:
		println("a")
	case 2:
		println("b")
	case 3:
		println("c")
	default:
		println("d")
	}

	if a > b || c > a{
		println("a > b && c > a")
	}
}

//Compilando projetos
