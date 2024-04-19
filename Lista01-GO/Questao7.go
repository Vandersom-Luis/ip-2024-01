// Questão 7

package main
import "fmt"

func main() {
  var farenheit, pChuva, celsius, mmChuva float64

  fmt.Scanf("%f\n%f", &farenheit, &pChuva)

  celsius = (5 * (farenheit - 32)) / 9
  mmChuva = pChuva * 25.4

  fmt.Printf("O VALOR EM CELSIUS E = %f\nA QUANTIDADE DE CHUVA É = %f", celsius, mmChuva)
}