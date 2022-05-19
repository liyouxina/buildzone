package main

import (
	"github.com/golang/protobuf/proto"
	repb "github.com/liyouxina/buildzone/protogen/remote_execution"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
)

func main() {
	bytes, _ := ioutil.ReadFile("/tmp/buildzone/2aead63a8a8cdf4efb9deeafb377bef9c19534075e358076122a13a7ac1b5ca2")
	tree := repb.Tree{}
	print(string(bytes))
	err := proto.Unmarshal(bytes, &tree)
	if err != nil {
		panic(err)
	}
	log.Infof("%+v", tree)
}
