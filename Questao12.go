// Questão 12

package main
import "fmt"

func main() {
    var horas, precoUm, precoDois, precoTotal, divisao int
    fmt.Scanf("%d", &horas)

    divisao = horas / 3
    precoUm = divisao * 10
    precoDois =  (horas - (divisao * 3)) * 5
    precoTotal = precoUm + precoDois

    fmt.Printf("O valor a pagar é = %d", precoTotal)
}