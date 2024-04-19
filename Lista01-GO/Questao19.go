// Questão 19

package main
import "fmt"

func main() {
    var n int
    var somaTotal float64

    fmt.Scanf("%d", &n)

    if (n > 1 && n % 1 == 0) {
        for i := 1 ; i <= n ; i++ {
            somaTotal += 1 / float64(i)
        }
        fmt.Println(somaTotal)
    } else {fmt.Printf("Número inválido!")}
}