package pb

import (
	"fmt"
	"reflect"

	"github.com/golang/protobuf/jsonpb"
)

func ParsePBMessage(msgName string, msgContent string) {
	msg, err := GetRPCMessageObjectByName(msgName)
	if err != nil {
		fmt.Printf("Parse pb request error:%s\n", err)
		return
	}
	err = jsonpb.UnmarshalString(msgContent, msg)
	if err != nil {
		fmt.Printf("Unmarshal pb request error:%s\n", err)
		return
	}
	msgExtension := GetRPCMessageExtension(msgName)

	fmt.Println(msg)
	fmt.Println(reflect.TypeOf(msgExtension))
	return
}
