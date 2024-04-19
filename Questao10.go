// Questão 10

package main
import "fmt"

func main() {
    var a, b, c, d, determinante float64

    fmt.Scanf("%f\n%f\n%f\n%f", &a, &b, &c, &d)

    determinante = (a * d) - (b * c)

    fmt.Printf("O VALOR DO DETERMINANTE É = %f", determinante)
}