// Questão 2

package main
import "fmt"

func remove(slice []int, s int) []int {
    return append(slice[:s], slice[s+1:]...)
}

func soma(array []int) int {
    if len(array) == 1 {
        return array[0]
    }

    return array[len(array) - 1] + soma(remove(array, len(array) - 1))
}



func main() {
    array := []int{1, 2, 3, 4} 

    fmt.Print(soma(array))
}