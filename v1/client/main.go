package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:5407") //和RPC Server連線

	if err != nil {
		log.Fatal("dialing", err)
	}

	var reply string
	err = client.Call("HelloService.Say", "World!!", &reply) //client.Call(server.func,parameter1,&parameter2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply) // return result
}
