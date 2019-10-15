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
    reqLen,err := conn.Read(buf)
    if err != nil{
        fmt.Println("error reading:", err.Error())
    }

    sysType := runtime.GOOS
    if sysType == "windows"{
        fmt.Println(string(buf[:reqLen]))
        wincopy(string(buf[:reqLen]))
    }

    conn.Write([]byte("Message received."))
}

func wincopy(str string) {
    cmd1 := exec.Command("cmd", "/c","echo", str)
    cmd2 := exec.Command("cmd", "/c","clip")
    cmd2.Stdout = os.Stdout
    in, _ := cmd2.StdinPipe()
    cmd1.Stdout = in
    cmd2.Start()
    cmd1.Run()
    in.Close()
    cmd2.Wait()
}
