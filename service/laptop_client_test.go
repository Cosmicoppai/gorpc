package service_test

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorpc/pb"
	"gorpc/sample"
	"gorpc/serialize"
	"gorpc/service"
	"log"
	"net"
	"testing"
)

func TestClientCreateLaptop(t *testing.T) {
	t.Parallel()
	laptopServer, serverAddr := startLaptopServer(t)
	laptopClient := newLaptopClient(t, serverAddr)
	laptop := sample.NewLaptop()
	expectedId := laptop.Id
	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}
	res, err := laptopClient.CreateLaptop(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	if res == nil || res.Id != expectedId {
		t.Fatal("Invalid result", res)
	}
	other, err := laptopServer.Store.Find(res.Id)
	if err != nil || other == nil {
		t.Fatal(err)
	}

	reqSameLaptop(t, laptop, other)

}

func reqSameLaptop(t *testing.T, laptop *pb.Laptop, other *pb.Laptop) {
	json1, err := serialize.ProtoBufToJson(laptop)
	if err != nil {
		t.Fatal(err)
	}
	json2, err := serialize.ProtoBufToJson(other)
	if err != nil {
		t.Fatal(err)
	}
	if string(json2) != string(json1) {
		t.Fatal(err)
	}

}

func startLaptopServer(t *testing.T) (*service.LaptopServer, string) {
	laptopServer := service.NewLaptopServer(service.NewInMemLaptopStore())
	grpcServer := grpc.NewServer()
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Println(err)
	}
	go func() {
		err := grpcServer.Serve(listener)
		if err != nil {
			t.Error(err)
			return
		}
	}()
	return laptopServer, listener.Addr().String()

}

func newLaptopClient(t *testing.T, serverAddr string) pb.LaptopServiceClient {
	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatal(err)
	}
	return pb.NewLaptopServiceClient(conn)

}
