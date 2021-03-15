package client

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("Hello User!")

	conn, err := grpc.Dial("localhost:4000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()
}

