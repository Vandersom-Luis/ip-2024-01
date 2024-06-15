// Questão 1

package main
import "fmt"

func elevado(x, n int) int {
    if n <= 0 {
        return 1
    }
    return x * elevado(x, n - 1)
}

func main() {
  fmt.Print(elevado(2, 3))
}