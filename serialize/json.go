package serialize

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func ProtoBufToJson(message proto.Message) ([]byte, error) {
	marshaller := protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
		Indent:          " ",
		UseProtoNames:   true,
	}
	data, err := marshaller.Marshal(message)
	if err != nil {
		return []byte(nil), err
	}
	return data, nil
}
