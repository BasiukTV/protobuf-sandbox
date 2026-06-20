package basic

import (
	"log"
	pb "my-protobuf/protogen/basic"
)

func BasicHello() {
	h := pb.Hello{
		Name: "World",
	}

	log.Println(&h)
}
