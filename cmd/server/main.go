package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"gorpc/pb"
	"gorpc/service"
	"log"
	"net"
)

func main() {
	port := flag.Int("port", 0, "The Server port")
	flag.Parse()
	log.Printf("Server is running on %d \n", *port)

	laptopServer := service.NewLaptopServer(service.NewInMemLaptopStore())
	grpcServer := grpc.NewServer()
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	address := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalln("Cannot start server:", err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalln(err)
	}

}
