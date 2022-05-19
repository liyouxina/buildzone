package execution_server

import (
	"context"
	"github.com/golang/protobuf/proto"
	repb "github.com/liyouxina/buildzone/protogen/remote_execution"
	"github.com/liyouxina/buildzone/server/enviroment"
	log "github.com/sirupsen/logrus"
	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/protobuf/types/known/anypb"
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
		action := &repb.Action{}
		err = proto.Unmarshal(blob, action)
		if err != nil {
			return err
		}
		blob, err = cache.Get(ctx, action.CommandDigest)
		if err != nil {
			return err
		}
		command := &repb.Command{}
		err = proto.Unmarshal(blob, command)
		if err != nil {
			return err
		}
		metadata, _ := anypb.New(&repb.ExecuteOperationMetadata{
			Stage: repb.ExecutionStage_COMPLETED,
			ActionDigest: &repb.Digest{
				Hash:      "8e9cf629af95185ce4d758f8f3aeced9d8c5f5b6abc68028199dc58cfccdbb2b",
				SizeBytes: 144,
			},
		})
		stream.Send(&longrunning.Operation{
			Name:     "blobs/8e9cf629af95185ce4d758f8f3aeced9d8c5f5b6abc68028199dc58cfccdbb2b/57808",
			Metadata: metadata,
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
	metadata, _ := anypb.New(&repb.ExecuteOperationMetadata{
		Stage: repb.ExecutionStage_COMPLETED,
	})
	stream.Send(&longrunning.Operation{
		Name:     "blobs/8e9cf629af95185ce4d758f8f3aeced9d8c5f5b6abc68028199dc58cfccdbb2b/57808",
		Metadata: metadata,
	})
	return nil
}

func (s *ExecutionServer) PublishOperation(stream repb.Execution_PublishOperationServer) error {
	log.Infof("ExecutionServer PublishOperation")
	return nil
}
