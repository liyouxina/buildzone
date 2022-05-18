package operation

import (
	repb "github.com/liyouxina/buildzone/protogen/remote_execution"
	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/protobuf/types/known/anypb"
)

func Assemble(stage repb.ExecutionStage_Value, name string, digest *repb.Digest, er *repb.ExecuteResponse) (*longrunning.Operation, error) {
	if digest == nil || er == nil {
		return nil, nil
	}
	metadata, err := anypb.New(&repb.ExecuteOperationMetadata{
		Stage:        stage,
		ActionDigest: digest,
	})
	if err != nil {
		return nil, err
	}
	operation := &longrunning.Operation{
		Name:     name,
		Metadata: metadata,
	}
	result, err := anypb.New(er)
	if err != nil {
		return nil, err
	}
	operation.Result = &longrunning.Operation_Response{Response: result}

	if stage == repb.ExecutionStage_COMPLETED {
		operation.Done = true
	}
	return operation, nil
}
