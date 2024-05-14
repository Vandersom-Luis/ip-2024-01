// Questão 05

package main
import "fmt"

func main() {
    var n, max, pos int
    var ok bool = true
    resultados := make([]int, 0)
    resultadosPos := make([]int, 0)

    for (ok) {
    fmt.Scan(&n)

    if n == 0 {
        break
    }

    slice := make([]int, n)

    for i := range slice {
    fmt.Scan(&slice[i])}

    max = 0
    pos = 0

    for i,v := range slice {
        if v > max {
            max = v
            pos = i
        }
    }

    resultadosPos = append(resultadosPos, pos)
    resultados = append(resultados, max)
    }
    for i := range resultados {
     fmt.Println(resultadosPos[i], resultados[i])
    }
}