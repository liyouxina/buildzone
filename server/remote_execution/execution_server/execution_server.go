package execution_server

import (
	"context"
	"github.com/golang/protobuf/proto"
	repb "github.com/liyouxina/buildzone/protogen/remote_execution"
	"github.com/liyouxina/buildzone/server/enviroment"
	"github.com/liyouxina/buildzone/server/remote_execution/operation"
	log "github.com/sirupsen/logrus"
	"google.golang.org/genproto/googleapis/longrunning"
)

type ExecutionServer struct {
	env *enviroment.Enviroment
}

func NewExecutionServer(env *enviroment.Enviroment) *ExecutionServer {
	return &ExecutionServer{
		env: env,
	}
}

func (s *ExecutionServer) Execute(req *repb.ExecuteRequest, stream repb.Execution_ExecuteServer) error {
	log.Infof("ExecutionServer Execute req %+v", *req)
	ctx := context.Background()
	if req.ActionDigest != nil {
		cache := s.env.GetCache()
		blob, err := cache.Get(ctx, req.ActionDigest)
		if err != nil {
			return err
		}
		actionResult := repb.ActionResult{}
		_ = proto.Unmarshal(blob, &actionResult)

		req.
		repb.Tree
		operation.Assemble(repb.ExecutionStage_COMPLETED, req.ActionDigest)
		stream.Send(&longrunning.Operation{
			Name: `{}`,
		})
	}
	return nil
}

func (s *ExecutionServer) Dispatch(ctx context.Context, req *repb.ExecuteRequest) (string, error) {
	log.Infof("ExecutionServer Dispatch %+v", *req)
	return "", nil
}

func (s *ExecutionServer) WaitExecution(req *repb.WaitExecutionRequest, stream repb.Execution_WaitExecutionServer) error {
	log.Infof("ExecutionServer WaitExecution %+v", *req)
	stream.Send(&longrunning.Operation{
		Name: `7aa839330d21a5a8b125698ef46ac614286bf64859ccfab1faae70f315374262`,
	})
	repb.ExecuteResponse
	return nil
}

func (s *ExecutionServer) PublishOperation(stream repb.Execution_PublishOperationServer) error {
	log.Infof("ExecutionServer PublishOperation")
	return nil
}