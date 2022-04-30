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
	msg := pb.Laptop{}
	err = serialize.ReadProtoBuffFromBinary(file.Name(), msg)
	if err != nil {
		log.Fatalln(err)
	}
	proto.Equal(message)
}
