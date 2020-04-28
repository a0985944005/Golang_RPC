# Golang_RPC
## 重構 hello world RPC!
在涉及RPC的應用中，作為開發人員一般至少有三種角色：
1.服務端實現RPC方法的開發人員
2.客戶端調用RPC方法的人員
3.(最重要)製定服務端和客戶端RPC接口規範的設計人員


重構HelloService服務，第一步需要明確服務的名字和接口：
```
const HelloServiceName = "path/to/pkg.HelloService"

type HelloServiceInterface = interface {
    Say(request string, reply *string) error
}

func RegisterHelloService(svc HelloServiceInterface) error {
    return rpc.RegisterName(HelloServiceName, svc)
}
```
首先是服務的名稱
然後是服務要實現的詳細方法列表
最後是註冊該類型服務的函數。
test.go ->  範例struct有無調用interface方法的內容