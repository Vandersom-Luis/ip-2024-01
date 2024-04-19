// Questão 18

package main
import "fmt"

func main() {
    var a, r, n, somaTotal int

    fmt.Scanf("%d %d %d", &a, &r, &n)

    for i := 0 ; i < n ; i++ {
        somaTotal += a + (r * i)
    }
    fmt.Println(somaTotal)
}