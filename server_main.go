package main

import (
	pb "Go_SIMPLE_SOCKET_STUDY/pb"
	"fmt"
	"net"
	"time"

	"github.com/golang/protobuf/proto"
)

func main() {

	fmt.Printf("Qx, GSS run. %s\n", time.Now())

	service := ":10001"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)

	checkErr(err)

	listen, err := net.ListenTCP("tcp", tcpAddr)
	checkErr(err)

	for {
		conn, err := listen.Accept()

		if err != nil {
			continue
		}

		t := &pb.Test{
			Label: "123",
			Type:  123,
		}

		data, err := proto.Marshal(t)

		if err != nil {
			fmt.Printf("Marshal error\n")
		}

		msg := &pb.Test{}

		err = proto.Unmarshal(data, msg)

		if err == nil {
			fmt.Printf("Label is %s\n", msg.Label)
			fmt.Printf("Type is %d\n", msg.Type)
		} else {
			fmt.Printf("Unmarshal error\n")
		}

		conn.Write(data)

		fmt.Printf("data is %b\n", data)
		fmt.Printf("send data...\n")
	}

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
