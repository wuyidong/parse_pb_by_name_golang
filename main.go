package main

import (
	"flag"
	"fmt"

	"pbtest/pb"
	"pbtest/util"
)

var key = flag.String("key", "", "message name, eg: rpc.account.create_account_request")
var input = flag.String("input", "", "input file, should be json")

func main() {
	flag.Parse()

	inputJson, err := util.LoadInputFile(*input)
	if err != nil {
		fmt.Printf("Load input json error: %s\n", err)

	}
	pb.ParsePBMessage(*key, inputJson)
}
