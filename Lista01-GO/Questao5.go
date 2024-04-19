// Questão 5

package main
import "fmt"

func main() {
    var conta int
    var consumo, valor float32
    var tipo string

    fmt.Scanf("%d %f %s", &conta, &consumo, &tipo)

    if (tipo == "R") {
        valor = 5 + 0.05 * consumo
    } else if (tipo == "C") {
        valor = 500  + (0.25 * (consumo - 80))
    } else if (tipo == "I") {
        valor = 800 + (0.04 * (consumo - 100))
    }

    fmt.Printf("CONTA = %d\nVALOR DA CONTA = %f", conta, valor)
}