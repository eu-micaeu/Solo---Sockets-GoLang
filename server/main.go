package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
)

func main() {
	
    // Iniciar o servidor na porta 8080
    listener, err := net.Listen("tcp", "0.0.0.0:8080")
    if err != nil {
        fmt.Println("Erro ao iniciar o servidor:", err)
        os.Exit(1)
    }
    defer listener.Close()

    fmt.Println("Servidor iniciado e ouvindo na porta 8080")

    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Erro ao aceitar a conexão do cliente:", err)
            continue
        }

        go handleClient(conn)
    }
}

func handleClient(conn net.Conn) {
    defer conn.Close()

    fmt.Println("Cliente conectado:", conn.RemoteAddr())

    reader := bufio.NewReader(conn)
    for {
        message, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("Erro ao ler mensagem do cliente:", err)
            return
        }

        fmt.Printf("Mensagem do cliente [%s]: %s", conn.RemoteAddr(), message)
    }
}
