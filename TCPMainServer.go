//TCPMainServer
package main
import (
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
    var buf [1024]byte
    var msg string
    for{
    num, _ :=c.Read(buf[0:])
    msg = string(buf[0:num])
    TakeDecision(msg,c)
  }
    c.Close()
}
func TakeDecision(msg string,c net.Conn){
    var cmd, key ,value string
    arr := strings.Split(msg," ")
    var flag int
    flag=0
    if( (len(arr) ==2 && arr[0]=="set" ) || (len(arr)<2 || len(arr)>3)){
        c.Write([]byte("Bad Command"))
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
    case "set": set(key,value,c)
                break
    case "get": get(key,c)
                break
    case "del": del(key,c)
                break
    case "close":c.Close()
    default:c.Write([]byte("Bad Command"))
    }
   }
}

func set(key ,value string,c net.Conn){
    lock.Lock()
    kvstore[key]=value
    lock.Unlock()
    c.Write([]byte("--set"+key))

}

func get(key string,c net.Conn){

    key =strings.TrimSpace(strings.Trim(key, "\n"))
    lock.Lock()
    value, bit:=kvstore[key]
    lock.Unlock()
    if bit!=true {
        c.Write([]byte("--Does Not Exist"))
    }else{
        c.Write([]byte(value))
    }
}

func del(key string,c net.Conn){
        key =strings.TrimSpace(strings.Trim(key, "\n"))
        lock.Lock()
        delete(kvstore,key)
        lock.Unlock()
        c.Write([]byte("--"+key+" Deleted"))
}

func main() {
    fmt.Println("Server Started")
    kvstore= make(map[string]string)
    server()
    var input string
    fmt.Scanln(&input)
}
