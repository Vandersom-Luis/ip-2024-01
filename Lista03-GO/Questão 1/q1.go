//Questão 01
package main
import "fmt"

func main() {

    var n, n2, x, igual int
    fmt.Scan(&n)
    numeros := make([]int, n)

    for i := 0 ; i < n ; i++ {
     fmt.Scan(&numeros[i])
    }

    fmt.Scan(&n2)
    comparacao := make ([]string, n2)

    for i := 0 ; i < n2 ; i++ {
        fmt.Scan(&x)

        for _,v := range numeros{
            if x == v {
                igual += 1
            }
        }
            if igual > 0 {
                comparacao[i] = "ACHEI"
                igual = 0
        } else {
            comparacao[i] = "NAO ACHEI"
        }

    }

    for _,v := range comparacao{
        fmt.Printf("%s\n", v)
    }
}