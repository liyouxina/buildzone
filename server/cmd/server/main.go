package main

import (
	repb "github.com/liyouxina/buildzone/protogen/remote_execution"
	"github.com/liyouxina/buildzone/server/enviroment"
	"github.com/liyouxina/buildzone/server/remote_cache/action_cache_server"
	"github.com/liyouxina/buildzone/server/remote_cache/byte_stream_server"
	"github.com/liyouxina/buildzone/server/remote_cache/capabilities_server"
	"github.com/liyouxina/buildzone/server/remote_cache/content_addressable_storage_server"
	bspb "google.golang.org/genproto/googleapis/bytestream"
	"google.golang.org/grpc"
	"net"
)

func main() {
	enviroment.GlobalEnviroment
	grpcServer := grpc.NewServer()
	capabilitiesServer := capabilities_server.NewCapabilitiesServer()
	repb.RegisterCapabilitiesServer(grpcServer, capabilitiesServer)
	casServer := content_addressable_storage_server.NewContentAddressableStorageServer()
	repb.RegisterContentAddressableStorageServer(grpcServer, casServer)
	byteStreamServer := byte_stream_server.NewByteStreamServer()
	bspb.RegisterByteStreamServer(grpcServer, byteStreamServer)
	actionCacheServer := action_cache_server.NewActionCacheServer()
	repb.RegisterActionCacheServer(grpcServer, actionCacheServer)
	lis, _ := net.Listen("tcp", "localhost:28701")
	grpcServer.Serve(lis)
}
