// Questão 07

package main
import "fmt"

func main() {
    var n, max, frequencia, x int
    var ok bool = true
    resultadosx := make([]int, 0)
    resultadosf := make([]int, 0)

    for (ok) {

        max = 0
        frequencia = 0

        fmt.Scan(&n)

        if n == 0  {
            break
        }

        slice := make([]int, n)

        for i := range slice {
            fmt.Scan(&slice[i])
        }

        for _,v := range slice {
            if v > max {
                max = v
            }
        }
        for x = 0 ; x <= max ; x++ {
            for _,v := range slice {
                if x == v {
                    frequencia += 1
                }
        }
        resultadosx = append(resultadosx, x)
        resultadosf = append(resultadosf, frequencia)
    }
  }
  for i := range resultadosx {
  fmt.Printf("(%d) %d\n", resultadosx[i], resultadosf[i])
  }
}