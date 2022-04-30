package main

import (
	"gorpc/sample"
	"gorpc/serialize"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("File name not specified")
	}
	fileName := os.Args[1]
	msg := sample.NewLaptop()
	err := serialize.WriteProtoBufToBinaryFile(msg, fileName)
	if err != nil {
		log.Fatalln(err)
	}
}
