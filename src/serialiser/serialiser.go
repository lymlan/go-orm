package serialiser

import (
	"fmt"
	"io/ioutil"
	"log"

	"google.golang.org/protobuf/proto"
)

func Serialise(filepath string, message proto.Message) {

	data, err := proto.Marshal(message)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	err = ioutil.WriteFile(filepath, data, 0644)
	if err != nil {
		log.Fatal("Something went wrong")
	}
	fmt.Println(data)
}

func Deserialise(filepath string, message proto.Message) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal("Something went wrong again")
	}
	err = proto.Unmarshal(data, message)
	if err != nil {
		log.Fatal("Something went wrong writing unmarshalling the data into the protobuffer")
	}

}
