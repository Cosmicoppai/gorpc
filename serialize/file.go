package serialize

import (
	"google.golang.org/protobuf/proto"
	"io/ioutil"
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
