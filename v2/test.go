package main

type Reader interface {
	Read()
}

type Writer interface {
	Write()
}

type A struct {
}

func (a A) Read() {
}

var _ Reader = &A{} // 編譯通過，確保A實現 Reader() 接口
var _ Writer = &A{} // 編譯不通過，A沒有實現 Writer() 接口

func main() {

}
