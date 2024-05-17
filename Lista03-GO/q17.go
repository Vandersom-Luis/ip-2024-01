// Questão 17

package main
import "fmt"

func main() {
    var n, um, repete int
    fmt.Scan(&n)

    slice := make([]int, n)

    for i := range slice {
        fmt.Scan(&slice[i])
    }

    for _,v := range slice {
        repete = 0
        for _,y := range slice {
            if y == v {
                repete++
            }
        }
       if repete == 1 {
           um++
       }
    }
    fmt.Print(um)
}