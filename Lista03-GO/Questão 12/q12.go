// Questão 12

package main
import "fmt"

func main() {
    var n int
    
    fmt.Scan(&n)
    slice := make([]int, n)
    
    for i := range slice {
        fmt.Scan(&slice[i])
    }
    
    crescente(slice)
    
    for i := range slice {
        fmt.Println(slice[i])
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