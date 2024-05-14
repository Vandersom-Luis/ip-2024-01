// Questão 03

package main
import "fmt"

func main() {

    var n int

    fmt.Scan(&n)

    slice := make([]int, n)

    for i := 0 ; i < n ; i++ {
        fmt.Scan(&slice[i])
    }

    for i := range slice{
        fmt.Print(slice[len(slice) - i - 1], " ")
    }
}