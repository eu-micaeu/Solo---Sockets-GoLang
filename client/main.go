package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
    "strings"
)

func main() {

    serverIP := "0.0.0.0" // Substitua pelo endereço IP do servidor que deseja acessar

    // Conectar ao servidor no endereço IP e porta 8080
    conn, err := net.Dial("tcp", serverIP+":8080")
    if err != nil {
        fmt.Println("Erro ao conectar ao servidor:", err)
        os.Exit(1)
    }
    defer conn.Close()

    fmt.Println("Conectado ao servidor.")

    reader := bufio.NewReader(os.Stdin)
    for {
        fmt.Print("Digite uma mensagem: ")
        message, _ := reader.ReadString('\n')
        message = strings.TrimSpace(message)

        _, err := conn.Write([]byte(message + "\n"))
        if err != nil {
            fmt.Println("Erro ao enviar mensagem para o servidor:", err)
            return
        }
    }
}
