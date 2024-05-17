// Questão 16

package main
import "fmt"

func main() {
    var n, k, npresentes int
    
    fmt.Scanf("%d %d", &n, &k)
    slice := make([]int, n)
    
    for i := range slice {
        fmt.Scan(&slice[i])
    }
    
    for _,v := range slice {
        if v <= 0 {
            npresentes++
        }
    }
    
    if npresentes < k {
        fmt.Println("SIM")
    } else {
        fmt.Println("NAO")
        
        for i := len(slice) - 1 ; i >= 0 ; i-- {
            if slice[i] <= 0 {
                fmt.Println(i+1)
            }
        }
    }
}