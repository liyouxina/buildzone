package disk_cache

import (
	"context"
	"fmt"
	repb "github.com/liyouxina/buildzone/protogen/remote_execution"
	"github.com/liyouxina/buildzone/server/util/disk"
	log "github.com/sirupsen/logrus"
	"io"
)

func NewDiskCache(path string) DiskCache {
	return DiskCache{
		path: path,
	}
}

type DiskCache struct {
	path string
}

func (c DiskCache) FindMissing(ctx context.Context, digests []*repb.Digest) ([]*repb.Digest, error) {
	missing := []*repb.Digest{}
	for _, digest := range digests {
		contains, err := disk.Exists(fmt.Sprintf("%s/%s", c.path, digest.Hash))
		if contains {
			continue
		}
		if err != nil {
			log.Error("disk error ", err.Error())
		}
		missing = append(missing, digest)
	}
	return missing, nil
}

func (c DiskCache) Contains(ctx context.Context, d *repb.Digest) (bool, error) {
	contains, err := disk.Exists(fmt.Sprintf("%s/%s", c.path, d.Hash))
	if contains {
		return true, nil
	}
	if err != nil {
		log.Error("disk error ", err.Error())
		return false, err
	}
	return false, nil
}

func (c DiskCache) Get(ctx context.Context, d *repb.Digest) ([]byte, error) {
	return disk.Read(fmt.Sprintf("%s/%s", c.path, d.Hash))
}

func (c DiskCache) GetMulti(ctx context.Context, digests []*repb.Digest) (map[*repb.Digest][]byte, error) {
	return nil, nil
}

func (c DiskCache) Set(ctx context.Context, d *repb.Digest, data []byte) error {
	return disk.WriteOver(fmt.Sprintf("%s/%s", c.path, d.Hash), data)
}

func (c DiskCache) SetMulti(ctx context.Context, kvs map[*repb.Digest][]byte) error {
	var err error
	for d, data := range kvs {
		fullPath := fmt.Sprintf("%s/%s", c.path, d.Hash)
		err = disk.WriteOver(fullPath, data)
	}
	if err != nil {
		return err
	}
	return nil
}

func (c DiskCache) Delete(ctx context.Context, d *repb.Digest) error {
	return nil
}

func (c DiskCache) Reader(ctx context.Context, d *repb.Digest, offset int64) (io.ReadCloser, error) {
	return nil, nil
}

func (c DiskCache) Writer(ctx context.Context, d *repb.Digest) (io.WriteCloser, error) {
	return nil, nil
}

func (c DiskCache) WriteAppend(ctx context.Context, d *repb.Digest, data []byte) error {
	return disk.WriteAppend(fmt.Sprintf("%s/%s", c.path, d.Hash), data)
}
