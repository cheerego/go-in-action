package main

import (
	"context"
	"github.com/cheerego/go-micro-in-action/consul/proto"
	"github.com/micro/go-micro"
	"log"
)

type Greeter struct{}

func (s *Greeter) Hello(ctx context.Context, req *greeter.Request, rsp *greeter.Response) error {
	log.Print("Received Say.Hello request")
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.greeter"),
	)

	// optionally setup command line usage
	service.Init()

	// Register Handlers
	greeter.RegisterGreeterHandler(service.Server(), new(Greeter))

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
