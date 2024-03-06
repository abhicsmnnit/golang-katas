package main

import (
	"fmt"

	"github.com/abhicsmnnit/golang-katas/11/generate_proto_example/data"
	"google.golang.org/protobuf/proto"
)

//go:generate protoc -I=. --go_out=. --go_opt=module=github.com/learning-go-book-2e/proto_generate --go_opt=Mperson.proto=github.com/learning-go-book-2e/proto_generate/data person.proto
func main() {
	p := &data.Person{
		Name:  "Abhinav Verma",
		Id:    1,
		Email: "abhicsmnnit@gmail.com",
	}
	fmt.Println(p)

	protoBytes, _ := proto.Marshal(p)
	fmt.Println(protoBytes)

	var person data.Person
	proto.Unmarshal(protoBytes, &person)
	fmt.Println(person)
}
