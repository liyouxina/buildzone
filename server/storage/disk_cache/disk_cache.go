package disk_cache

import (
	"context"
	repb "github.com/liyouxina/buildzone/protogen/remote_execution"
)

type DiskCache struct {
	Path string
}

func (c *DiskCache) FindMissing(ctx context.Context, digests []*repb.Digest) ([]*repb.Digest, error) {

	return nil, nil
}
