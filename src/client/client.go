package main

import (
	"log"

	"../person"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	serverAddress = "localhost:3000"
)

func main() {
	connection, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to server: %v", err)
	}
	defer connection.Close()
	client := person.NewPersonServiceClient(connection)

	v := &person.Void{}
	c := &person.ColorMessage{Color: "Blue"}

	pete, err := client.GetPerson(context.Background(), v)
	log.Printf("Pete: %v", pete)

	voidRet, err := client.DyeHair(context.Background(), c)
	log.Printf("Pete's hair color has changed. %v", voidRet)

	peteBlue, err := client.GetPerson(context.Background(), v)
	log.Printf("Pete: %v", peteBlue)
}
