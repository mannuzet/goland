package main

import (
    "fmt"
)

func main() {
    var balance float32
    fmt.Println("Insira o valor inicial:")
    fmt.Scan(&balance)
    escolherOperacao(&balance)
}

func escolherOperacao(balance *float32) {
    var option int
    fmt.Println("Do you want to withdraw (1), deposit (2) or close the session (3)?")
    fmt.Scan(&option)
    switch option{ 
 
    case 1:
        sacar(balance)
    case 2:
        depositar(balance)
    case 3:
        encerrar()
    default:
        fmt.Println("Invalid option")
    }
}

func depositar(balance *float32) {
    var value float32
    fmt.Println("what amount do you want to deposit?")
    fmt.Scan(&value)
    *balance += value
    fmt.Println("Seu saldo atual é:", *balance)
}

func sacar(balance *float32) {
    var value float32
    fmt.Println("How much do you want to withdraw?")
    fmt.Scan(&value)
    if value > *balance {
        fmt.Println("Saldo insuficiente.")
    } else {
        *balance -= value
        fmt.Println("Seu saldo atual é:", *balance)
    }
}

func encerrar() {
    fmt.Println("Session closed")
}