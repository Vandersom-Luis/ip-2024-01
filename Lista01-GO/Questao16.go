// Questão 16

package main
import "fmt"

func main() {
    var salario, reajuste, salarioNovo float64

    fmt.Scanf("%f", &salario)

    if (salario <= 300) {
    reajuste = 0.5
    } else {
    reajuste = 0.3 }

    salarioNovo = salario + (salario * reajuste)

    fmt.Printf("Salário com reajuste = %f", salarioNovo)
}