package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorpc/pb"
	"log"
)

// LaptopServer is the server that provides laptop services
type LaptopServer struct {
	pb.UnimplementedLaptopServiceServer
	Store LaptopStore
}

func NewLaptopServer(store LaptopStore) *LaptopServer {
	return &LaptopServer{Store: store}
}

// CreateLaptop is Unary RPC to create a new laptop
func (server *LaptopServer) CreateLaptop(ctx context.Context, req *pb.CreateLaptopRequest) (*pb.CreateLaptopResponse, error) {
	laptop := req.GetLaptop()
	log.Printf("Receive a create-laptop request with id: %s", laptop.Id)

	if len(laptop.Id) > 0 {
		_, err := uuid.Parse(laptop.Id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "Laptop Id is not a valid UUID")
		}
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal Server Error")
		}
		laptop.Id = id.String()
	}
	// save the laptop
	if err := server.Store.Save(laptop); err != nil {
		if errors.Is(err, ErrAlreadyExist) {
			return nil, status.Errorf(codes.AlreadyExists, "Record already exists")
		}
		return nil, status.Errorf(codes.Internal, "Internal server error")
	}
	return &pb.CreateLaptopResponse{Id: laptop.Id}, nil
}
