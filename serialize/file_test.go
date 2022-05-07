package serialize_test

import (
	"google.golang.org/protobuf/proto"
	"gorpc/pb"
	"gorpc/sample"
	"gorpc/serialize"
	"log"
	"os"
	"testing"
)

func TestSerialize(t *testing.T) {
	t.Parallel()

	file, err := createFile("test")
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
	if !proto.Equal(message, message2) { // check if both Messages are equal
		log.Fatalln("Messages are not equal !")
	}
	jsonFile, err := createFile("test.json")
	if err != nil {
		log.Fatalln(err)
	}
	err = serialize.WriteProtoBufToJson(message, jsonFile.Name())
	if err != nil {
		log.Fatalln(err)
	}
}

func createFile(filename string) (*os.File, error) {
	return os.Create(filename)
}
