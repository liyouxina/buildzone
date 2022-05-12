package execution_server

import (
	"context"
	repb "github.com/liyouxina/buildzone/protogen/remote_execution"
	"github.com/liyouxina/buildzone/server/enviroment"
	log "github.com/sirupsen/logrus"
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
	log.Infof("ExecutionServer Execute")
	return nil
}

func (s *ExecutionServer) Dispatch(ctx context.Context, req *repb.ExecuteRequest) (string, error) {
	log.Infof("ExecutionServer Dispatch")
	return "", nil
}

func (s *ExecutionServer) WaitExecution(req *repb.WaitExecutionRequest, stream repb.Execution_WaitExecutionServer) error {
	log.Infof("ExecutionServer WaitExecution")
	return nil
}

func (s *ExecutionServer) MarkExecutionFailed(ctx context.Context, taskID string, reason error) error {
	log.Infof("ExecutionServer MarkExecutionFailed")
	return nil
}

func (s *ExecutionServer) PublishOperation(stream repb.Execution_PublishOperationServer) error {
	log.Infof("ExecutionServer PublishOperation")
	return nil
}

func (s *ExecutionServer) Cancel(ctx context.Context, invocationID string) error {
	log.Infof("ExecutionServer Cancel")
	return nil
}