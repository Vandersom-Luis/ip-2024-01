// Questão 6

package main
import "fmt"

func main() {
    var nEntradas int
    var farenheit, celsius float64
    fmt.Scanf("%d", &nEntradas)

    for (nEntradas > 0) {
        fmt.Scanf("%f", &farenheit)
        celsius = (5 * (farenheit - 32)) / 9
        fmt.Printf("%f FAHRENHEIT EQUIVALE A %f CELSIUS\n", farenheit, celsius)
    }
}
