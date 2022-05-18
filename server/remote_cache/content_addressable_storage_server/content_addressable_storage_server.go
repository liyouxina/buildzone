package content_addressable_storage_server

import (
	"context"
	repb "github.com/liyouxina/buildzone/protogen/remote_execution"
	"github.com/liyouxina/buildzone/server/enviroment"
	log "github.com/sirupsen/logrus"
	statuspb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
)

type ContentAddressableStorageServer struct {
	env *enviroment.Enviroment
}

func NewContentAddressableStorageServer(env *enviroment.Enviroment) *ContentAddressableStorageServer {
	return &ContentAddressableStorageServer{
		env: env,
	}
}

func (s *ContentAddressableStorageServer) FindMissingBlobs(ctx context.Context, req *repb.FindMissingBlobsRequest) (*repb.FindMissingBlobsResponse, error) {
	log.Infof("ContentAddressableStorageServer FindMissingBlobs %+v", *req)
	response := &repb.FindMissingBlobsResponse{}
	cache := s.env.GetCache()
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
	response := &repb.BatchUpdateBlobsResponse{}
	cache := s.env.GetCache()
	kvs := make(map[*repb.Digest][]byte, len(req.Requests))
	for _, res := range req.Requests {
		kvs[res.Digest] = res.Data
	}
	err := cache.SetMulti(ctx, kvs)
	if err != nil {
		return nil, err
	}
	for uploadDigest := range kvs {
		response.Responses = append(response.Responses, &repb.BatchUpdateBlobsResponse_Response{
			Digest: uploadDigest,
			Status: &statuspb.Status{Code: int32(codes.OK)},
		})
	}
	return response, nil
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
