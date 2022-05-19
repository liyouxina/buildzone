package byte_stream_server

import (
	"context"
	repb "github.com/liyouxina/buildzone/protogen/remote_execution"
	"github.com/liyouxina/buildzone/server/enviroment"
	dg "github.com/liyouxina/buildzone/server/util/digest"
	log "github.com/sirupsen/logrus"
	bspb "google.golang.org/genproto/googleapis/bytestream"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
)

func NewByteStreamServer(env *enviroment.Enviroment) *ByteStreamServer {
	return &ByteStreamServer{
		env: env,
	}
}

type ByteStreamServer struct {
	env *enviroment.Enviroment
}

func (s *ByteStreamServer) Read(req *bspb.ReadRequest, stream bspb.ByteStream_ReadServer) error {
	log.Infof("ByteStreamServer Read %+v", *req)
	ctx := stream.Context()
	cache := s.env.GetCache()
	digest, err := dg.ParseFromResourceName(req.GetResourceName(), dg.DownloadRegex)
	if err != nil {
		return status.Error(codes.NotFound, "ByteStreamServer Read")
	}
	blob, err := cache.Get(ctx, digest)
	if err != nil {
		return status.Errorf(codes.NotFound, "ByteStreamServer Read cache get")
	}
	stream.Send(&bspb.ReadResponse{
		Data: blob,
	})
	return nil
}

func (s *ByteStreamServer) Write(stream bspb.ByteStream_WriteServer) error {
	log.Infof("ByteStreamServer Write")
	ctx := stream.Context()
	var digest *repb.Digest
	var downloadSize int64
	cache := s.env.GetCache()
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if digest == nil {
			if digest, err = dg.ParseFromResourceName(req.GetResourceName(), dg.UploadRegex); err != nil {
				return err
			}
		}
		if err = cache.WriteAppend(ctx, digest, req.Data); err != nil {
			return err
		}
		downloadSize += int64(len(req.Data))
		if req.FinishWrite {
			return stream.SendAndClose(&bspb.WriteResponse{
				CommittedSize: downloadSize,
			})
		}
	}
	return nil
}

func (s *ByteStreamServer) QueryWriteStatus(ctx context.Context, req *bspb.QueryWriteStatusRequest) (*bspb.QueryWriteStatusResponse, error) {
	log.Infof("ByteStreamServer QueryWriteStatus %+v", *req)
	return &bspb.QueryWriteStatusResponse{
		CommittedSize: 0,
		Complete:      false,
	}, nil
}
