package byte_stream_server

import (
	"context"
	"errors"
	"fmt"
	repb "github.com/liyouxina/buildzone/protogen/remote_execution"
	"github.com/liyouxina/buildzone/server/enviroment"
	log "github.com/sirupsen/logrus"
	bspb "google.golang.org/genproto/googleapis/bytestream"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"regexp"
	"strconv"
)

var (
	uploadRegex   = regexp.MustCompile(`^(?:(?:(?P<instance_name>.*)/)?uploads/(?P<uuid>[a-f0-9-]{36})/)?(?P<blob_type>blobs|compressed-blobs/zstd)/(?P<hash>[a-f0-9]{64})/(?P<size>\d+)`)
	downloadRegex = regexp.MustCompile(`^(?:(?P<instance_name>.*)/)?(?P<blob_type>blobs|compressed-blobs/zstd)/(?P<hash>[a-f0-9]{64})/(?P<size>\d+)`)
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
	digest, err := parseFromResourceName(req.GetResourceName(), downloadRegex)
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
			if digest, err = parseFromResourceName(req.GetResourceName(), uploadRegex); err != nil {
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

func parseFromResourceName(resourceName string, matcher *regexp.Regexp) (*repb.Digest, error) {
	match := matcher.FindStringSubmatch(resourceName)
	result := make(map[string]string, len(match))
	for i, name := range matcher.SubexpNames() {
		if i != 0 && name != "" && i < len(match) {
			result[name] = match[i]
		}
	}
	hash, hashOK := result["hash"]
	sizeStr, sizeOK := result["size"]
	if !hashOK || !sizeOK {
		return nil, errors.New(fmt.Sprintf("Unparsable resource name: %s", resourceName))
	}
	sizeBytes, _ := strconv.ParseInt(sizeStr, 10, 0)
	return &repb.Digest{
		Hash:      hash,
		SizeBytes: sizeBytes,
	}, nil
}
