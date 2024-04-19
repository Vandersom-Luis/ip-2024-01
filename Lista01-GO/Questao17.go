// Questão 17

package main
import "fmt"

func main() {
    var x, y int

    fmt.Scanf("%d %d", &x, &y)

    if (x % 2 == 0) {
        for i := 0 ; i < y ; i++ {
            fmt.Printf("%d ", x + (i * 2))
        }
    } else {
        fmt.Printf("O PRIMEIRO NÚMERO NÃO É PAR")
    }
    fmt.Printf(" \n")
}