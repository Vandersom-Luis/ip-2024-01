// Questão 4

package main
import "fmt"

func main() {
    var salarioMinimo, consumo float64
    fmt.Scanf("%f\n%f", &salarioMinimo, &consumo)
    var custo_por_kw float64 = 0.7 * salarioMinimo / 100
    var custo_do_consumo float64 = custo_por_kw * consumo
    var custo_com_desconto float64 = 0.9 * custo_do_consumo


    fmt.Printf("Custo por kW: R$ %f\nCusto do consumo: R$ %f\nCusto com desconto: R$ %f", custo_por_kw, custo_do_consumo, custo_com_desconto)
}