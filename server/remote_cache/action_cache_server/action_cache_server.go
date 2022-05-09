package action_cache_server

import (
	"context"
	repb "github.com/liyouxina/buildzone/protogen/remote_execution"
)

type ActionCacheServer struct {
}

func NewActionCacheServer() *ActionCacheServer {
	return &ActionCacheServer{}
}

func (s *ActionCacheServer) GetActionResult(ctx context.Context, req *repb.GetActionResultRequest) (*repb.ActionResult, error) {
	return nil, nil
}

func (s *ActionCacheServer) UpdateActionResult(ctx context.Context, req *repb.UpdateActionResultRequest) (*repb.ActionResult, error) {
	return nil, nil
}
