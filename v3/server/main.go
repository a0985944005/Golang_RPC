package main

import (
	"errors"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"time"
)

type Nums struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

//乘法計算
func (t *Arith) Multiply(nums *Nums, reply *int) error {
	time.Sleep(time.Second * 3) //dealy三秒 同步會等待 異步會繼續執行下面的
	*reply = nums.A * nums.B
	return nil
}

//商數和餘數
func (t *Arith) Divide(nums *Nums, quo *Quotient) error {
	time.Sleep(time.Second * 3)
	if nums.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = nums.A / nums.B
	quo.Rem = nums.A % nums.B
	return nil
}

func main() {
	//創建對象
	arith := new(Arith)
	//rpc server註冊了一个arith對象 公開方法讓client調用
	rpc.Register(arith)
	//指定rpc的傳輸協議 這裡採用http協議作為rpc調用的方法 也可以用rpc.ServeConn處理單個連接請求
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error", e)
	}
	go http.Serve(l, nil)
	os.Stdin.Read(make([]byte, 1))
}
