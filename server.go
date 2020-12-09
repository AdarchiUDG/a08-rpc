package main

import (
	"net"
	"net/rpc"
	"fmt"
	"os"
	"./school"
)

func main() {
	sc := new(school.School)
	sc.Students = make(map[string]map[string]float64)
  rpc.Register(sc)

  tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
  if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}	

  listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
	
  for {
			conn, err := listener.Accept()
      if err != nil {
          continue
      }
      rpc.ServeConn(conn)
  }
}