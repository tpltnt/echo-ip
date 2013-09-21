package main

import(
        "log"
        "net"
)


func echo(connection net.Conn) {
}


func main() {
     // open socket
     listener, err := net.Listen("tcp", ":1256")
     // log error & close
     if nil != err {
        log.Fatal(err)
     }
     defer listener.Close()

     // handle incoming connections
     for {
         connection, err := listener.Accept()
         if nil != err {
            log.Println(err)
         }

         go echo(connection)
     }
}