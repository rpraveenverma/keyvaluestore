//TCPNewClient
package main

import (
    "bufio"
    "os"
    "encoding/gob"
    "fmt"
    "net"
)

func client() {
    // connect to the server
    c, err := net.Dial("tcp", "127.0.0.1:9999")
    if err != nil {
        fmt.Println(err)
        return
    }
    // send the message
    var msg string
    for{
    TakeInput(&msg)
    err = gob.NewEncoder(c).Encode(msg)
    if err != nil {
        fmt.Println(err)
    }
    }
    c.Close()
}

func TakeInput(msg *string) {
    bio := bufio.NewReader(os.Stdin)
    input, err := bio.ReadString('\n')

    if err != nil {
          panic(err)
        }
    *msg=input
}
func main() {
    client()
    var input string
    fmt.Scanln(&input)
}
