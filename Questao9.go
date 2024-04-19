// Questão 9

package main
import "fmt"

func main() {
    var a, b, c, delta float64

    fmt.Scanf("%f\n%f\n%f", &a, &b, &c)
    delta = (b * b) - (4 * a * c)

    fmt.Printf("O VALOR DE DELTA É = %f", delta)

}