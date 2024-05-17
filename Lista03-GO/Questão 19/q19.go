// Questão 19

package main
import "fmt"

func main() {
    var n int  
    
    fmt.Scan(&n)
    slice := make([]int, n)
    unico := make([]int, 0)
    
    for i := range slice {
        fmt.Scan(&slice[i])
    }
    
    for _,v := range slice {
        if v != 0 {
            unico = append(unico, v)
            for x,y := range slice {
                if y == v {
                    slice[x] = 0
                }
            }
        }
    }
	
    for i := range unico {
        fmt.Println(unico[i])
    }
}