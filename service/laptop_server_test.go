package service_test

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorpc/pb"
	"gorpc/sample"
	"gorpc/service"
	"log"
	"testing"
)

func TestCreateLaptop(t *testing.T) {
	t.Parallel()

	laptopNoId := sample.NewLaptop()
	laptopNoId.Id = ""

	laptopInvalidId := sample.NewLaptop()
	laptopInvalidId.Id = "invalid-uuid"

	laptopDuplicateId := sample.NewLaptop()
	storeDuplicate := service.NewInMemLaptopStore()
	err := storeDuplicate.Save(laptopDuplicateId)
	if err != nil {
		log.Fatalln(err)
	}

	testCases := []struct {
		name   string
		laptop *pb.Laptop
		store  service.LaptopStore
		code   codes.Code
	}{
		{
			name:   "success with id",
			laptop: sample.NewLaptop(),
			store:  service.NewInMemLaptopStore(),
			code:   codes.OK,
		},
		{
			name:   "success with no id",
			laptop: laptopNoId,
			store:  service.NewInMemLaptopStore(),
			code:   codes.OK,
		},
		{
			name:   "failure with invalid id",
			laptop: laptopInvalidId,
			store:  service.NewInMemLaptopStore(),
			code:   codes.InvalidArgument,
		},
		{
			name:   "success with duplicate id",
			laptop: laptopDuplicateId,
			store:  storeDuplicate,
			code:   codes.AlreadyExists,
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			req := &pb.CreateLaptopRequest{Laptop: tc.laptop}
			server := service.NewLaptopServer(tc.store)
			res, err := server.CreateLaptop(context.Background(), req)
			if tc.code == codes.OK {
				if err != nil || res.Id == "" || res == nil {
					if len(tc.laptop.Id) > 0 && res.Id != tc.laptop.Id {
						log.Fatalln("Response id is not equal to request id")
					}
				}
			} else {
				st, _ := status.FromError(err)
				if st.Code() != tc.code {
					log.Println("Invalid status code returned", st.Code())
				}
			}

		})
	}
}
