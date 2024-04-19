// Questão 14

package main
import "fmt"
import "math"

func main() {
    var altura, arestaBase, areaBase, volume float64

    fmt.Scanf("%f\n%f", &altura, &arestaBase)

    areaBase = 3 * arestaBase * arestaBase * math.Sqrt(3) / 2
    volume = areaBase * altura / 3

    fmt.Printf("O VOLUME DA PIRÂMIDE É = %f METROS CÚBICOS", volume)
}