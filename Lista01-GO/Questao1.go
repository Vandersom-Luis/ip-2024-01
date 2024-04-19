// Questão 1

package main

import "fmt"

func main() {
	var nota1, nota2, nota3, media float64

	fmt.Println("Digite três valores")
	fmt.Scanf("%f", &nota1)
	fmt.Scanf("%f", &nota2)
	fmt.Scanf("%f", &nota3)

	media = (nota1 + nota2 + nota3) / 3
	fmt.Println("MEDIA =", media)

	if media >= 6 {
	    fmt.Println("APROVADO")
	} else {
	    fmt.Println("REPROVADO")
	}
}