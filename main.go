package main

import (
	"fmt"

	"github.com/michu990902/go-pb-test/src/example_simple"
)

func main() {
	doSimple()

}

func doSimple() {
	sm := example_simple.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My simple message",
		SampleList: []int32{1, 4, 7, 8},
	}
	fmt.Println(sm)
}

// func jsonDemo(sm proto.Message){
// 	smAsString := toJSON(sm)
// 	fmt.Println(smAsString)
// }

// func toJSON(sm proto.Message) string{
// 	marshaler := jsonpb.Marshaler{}
// 	marshaler.MarshalToString(pb proto.Message)
// }
