package main

import(
        "log"
        "net"
)


// echo function
func echo(connection net.Conn) {
     // get remote address
     clientaddr := connection.RemoteAddr()
     // convert address string to bytes
     payload := []byte(clientaddr.String())
     // send them back
     connection.Write(payload)
     // clean up
     err := connection.Close()
     if nil != err {
        log.Println(err)
     }
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