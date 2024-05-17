// Questão 06

package main
import "fmt"

func main() {
    var n, soma int

    fmt.Scan(&n)

    slice := make([]int, n)

    for i := range slice {
        fmt.Scan(&slice[i])
        soma += slice[i]
    }
    fmt.Print(soma)
}