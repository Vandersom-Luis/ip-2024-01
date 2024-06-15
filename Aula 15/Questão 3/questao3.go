// Questão 3

package main
import "fmt"

func reverter(array []int, começo int, final int) []int {
    
    if começo >= final {
        return array
    }
    
    array[começo], array[final] = array[final], array[começo]
    
    reverter(array, começo + 1, final - 1)
    return array
}


func main() {
    array := []int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    fmt.Print(reverter(array, 0, len(array) - 1))
}