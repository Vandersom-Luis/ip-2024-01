// Questão 15

package main
import "fmt"

func main() {
    var n1, n2, z int
    
    fmt.Scan(&n2)
    menor := make([]int, n2)
    comparacoes := make([]int, n2)
    

    for i := range menor {
        menor[i] = 99999999
    }

    for x := 0 ; x < n2 ; x++ {
        z = 0
        
        fmt.Scan(&n1)
        slice := make([]int, n1)
    
        for i := range slice {
            fmt.Scan(&slice[i])
        }
    
        for _,v := range slice {
            for _,y := range slice {
                if y > v {
                    if y - v < menor[x] {
                        menor[x] = y - v
                    }
                }
                if v > y {
                    if v - y < menor[x] {
                        menor[x] = v - y
                    }
                }
                if v == y {
                    z++
                }
                if z > n1 {
                 menor[x] = 0
                }
            }
        }
        comparacoes[x] = (n1 * (n1-1)) / 2
    }
    for i := range menor {
        fmt.Println(menor[i], " ", comparacoes[i])
    }
}