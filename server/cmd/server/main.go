package main

import (
	repb "github.com/liyouxina/buildzone/protogen/remote_execution"
	"github.com/liyouxina/buildzone/server/enviroment"
	"github.com/liyouxina/buildzone/server/remote_cache/action_cache_server"
	"github.com/liyouxina/buildzone/server/remote_cache/byte_stream_server"
	"github.com/liyouxina/buildzone/server/remote_cache/capabilities_server"
	"github.com/liyouxina/buildzone/server/remote_cache/content_addressable_storage_server"
	"github.com/liyouxina/buildzone/server/remote_execution/execution_server"
	"github.com/liyouxina/buildzone/server/storage/disk_cache"
	log "github.com/sirupsen/logrus"
	bspb "google.golang.org/genproto/googleapis/bytestream"
	"google.golang.org/grpc"
	"net"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	diskCache := disk_cache.NewDiskCache("/tmp/buildzone")
	env := NewEnviromentFromConfig()
	env.SetCache(diskCache)
	grpcServer := grpc.NewServer()
	capabilitiesServer := capabilities_server.NewCapabilitiesServer()
	repb.RegisterCapabilitiesServer(grpcServer, capabilitiesServer)
	casServer := content_addressable_storage_server.NewContentAddressableStorageServer(env)
	repb.RegisterContentAddressableStorageServer(grpcServer, casServer)
	byteStreamServer := byte_stream_server.NewByteStreamServer(env)
	bspb.RegisterByteStreamServer(grpcServer, byteStreamServer)
	actionCacheServer := action_cache_server.NewActionCacheServer(env)
	repb.RegisterActionCacheServer(grpcServer, actionCacheServer)
	executionServer := execution_server.NewExecutionServer(env)
	repb.RegisterExecutionServer(grpcServer, executionServer)
	lis, _ := net.Listen("tcp", "localhost:28701")
	grpcServer.Serve(lis)
}

func NewEnviromentFromConfig() *enviroment.Enviroment {
	// TODO
	return &enviroment.Enviroment{}
}
