// Questão 18

package main
import "fmt"

func main() {
    var n, soma1, soma2 int

    fmt.Scan(&n)
    valido1 := make([]bool, n)
    valido2 := make([]bool, n)

    cpf := make([]int, 11)
    a := make([]int, 9)
    b := make([]int, 2)

    for x := 0 ; x < n ; x++ {

        for i := range cpf {
            fmt.Scan(&cpf[i])
        }

        for i := range a {
            a[i] = cpf[i]
        }
        b[0] = cpf[9]
        b[1] = cpf[10]

        soma1 = 0
        for i := range a {
            soma1 += a[i] * (i + 1)

        }
        if soma1 % 11 == 10 || soma1 % 11 == 0 {
            if b[0] == 0 {
                valido1[x] = true
            } else {valido1[x] = false}
        } else {
            if soma1 % 11 == b[0] {
                valido1[x] = true
            } else {valido1[x] = false}
        }

        soma2 = 0
        for i := range a {
            soma2 += a[i] * (9 - i)
        }
        if soma2 % 11 == 10 || soma2 % 11 == 0 {
            if b[1] == 0 {
                valido2[x] = true
            } else {valido2[x] = false}
        } else {
            if soma2 % 11 == b[1] {
                valido2[x] = true
            } else {valido2[x] = false}
        }
    }

    for i := range valido1 {
        if valido1[i] == true && valido2[i] == true {
            fmt.Println("CPF valido")
        } else { fmt.Println("CPF invalido") }
    }
}