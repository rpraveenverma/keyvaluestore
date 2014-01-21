//TCPMainServer
package main
import (
    "encoding/gob"
    "fmt"
    "net"
    "strings"
    "sync"
)

var kvstore map[string]string
var connection net.Conn
var lock sync.Mutex

func server() {
    // listen on a port
    ln, err := net.Listen("tcp", ":9999")
    if err != nil {
        fmt.Println(err)
        return
    }
    for {
        // accept a connection
        c, err := ln.Accept()
        connection=c
        if err != nil {
            fmt.Println(err)
            continue
        }
        // handle the connection
        go handleServerConnection(c)
    }
}

func handleServerConnection(c net.Conn) {
    // receive the message

    for {
    var msg string
    err := gob.NewDecoder(c).Decode(&msg)
    if err != nil {
    } else {
        //print("command is: ",msg)
        TakeDecision(msg)
    }
   }
    c.Close()
}
func TakeDecision(msg string){
    var cmd, key ,value string
    arr := strings.Split(msg," ")
    var flag int
    flag=0
    if( (len(arr) ==2 && arr[0]=="set" ) || (len(arr)<2 || len(arr)>3)){
        fmt.Println("Bad Command")
    }else if len(arr)==2{
    flag=1
    cmd=arr[0]
    key=arr[1]
    }else{
    flag=2
    cmd=arr[0]
    key=arr[1]
    value=arr[2]
    }
    if flag!=0{
    switch cmd{
    case "set": set(key,value)
                break
    case "get": get(key)
                break
    case "del": del(key)
                break
    default: fmt.Println("Bad Command")
    }
   }

}

func set(key ,value string){
    lock.Lock()
    kvstore[key]=value
    lock.Unlock()
    fmt.Println("New hashmap is \n")
    fmt.Println(kvstore)
}

func get(key string){
    key =strings.TrimSpace(strings.Trim(key, "\n"))
    fmt.Println(kvstore[key])
}

func del(key string){
        key =strings.TrimSpace(strings.Trim(key, "\n"))
        lock.Lock()
        delete(kvstore,key)
        lock.Unlock()
        fmt.Println(kvstore)
}

func sendtoclient(){

}

func main() {
    kvstore= make(map[string]string)
    server()
    var input string
    fmt.Scanln(&input)
}
