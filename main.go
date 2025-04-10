package main

import (
    "fmt"
)

func main() {
    var idade int

    fmt.Print("Digite sua idade: ")
    _, err := fmt.Scan(&idade)

    if err != nil {
        fmt.Println("Erro ao ler a idade. Certifique-se de digitar um número.")
        return
    }

    if idade < 18 {
        fmt.Println("Você é menor de idade.")
    } else if idade <= 60 {
        fmt.Println("Você é adulto.")
    } else {
        fmt.Println("Você é idoso.")
    }
}