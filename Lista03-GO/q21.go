// Questão 21

package main
import "fmt"

func main() {

    var n int
    fmt.Scan(&n)

    slice := make([]int, n)
    pares := make([]int, 0)
    impares := make([]int, 0)

    for i := range slice {
        fmt.Scan(&slice[i])
    }

    for _,v := range slice {
        if v % 2 == 0 {
            pares = append(pares, v)
        } else { impares = append(impares, v)}
    }
    crescente(pares)
    decrescente(impares)

    for _,v := range pares {
        fmt.Println(v)
    }

    for _,v := range impares {
        fmt.Println(v)
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

func decrescente(array []int) {
    tamanho := len(array)
    for i := 0; i < tamanho-1; i++ {
        trocou := false
        for j := 0; j < tamanho-i-1; j++ {
            if array[j] < array[j+1] {
                array[j], array[j+1] = array[j+1], array[j]
                trocou = true
            }
        }
        if !trocou {
            break
        }
    }
}