package digest

import (
	"errors"
	"fmt"
	repb "github.com/liyouxina/buildzone/protogen/remote_execution"
	"regexp"
	"strconv"
)

var (
	UploadRegex   = regexp.MustCompile(`^(?:(?:(?P<instance_name>.*)/)?uploads/(?P<uuid>[a-f0-9-]{36})/)?(?P<blob_type>blobs|compressed-blobs/zstd)/(?P<hash>[a-f0-9]{64})/(?P<size>\d+)`)
	DownloadRegex = regexp.MustCompile(`^(?:(?P<instance_name>.*)/)?(?P<blob_type>blobs|compressed-blobs/zstd)/(?P<hash>[a-f0-9]{64})/(?P<size>\d+)`)
)

func ParseFromResourceName(resourceName string, matcher *regexp.Regexp) (*repb.Digest, error) {
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
