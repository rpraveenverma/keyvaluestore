//TCPNewClient
package main

import (

    "bufio"
    "os"
    "fmt"
    "net"

)
func Client() {
    // connect to the server
    c, err := net.Dial("tcp", "127.0.0.1:9999")
    if err != nil{
        fmt.Println(err)
        return
    }
    // send the message
    var msg string
    for{
    takeInput(&msg)
    c.Write([]byte(msg))
    // reading response sent by server
    var buf [1024]byte
    n, _ :=c.Read(buf[0:])
    fmt.Println(string(buf[0:n]))
    }

    c.Close()
}
func takeInput(msg *string) {
    bio := bufio.NewReader(os.Stdin)
    input, err := bio.ReadString('\n')
    if err != nil {
          panic(err)
        }
    *msg=input
}
func main() {
    fmt.Println("Welcome To Key Value Store")
    fmt.Println("you can use---set-------get-------del-----and----close it--- commands to perform task")
    fmt.Println("====================================================================")
    client()
    var input string
    fmt.Scanln(&input)
}
