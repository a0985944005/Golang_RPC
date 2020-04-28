package main

import (
	"log"
	"net"
	"net/rpc"
)

const HelloServiceName = "HelloWorld" //建立一個Service名稱  (client.server都需要)

type HelloServiceInterface = interface { //建一個interface (client.server都需要)
	Say(request string, reply *string) error
}

type HelloService struct{}

func RegisterHelloService(svc HelloServiceInterface) error { //建立RPC Service
	return rpc.RegisterName(HelloServiceName, svc)
}

func (p *HelloService) Say(request string, reply *string) error {
	*reply = "Hello" + request
	return nil
}

func main() {
	RegisterHelloService(new(HelloService))

	listener, err := net.Listen("tcp", ":5407")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		go rpc.ServeConn(conn)
	}
}
