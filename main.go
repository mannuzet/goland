package main

import (
    "fmt"
)

func main() {

    var saldo float64

    fmt.Print("Digite o saldo inicial: R$ ")
    fmt.Scan(&saldo)

    for {
        var comando string
        var valor float64

        fmt.Printf("\nSaldo atual: R$ %.2f\n", saldo)

        fmt.Print("\nDigite o comando (saque, deposito ou sair): ")
        fmt.Scan(&comando)

        if comando == "sair" {
            fmt.Println("Saindo do programa...")
            break
        }

        fmt.Print("Digite o valor: R$ ")
        fmt.Scan(&valor)

        if comando == "saque" {
            if valor > saldo {
                fmt.Println("Saldo insuficiente para realizar o saque.")
            } else {
                saldo -= valor
                fmt.Printf("Saque de R$ %.2f realizado com sucesso.\n", valor)
            }
        } else if comando == "deposito" {
            saldo += valor
            fmt.Printf("Depósito de R$ %.2f realizado com sucesso.\n", valor)
        } else {
            fmt.Println("Comando inválido. Tente novamente.")
        }
    }
}