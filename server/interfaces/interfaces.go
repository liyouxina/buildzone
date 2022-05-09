package interfaces

import (
	"context"
	repb "github.com/liyouxina/buildzone/protogen/remote_execution"
	"io"
)

type CacheType int
type Cache interface {
	Contains(ctx context.Context, d *repb.Digest) (bool, error)
	FindMissing(ctx context.Context, digests []*repb.Digest) ([]*repb.Digest, error)
	Get(ctx context.Context, d *repb.Digest) ([]byte, error)
	GetMulti(ctx context.Context, digests []*repb.Digest) (map[*repb.Digest][]byte, error)
	Set(ctx context.Context, d *repb.Digest, data []byte) error
	SetMulti(ctx context.Context, kvs map[*repb.Digest][]byte) error
	Delete(ctx context.Context, d *repb.Digest) error
	Reader(ctx context.Context, d *repb.Digest, offset int64) (io.ReadCloser, error)
	Writer(ctx context.Context, d *repb.Digest) (io.WriteCloser, error)
}
