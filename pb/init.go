package pb

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/golang/protobuf/proto"
	_ "pbtest/proto/account"
	"pbtest/proto/rpc"
)

// message id -> *proto.ExtensionDesc
// 记录 rpc.Body 的拓展消息
var RPCMessageBodyExtensions map[int32]*proto.ExtensionDesc

func init() {
	RPCMessageBodyExtensions = proto.RegisteredExtensions((*rpc.Body)(nil))
}

// some utils for UMessage

// msgName: rpc.account.create_account_request
func GetRPCMessageObjectByName(msgName string) (msg proto.Message, err error) {
	msgType := reflect.TypeOf(GetRPCMessageExtension(msgName).ExtensionType)
	if msgType == nil {
		err = fmt.Errorf("can't find message type")
		return
	}
	// msgType is pointer
	msg = reflect.Indirect(reflect.New(msgType.Elem())).Addr().Interface().(proto.Message)
	return
}

// msgName: rpc.account.create_account_request
// namePrefix: rpc.account
// name: create_account_request
func GetNamePrefix(msgName string) (prefix string) {
	items := strings.Split(msgName, ".")
	prefix = strings.Join(items[0:len(items)-1], ".")
	return
}

func GetName(msgName string) (name string) {
	items := strings.Split(msgName, ".")
	name = items[len(items)-1]
	return
}

func GetRPCMessageId(msgName string) (msgId int32) {
	msgTypeName := GetNamePrefix(msgName) + ".MessageType"
	mapMsgNameId := proto.EnumValueMap(msgTypeName)
	msgId = mapMsgNameId[strings.ToUpper(GetName(msgName))]
	return
}

func GetRPCMessageExtension(msgName string) (extension *proto.ExtensionDesc) {
	msgId := GetRPCMessageId(msgName)
	extension = RPCMessageBodyExtensions[msgId]
	return
}
