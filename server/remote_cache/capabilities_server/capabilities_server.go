package capabilities_server

import (
	"context"
	repb "github.com/liyouxina/buildzone/protogen/remote_execution"
)

type CapabilitiesServer struct {
	supportCAS        bool
	supportRemoteExec bool
	supportZstd       bool
}

func NewCapabilitiesServer() *CapabilitiesServer {
	return &CapabilitiesServer{}
}
func (s *CapabilitiesServer) GetCapabilities(ctx context.Context, req *repb.GetCapabilitiesRequest) (*repb.ServerCapabilities, error) {
	return nil, nil
}
