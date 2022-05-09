package action_cache_server

import (
	"context"
	"errors"
	"fmt"
	repb "github.com/liyouxina/buildzone/protogen/remote_execution"
	log "github.com/sirupsen/logrus"
)

type ActionCacheServer struct {
}

func NewActionCacheServer() *ActionCacheServer {
	return &ActionCacheServer{}
}

func (s *ActionCacheServer) GetActionResult(ctx context.Context, req *repb.GetActionResultRequest) (*repb.ActionResult, error) {
	if req.ActionDigest == nil {
		return nil, errors.New("ActionDigest is required")
	}
	return nil, nil
}

func (s *ActionCacheServer) UpdateActionResult(ctx context.Context, req *repb.UpdateActionResultRequest) (*repb.ActionResult, error) {
	log.Infof("ActionCacheServer UpdateActionResult")
	fmt.Println(req.ActionResult)
	return nil, nil
}
