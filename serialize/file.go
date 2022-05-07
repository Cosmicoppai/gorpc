package serialize

import (
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
)

func WriteProtoBufToBinaryFile(msg proto.Message, fileName string) error {
	data, err := proto.Marshal(msg)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(fileName, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func ReadProtoBuffFromBinary(fileNames string, message proto.Message) error {

	data, err := ioutil.ReadFile(fileNames)
	if err != nil {
		log.Fatalln(err)
	}
	err = proto.Unmarshal(data, message)
	if err != nil {
		log.Fatalln("Cannot unmarshal the data", err)
	}
	return nil

}

func WriteProtoBufToJson(message proto.Message, filename string) error {
	data, err := ProtoBufToJson(message)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
