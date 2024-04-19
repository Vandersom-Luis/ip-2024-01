// Questão 20

package main
import "fmt"

func main() {
    var horas, minutos, segundos, totalSegundos int

    fmt.Scanf("%d\n%d\n%d", &horas, &minutos, &segundos)

    totalSegundos = (horas * 3600) + (minutos * 60) + segundos

    fmt.Printf("O TEMPO EM SEGUNDOS É = %d", totalSegundos)
}