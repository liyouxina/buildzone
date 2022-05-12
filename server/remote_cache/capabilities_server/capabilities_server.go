package capabilities_server

import (
	"context"
	repb "github.com/liyouxina/buildzone/protogen/remote_execution"
	smpb "github.com/liyouxina/buildzone/protogen/semver"
	log "github.com/sirupsen/logrus"
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
	log.Infof("CapabilitiesServer GetCapabilities")
	response := repb.ServerCapabilities{
		LowApiVersion:  &smpb.SemVer{Major: int32(2)},
		HighApiVersion: &smpb.SemVer{Major: int32(99), Minor: int32(9)},
	}
	response.CacheCapabilities = &repb.CacheCapabilities{
		DigestFunction: []repb.DigestFunction_Value{repb.DigestFunction_SHA256},
		ActionCacheUpdateCapabilities: &repb.ActionCacheUpdateCapabilities{
			UpdateEnabled: true,
		},
		CachePriorityCapabilities: &repb.PriorityCapabilities{
			Priorities: []*repb.PriorityCapabilities_PriorityRange{
				{
					MinPriority: 0,
					MaxPriority: 0,
				},
			},
		},
		MaxBatchTotalSizeBytes:      0,
		SymlinkAbsolutePathStrategy: repb.SymlinkAbsolutePathStrategy_ALLOWED,
	}
	response.ExecutionCapabilities = &repb.ExecutionCapabilities{
		DigestFunction: repb.DigestFunction_SHA256,
		ExecEnabled: true,
	}
	return &response, nil
}
