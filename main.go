package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	complexpb "github.com/michu990902/go-pb-test/src/complex"
	enumpb "github.com/michu990902/go-pb-test/src/enum_example"
	"github.com/michu990902/go-pb-test/src/simple"
)

func main() {
	// sm := doSimple()
	// jsonDemo(sm)

	// doEnum()

	doComplex()
}

func doComplex() {
	cm := complexpb.ComplexMessage{
		OneDummy: &complexpb.DummyMessage{
			Id:   1,
			Name: "test1",
		},
		MultipleDummy: []*complexpb.DummyMessage{
			&complexpb.DummyMessage{
				Id:   2,
				Name: "test2",
			},
			&complexpb.DummyMessage{
				Id:   3,
				Name: "test3",
			},
		},
	}
	fmt.Println(cm)
}

func doEnum() {
	em := enumpb.EnumMessage{
		Id:           42,
		DayOfTheWeek: enumpb.DayOfWeek_TUESDAY,
	}
	fmt.Println(em)
}

func jsonDemo(sm proto.Message) {
	smAsString := toJSON(sm)
	fmt.Println(smAsString)

	sm2 := &simple.SimpleMessage{}
	fromJSON(smAsString, sm2)
	fmt.Println("Successfully created proto struct:", sm2)
}

func doSimple() proto.Message {
	sm := &simple.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My simple message",
		SampleTest: []int32{1, 4, 7, 8},
	}
	fmt.Println(sm)
	fmt.Println(sm.GetId())
	return sm
}

// func jsonDemo(sm proto.Message){
// 	smAsString := toJSON(sm)
// 	fmt.Println(smAsString)
// }

func toJSON(sm proto.Message) string {
	marshaler := jsonpb.Marshaler{}
	out, err := marshaler.MarshalToString(sm)
	if err != nil {
		log.Fatalln("Can't convert to JSON: ", err)
	}
	return out
}

func fromJSON(in string, pb proto.Message) {
	err := jsonpb.UnmarshalString(in, pb)
	if err != nil {
		log.Fatalln("Couldn't unmarshal the JSON into the pb struct: ", err)
	}
}
