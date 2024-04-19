// Questão 13

package main
import "fmt"

func main() {
    var nota float64
    var conceito string

    fmt.Scanf("%f", &nota)

    if (nota < 6) {
        conceito = "D"
    } else if (nota >= 6 && nota < 7.5) {
        conceito = "C"
    } else if (nota >= 7.5 && nota < 9) {
        conceito = "B"
    } else {
        conceito = "A"
    }

    fmt.Printf("NOTA = %f CONCEITO = %s", nota, conceito)
}