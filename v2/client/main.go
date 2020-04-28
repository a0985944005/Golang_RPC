package main

import (
	"log"
	"net/rpc"
)

const HelloServiceName = "HelloWorld" //建立一個Service名稱  (client.server都需要)

type HelloServiceInterface = interface { //建一個interface (client.server都需要)
	Say(request string, reply *string) error
}

type HelloServiceClient struct {
	*rpc.Client
}

var _ HelloServiceInterface = (*HelloServiceClient)(nil) //這行寫法是在測試某個struct `HelloServiceClient` 是否實現interface `HelloServiceInterface`的時候用的

func (p *HelloServiceClient) Say(request string, reply *string) error { //client端的struct實現了interface中的Say()
	return p.Client.Call(HelloServiceName+".Say", request, reply) //Call server端的func
}

func DialService(network, address string) (*HelloServiceClient, error) { //建立RPC連線
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: c}, nil
}

func main() {
	client, err := DialService("tcp", "localhost:5407")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string

	err = client.Say(" World!!", &reply)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(reply)
}
