//Questão 2

package main

import "fmt"

func main() {

    var nCasos int
    var nAtual int = 1
    var pPopular, pGeral, pArquibancada, pCadeiras, nPessoas, nPopular, nGeral, nArquibancada, nCadeiras, renda float64

	fmt.Scanf("%d", &nCasos)

	for (nAtual != nCasos + 1) {
	    fmt.Scanf("%f %f %f %f %f", &nPessoas, &pPopular, &pGeral, &pArquibancada, &pCadeiras)

	    nPopular = nPessoas * pPopular / 100
        nGeral = nPessoas * pGeral / 100
        nArquibancada = nPessoas * pArquibancada / 100
        nCadeiras = nPessoas * pCadeiras / 100
        renda = nPopular * 1 + nGeral * 5 + nArquibancada * 10 + nCadeiras * 20
        fmt.Printf("A RENDA DO JOGO N. %d E = %f\n", nAtual, renda)
        nAtual++ }
}