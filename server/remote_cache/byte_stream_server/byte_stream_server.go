package byte_stream_server

import (
	"context"
	log "github.com/sirupsen/logrus"

	bspb "google.golang.org/genproto/googleapis/bytestream"
)

func NewByteStreamServer() *ByteStreamServer {
	return nil
}

type ByteStreamServer struct {
}

func (s *ByteStreamServer) Read(req *bspb.ReadRequest, stream bspb.ByteStream_ReadServer) error {
	log.Infof("ByteStreamServer Read")
	return nil
}

func (s *ByteStreamServer) Write(stream bspb.ByteStream_WriteServer) error {
	log.Infof("ByteStreamServer Write")
	return nil
}

func (s *ByteStreamServer) QueryWriteStatus(ctx context.Context, req *bspb.QueryWriteStatusRequest) (*bspb.QueryWriteStatusResponse, error) {
	log.Infof("ByteStreamServer QueryWriteStatus")
	return &bspb.QueryWriteStatusResponse{
		CommittedSize: 0,
		Complete:      false,
	}, nil
}
