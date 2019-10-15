// Package main provides ...
package main

import (
    "fmt"
    "net"
    "os"
    "os/exec"
    "io"
)

const (
    CONN_HOST  string = "0.0.0.0"
    CONN_PORT string = "8899"
    CONN_TYPE string = "tcp"
)

func main() {
    // Listen
    l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
    if err != nil{
        fmt.Println("Error listening:", err.Error())
        os.Exit(1)
    }
    defer l.Close()
    fmt.Println("Listening on "+CONN_HOST+":"+CONN_PORT)

    for{
        // listen for an incoming connection
        conn, err := l.Accept()
        if err != nil{
            fmt.Println("Error accetping:", err.Error())
            os.Exit(1)
        }
        go handleRequest(conn)
    }
}

func handleRequest(conn net.Conn) {
    defer conn.Close()
    cmd := exec.Command("cmd", "/c","clip")
    stdin, err := cmd.StdinPipe()
    if err != nil {
        fmt.Println("Error pipe init:", err)
        os.Exit(1)
    }
    if  err = cmd.Start();err != nil {
        fmt.Println("Error process start:", err)
        os.Exit(1)
    }
    if copied ,err := io.Copy(stdin,conn); err !=nil{
        fmt.Println("Error pipe copy:", err)
        os.Exit(1)
    }else{
        fmt.Println("copied:", copied)
    }
    stdin.Close()
    if err = cmd.Wait();err!=nil{
        fmt.Println("Error wait:", err)
    }
}
