package main

import (
	"log"
	"net"
	"net/rpc"
)

const HelloServiceName = "path/to/pkg.HelloService"

type HelloServiceInterface interface {
	Hello(request string, reply *string) error
}

type HelloService struct{}

func (hs *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

func main() {
	var hs HelloServiceInterface
	hs = new(HelloService)
	err := RegisterHelloService(hs)
	if err != nil {
		return
	}
	listener, err := net.Listen("tcp", ":1235")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error: ", err)
	}
	rpc.ServeConn(conn)
}
