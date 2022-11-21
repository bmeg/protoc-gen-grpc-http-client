package tests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/bmeg/protoc-gen-grpc-http-client/simple"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway/httprule"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
)

func build(temp httprule.Template, msg proto.Message) string {
	//stacksize := len(temp.OpCodes) / 2

	var stack []string

	for i := 0; i < len(temp.OpCodes); i += 2 {
		op := utilities.OpCode(temp.OpCodes[i])
		operand := temp.OpCodes[i+1]

		if op == utilities.OpPush {
			//fmt.Printf("Push %d\n", operand)
		} else if op == utilities.OpPushM {
			//fmt.Printf("PushM %d\n", operand)
		} else if op == utilities.OpLitPush {
			stack = append(stack, temp.Pool[operand])
			//fmt.Printf("LitPush %s\n", temp.Pool[operand])
		} else if op == utilities.OpConcatN {
			//fmt.Printf("ConcatN %d\n", operand)
			n := operand
			l := len(stack) - n
			stack = append(stack[:l], strings.Join(stack[l:], "/"))
		} else if op == utilities.OpCapture {
			//fmt.Printf("Capture: %s\n", temp.Pool[operand])
			msg.ProtoReflect().Range(func(field protoreflect.FieldDescriptor, val protoreflect.Value) bool {
				if field.JSONName() == temp.Pool[operand] {
					stack = append(stack, val.String())
				}
				//fmt.Printf("Field: %s\n", field)
				//fmt.Printf("Value: %s\n", val)
				return true
			})

		}
	}
	segs := strings.Join(stack, "/")
	if temp.Verb != "" {
		return fmt.Sprintf("/%s:%s", segs, temp.Verb)
	}
	return "/" + segs
}

func checkUrlBuild(t *testing.T, path string, msg proto.Message, result string) error {
	c, err := httprule.Parse(path)
	if err != nil {
		t.Error(err)
	}
	s := build(c.Compile(), msg)
	if s != result {
		return fmt.Errorf("%s != %s", s, result)
	}
	return nil
}

type pathTest struct {
	path   string
	msg    proto.Message
	result string
}

func TestPathParse(t *testing.T) {

	tests := []pathTest{
		{"/v1/graph/{graph}", &simple.GraphID{Graph: "gname"}, "/v1/graph/gname"},
		{"/v1/graph/{graph}/edge/{id}", &simple.ElementID{Graph: "gname", Id: "edge1"}, "/v1/graph/gname/edge/edge1"},
		{"/v1/graph", &simple.Empty{}, "/v1/graph"},
	}

	for _, p := range tests {
		err := checkUrlBuild(t, p.path, p.msg, p.result)
		if err != nil {
			t.Error(err)
		}
	}

}
