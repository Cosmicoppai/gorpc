package serialize_test

import (
	"google.golang.org/protobuf/proto"
	"gorpc/sample"
	"gorpc/serialize"
	"gorpc/src/pb"
	"log"
	"os"
	"testing"
)

func TestSerialize(t *testing.T) {
	t.Parallel()

	file, err := os.Create("test")
	if err != nil {
		t.Fatal(err)
	}
	message := sample.NewLaptop()
	err = serialize.WriteProtoBufToBinaryFile(message, file.Name())
	if err != nil {
		t.Fatal(err)
	}
	message2 := &pb.Laptop{}
	err = serialize.ReadProtoBuffFromBinary(file.Name(), message2)
	if err != nil {
		log.Fatalln(err)
	}
	proto.Equal(message, message2)
}
