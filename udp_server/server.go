package main

import (
	"encoding/binary"
	"fmt"
	"net"
)


func main() {
    localAddr, err := net.ResolveUDPAddr("udp", "0.0.0.0:35912")
    if err != nil {
        fmt.Println("Couldn't resolve addr")
        return
    }

    conn, err := net.ListenUDP("udp", localAddr)
    if err != nil {
        fmt.Println("Failed listening to local IP", err)
        return
    }

    defer conn.Close()

    wandaAddr, err := net.ResolveUDPAddr("udp", "192.168.1.10:35912")
    if err != nil {
        fmt.Println("Failed listening to Remote IP", err)
        return
    }

    message := []byte("Hello")
    _, err = conn.WriteToUDP(message, wandaAddr)
    if err != nil {
        fmt.Println("Error sending UDP:", err)
    }

    fmt.Println("Sent packet")

    buffer := make([]byte, 4)
    for {
        _, addr, err := conn.ReadFromUDP(buffer)
        if err != nil {
            fmt.Println("error reading udp")
            continue
        }
        fmt.Printf("%d from %s\n", binary.LittleEndian.Uint32(buffer), addr)
    }

}
