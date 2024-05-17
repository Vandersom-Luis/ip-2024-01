// Questão 11

package main
import "fmt"

func main() {
    var n, maior int
    var menor int = 9999999999
    
    fmt.Scan(&n)
    slice := make([]int, n)
    invertida := make([]int, n)
    
    for i := range slice {
        fmt.Scan(&slice[i])
    }
    
    for i := range invertida {
        invertida[i] = slice[len(slice) - i - 1]
    }
    
    for i := range slice {
        fmt.Print(slice[i])
        if i != len(slice) - 1 {
            fmt.Print(" ")
        }
    }
    fmt.Print("\n")
    
    for i := range invertida {
        fmt.Print(invertida[i])
                if i != len(slice) - 1 {
            fmt.Print(" ")
        }
    }
    fmt.Print("\n")
    
    for _,v := range slice {
        if v > maior {
            maior = v
        }
    }
    for _,v := range invertida {
        if v < menor {
            menor = v
        }
    }
    fmt.Print(maior, "\n", menor)
}