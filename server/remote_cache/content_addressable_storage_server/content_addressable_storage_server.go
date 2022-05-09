package content_addressable_storage_server

import (
	"context"
	repb "github.com/liyouxina/buildzone/protogen/remote_execution"
	"github.com/liyouxina/buildzone/server/enviroment"
	log "github.com/sirupsen/logrus"
)

const gRPCMaxSize = int64(4194304 - 2000)

type ContentAddressableStorageServer struct {
}

func NewContentAddressableStorageServer() *ContentAddressableStorageServer {
	return &ContentAddressableStorageServer{}
}

func (s *ContentAddressableStorageServer) FindMissingBlobs(ctx context.Context, req *repb.FindMissingBlobsRequest) (*repb.FindMissingBlobsResponse, error) {
	log.Infof("ContentAddressableStorageServer FindMissingBlobs")
	response := &repb.FindMissingBlobsResponse{}
	cache := enviroment.GlobalEnviromentContext.GetCache()
	missing, err := cache.FindMissing(ctx, req.BlobDigests)
	if err != nil {
		return nil, err
	}
	for _, m := range missing {
		response.MissingBlobDigests = append(response.MissingBlobDigests, m)
	}
	return response, nil
}

func (s *ContentAddressableStorageServer) BatchUpdateBlobs(ctx context.Context, req *repb.BatchUpdateBlobsRequest) (*repb.BatchUpdateBlobsResponse, error) {
	log.Infof("ContentAddressableStorageServer BatchUpdateBlobs")
	rsp := &repb.BatchUpdateBlobsResponse{}
	return rsp, nil
}

func (s *ContentAddressableStorageServer) BatchReadBlobs(ctx context.Context, req *repb.BatchReadBlobsRequest) (*repb.BatchReadBlobsResponse, error) {
	log.Infof("ContentAddressableStorageServer BatchReadBlobs")
	rsp := &repb.BatchReadBlobsResponse{}
	return rsp, nil
}

func (s *ContentAddressableStorageServer) GetTree(req *repb.GetTreeRequest, stream repb.ContentAddressableStorage_GetTreeServer) error {
	log.Infof("ContentAddressableStorageServer GetTree")
	return nil
}
