// Questão 3

package main
import "fmt"

func main() {
    var n1, n2, n3, nconcatenado int

    fmt.Scanf("%d", &n1)
    fmt.Scanf("%d", &n2)
    fmt.Scanf("%d", &n3)

    if (n1 / 10 >= 1 || n2 / 10 >= 1 || n3 / 10 >= 1){
        fmt.Printf("DIGITO INVALIDO")
    } else {
    nconcatenado = n1 * 100 + n2 * 10 + n3
    fmt.Printf("%d, %d", nconcatenado, nconcatenado * nconcatenado)
    }
}