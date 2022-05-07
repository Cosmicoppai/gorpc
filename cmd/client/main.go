package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorpc/pb"
	"gorpc/sample"
	"log"
)

func main() {
	serverAddr := flag.String("address", "", "The Sever Address")
	flag.Parse()
	log.Printf("dial server %s \n", *serverAddr)

	conn, err := grpc.Dial(*serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	laptopClient := pb.NewLaptopServiceClient(conn)

	laptop := sample.NewLaptop()
	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}
	res, err := laptopClient.CreateLaptop(context.Background(), req)
	if err != nil {
		log.Println("Error while making request", err)
		return
	}
	log.Println("Created Laptop with id: ", res.Id)

}
