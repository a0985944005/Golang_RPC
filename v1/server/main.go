package main

import (
	"log"
	"net"
	"net/rpc"
)

type HelloService struct{} //宣告一個RPC使用的struct

//在RPC的規定中，輸入的參數只能有兩個且第二個必須是指標參數，回傳的變數至少要有一個error
func (p *HelloService) Say(request string, reply *string) error {
	*reply = "Hello " + request
	return nil
}

func main() {
	rpc.RegisterName("HelloService", new(HelloService))

	listener, err := net.Listen("tcp", ":5407") //port前面要記得+:
	if err != nil {
		log.Fatal("Accept error:", err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}

	rpc.ServeConn(conn) //rpc.ServeConn函數在該TCP鏈接上為對方提供RPC服務
}
