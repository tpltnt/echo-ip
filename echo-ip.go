package main

import(
        "log"
        "net"
)

func main() {
     // open socket
     listener, err := net.Listen("tcp", ":1256")
     // log error & close
     if nil != err {
        log.Fatal(err)
     }
     defer listener.Close()

     for {
         connection, err := listener.Accept()
         if nil != err {
            log.Println(err)
         }
     }
}