package firesearch

import (
	"context"
	"fmt"
	"github.com/pacedotdev/firesearch-sdk/clients/go/firesearch"
)

// AccessKeyStore is responsible for retrieving access keys from firesearch
type AccessKeyStore struct {
	Service Service
}

// GetAccessKey generates an access key for an index
func (c *AccessKeyStore) GetAccessKey(ctx context.Context, index string) (string, error) {
	accessKeyService := firesearch.NewAccessKeyService(c.Service.Client)
	keyReq := firesearch.GenerateKeyRequest{IndexPathPrefix: index}
	keyResp, err := accessKeyService.GenerateKey(ctx, keyReq)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve access key: %w", err)
	}
	return keyResp.AccessKey, nil
}
