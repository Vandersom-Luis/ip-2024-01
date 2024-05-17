// Questão 02

package main
import "fmt"

var n, k, nmaior int

func n1() int {
    fmt.Scan(&n)

    if n < 1 || n > 1000 {
        n1()
    }
    return n
}

func main() {
    v := make([]int, n1())

    for i := 0 ; i < n ; i++ {
        fmt.Scan(&v[i])
    }

    fmt.Scan(&k)

    for _,v := range v {
        if k <= v {
            nmaior += 1
        }
    }
    fmt.Print(nmaior)
}