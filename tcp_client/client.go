package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

func main() {
    wandaAddress := "192.168.1.10:35912"


    conn, err := net.Dial("tcp", wandaAddress)

    if err != nil {
        fmt.Println("Failed to connect: ", err)
        return
    }

    defer conn.Close()


    for {
        fmt.Print("Enter an integer: ")

        var command uint32
        _, err := fmt.Scanf("%d", &command)
        if err != nil {
            fmt.Println("Failed to grab integer.")
        }

        serializedCommand := make([]byte, 4)
        binary.LittleEndian.PutUint32(serializedCommand, command)

        _, err = conn.Write(serializedCommand)

        if err != nil {
            fmt.Println("Failed to send command: ", err)
        }

        responseBytes := make([]byte, 4)
        _, err = conn.Read(responseBytes)
        if err != nil {
            fmt.Println("Bad response: ", err)
        }

        responseValue := binary.LittleEndian.Uint32(responseBytes)
        fmt.Println("Response from server: ", responseValue)
    }


}
