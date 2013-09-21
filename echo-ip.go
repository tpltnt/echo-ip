package main

import(
        "log"
        "net"
)

func main() {
     // open socket
     listensocket, err := net.Listen("tcp", ":1256")
     // log error & close
     if err != nil {
        log.Fatal(err);
     }
}