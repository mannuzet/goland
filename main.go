package main

import (
    "fmt"
)
func dadosPessoa(nome string, idade int) (int, string) {
    // string
    if idade >= 18 {
        return idade, "Maior de idade"
    }
    return idade, "Menor de idade"
}

func main() {
    // string
    nome := "João"
    idade := 20

    // string
    idadeCalculada, classificacao := dadosPessoa(nome, idade)

    // string
    fmt.Printf("Nome: %s\n", nome)
    fmt.Printf("Idade: %d\n", idadeCalculada)
    fmt.Println(classificacao)
}