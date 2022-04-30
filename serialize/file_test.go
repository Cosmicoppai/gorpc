package serialize_test

import (
	"gorpc/sample"
	"gorpc/serialize"
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
}
