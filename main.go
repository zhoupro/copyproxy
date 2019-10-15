// Package main provides ...
package main

import (
    "fmt"
    "net"
    "os"
    "runtime"
    "os/exec"
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
    buf := make([]byte, 4096)
    _,err := conn.Read(buf)
    if err != nil{
        fmt.Println("error reading:", err.Error())
    }

    sysType := runtime.GOOS
    if sysType == "windows"{
        d := exec.Command("cmd","/c","echo",string(buf), "|","clip")
        d.Run()
    }

    conn.Write([]byte("Message received."))
}

