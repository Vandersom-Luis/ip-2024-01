// Questão 13

package main
import "fmt"

func main() {
    var n, frequencia, fmaior int
    var nmaior = 9999999999999
    
    fmt.Scan(&n)
    slice := make([]int, n)
    
    for i := range slice {
        fmt.Scan(&slice[i])
    }
    
    for _,v := range slice {
        frequencia = 0
        for _,y := range slice {
            if v == y {
                frequencia++
            }
        }
            if frequencia == fmaior {
                if v < nmaior {
                    nmaior = v
                }
            }
            
            if frequencia > fmaior{
            nmaior = v
            fmaior = frequencia
            }
    }
    fmt.Print(nmaior, " \n", fmaior)
}