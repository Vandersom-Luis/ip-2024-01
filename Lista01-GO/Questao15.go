// Questão 15

package main
import "fmt"
import "math"

func main() {
    var n int
    fmt.Scanf("%d", &n)

    if (n > 5 && n < 2000) {
        for i := 1 ; i <= n ; i++ {
          if (i % 2 == 0) {
              fmt.Printf("%d^2 = %d\n", i, int(math.Pow(float64(i), 2)))
          }
        }
    } else {fmt.Printf("Número inválido")}
}