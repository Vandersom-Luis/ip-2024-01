// Questão 14

package main
import "fmt"
import "math"

func main() {
    
    var n float64 
    
    fmt.Scan(&n)
    slice := make([]int, int(n))
    
    for i := range slice {
        fmt.Scan(&slice[i])
    }
    crescente(slice)
    
    
    
    if int(n) % 2 != 0 {
        fmt.Print(slice[int(math.Ceil((n / 2)-1))])
    } else {
        fmt.Print(float64(slice[int((n/2)-1)] + slice[int(n/2)] ) / 2)
    }
}

func crescente(array []int) {
    tamanho := len(array)
    for i := 0; i < tamanho-1; i++ {
        trocou := false
        for j := 0; j < tamanho-i-1; j++ {
            if array[j] > array[j+1] {
                array[j], array[j+1] = array[j+1], array[j]
                trocou = true
            }
        }
        if !trocou {
            break
        }
    }
}