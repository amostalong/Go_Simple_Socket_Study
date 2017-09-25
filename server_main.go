package main

import (
	pb "Go_SIMPLE_SOCKET_STUDY/pb"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/golang/protobuf/proto"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func main() {

	fmt.Printf("Qx, GSS run. %s\n", time.Now())

	session, mgoErr := mgo.Dial("127.0.0.1:27017")

	defer session.Close()

	if mgoErr != nil {
		panic(mgoErr)
	}

	session.SetMode(mgo.Monotonic, true)

	type Person struct {
		Name     string "bson:'Name'"
		Age      string "bson:'Age'"
		Birthday string "bson:'Birthday'"
	}

	c := Person{}

	col := session.DB("prometheus").C("persion")

	insertErr := col.Insert(&Person{Name: "cq", Age: "1", Birthday: "20001010"})

	if insertErr != nil {
		log.Fatal(insertErr)
	}

	findErr := col.Find(bson.M{"name": "cq"}).One(&c)

	if findErr != nil {
		log.Fatal(findErr)
	} else {
		fmt.Printf("Have found.\n")
	}

	fmt.Println("age: ", c.Age)

	_, delErr := col.RemoveAll(bson.M{"name": "cq"})

	if delErr != nil {
		log.Fatal(delErr)
	}

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
