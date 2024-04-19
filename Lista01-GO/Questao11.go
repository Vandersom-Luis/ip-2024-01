// Questão 11

package main
import "fmt"

func main() {
    var n int
    fmt.Scanf("%d", &n)

    if (n % 3 == 0) && (n % 5 == 0) {
        fmt.Printf("O NÚMERO É DIVISÍVEL\n")
    } else {
        fmt.Printf("O NÚMERO NÃO É DIVISÍVEL\n")
    }
}