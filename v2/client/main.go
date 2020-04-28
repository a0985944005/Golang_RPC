package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
)

type Nums struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {
	//調用rpc server提供的方法之前，先與rpc server建立连接
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialHttp error", err)
		return
	}
	//"同步"調用Server端提供的方法

	nums := &Nums{7, 8}
	var reply int
	//可以查看原始碼 其實Call同步調用是用異步調用實現的
	err = client.Call("Arith.Multiply", nums, &reply)

	if err != nil {
		log.Fatal("call arith.Multiply error", err)
	}
	fmt.Printf("Arith:%d*%d=%d\n", nums.A, nums.B, reply)

	//"異步"調用Server端提供的方法

	quo := Quotient{}

	divCall := client.Go("Arith.Divide", nums, &quo, nil)

	//使用select模型監聽通道有數據時執行，若沒有則執行default
	for {
		select {
		case <-divCall.Done: //Struct Call 內的Done是chan
			fmt.Printf("商是%d,餘數是%d\n", quo.Quo, quo.Rem)
		default:
			fmt.Println("...")
			time.Sleep(time.Second * 1)
		}
	}

}
