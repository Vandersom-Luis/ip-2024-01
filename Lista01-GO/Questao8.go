// Questão 8

package main
import "fmt"

func main() {
    var raio, altura, pi, aLateral, aCirculo, aTotal, custo float64

    fmt.Scanf("%f\n%f", &raio, &altura)

    pi = 3.14159
    aLateral = 2 * pi * raio * altura
    aCirculo = pi * raio * raio
    aTotal = 2 * aCirculo + aLateral
    custo = 100 * aTotal

    fmt.Printf("O VALOR DO CUSTO E = %f", custo)

}