package action_cache_server

import (
	"context"
	"github.com/golang/protobuf/proto"
	repb "github.com/liyouxina/buildzone/protogen/remote_execution"
	"github.com/liyouxina/buildzone/server/enviroment"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ActionCacheServer struct {
	env *enviroment.Enviroment
}

func NewActionCacheServer(env *enviroment.Enviroment) *ActionCacheServer {
	return &ActionCacheServer{
		env: env,
	}
}

func (s *ActionCacheServer) GetActionResult(ctx context.Context, req *repb.GetActionResultRequest) (*repb.ActionResult, error) {
	log.Infof("ActionCacheServer GetActionResult %+v", *req)
	response := &repb.ActionResult{}
	cache := s.env.GetCache()
	blob, err := cache.Get(ctx, req.ActionDigest)
	if err != nil || blob == nil || len(blob) == 0 {
		return nil, status.Error(codes.NotFound, "ActionResult (%s) not found: %s")
	}
	if err := proto.Unmarshal(blob, response); err != nil {
		log.Errorf("ActionCacheServer GetActionResult error unmarshal %v", err)
		return nil, err
	}
	log.Infof("ActionCacheServer GetActionResult response %+v", response)
	return response, nil
}

func (s *ActionCacheServer) UpdateActionResult(ctx context.Context, req *repb.UpdateActionResultRequest) (*repb.ActionResult, error) {
	log.Infof("ActionCacheServer UpdateActionResult %+v", *req)
	cache := s.env.GetCache()
	d := req.GetActionDigest()
	data, _ := proto.Marshal(req.ActionResult)
	err := cache.Set(ctx, d, data)
	if err != nil {
		return nil, err
	}
	return req.ActionResult, nil
}
