package byte_stream_server

import (
	"context"

	bspb "google.golang.org/genproto/googleapis/bytestream"
)

const (
	readBufSizeBytes = (1024 * 1024 * 4) - (1024 * 256)
)

func NewByteStreamServer() *ByteStreamServer {
	return nil
}

type ByteStreamServer struct {
}

func (s *ByteStreamServer) Read(req *bspb.ReadRequest, stream bspb.ByteStream_ReadServer) error {
	return nil
}

func (s *ByteStreamServer) Write(stream bspb.ByteStream_WriteServer) error {

	return nil
}

func (s *ByteStreamServer) QueryWriteStatus(ctx context.Context, req *bspb.QueryWriteStatusRequest) (*bspb.QueryWriteStatusResponse, error) {

	return &bspb.QueryWriteStatusResponse{
		CommittedSize: 0,
		Complete:      false,
	}, nil
}
