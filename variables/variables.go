package main

import "fmt"

func main() {
	// caso a variável não serja "instanciada" com valor, será atribuida a ela um valor padrão
	// se a variável é criada, ela terá que ser usada
	var teste string = "teste"
	// go possui inferência de variáveis
	var teste2 = "teste2"
	teste3 := "teste3"

	fmt.Println("Forma 1:", teste, ". Forma 2:", teste2, ". Forma 3:", teste3)
}
