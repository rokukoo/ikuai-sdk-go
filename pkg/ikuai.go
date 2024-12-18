package ikuai

import (
	"fmt"
	"log"
	"net/http"
)

type RpcClient struct {
	config *Config
	token  string
	client *http.Client
}

func (c *RpcClient) cookie() string {
	return fmt.Sprintf("sess_key=%s;username=%s;login=1", c.token, c.config.Username)
}

func (c *RpcClient) Printf(format string, v ...any) {
	log.Printf("[IKuai] "+format, v)
}

var defaultConfig *Config
var defaultClient *RpcClient

func DefaultClient() (*RpcClient, error) {
	if defaultConfig == nil {
		return nil, fmt.Errorf("default config is not initialized")
	}
	if defaultClient == nil {
		return nil, fmt.Errorf("default client is not initialized")
	}
	return defaultClient, nil
}

func SetDefaultConfig(config *Config) (err error) {
	defaultConfig = config
	defaultClient = NewRpcClient(defaultConfig)
	err = defaultClient.login()
	return err
}

func NewRpcClient(config *Config) *RpcClient {
	return &RpcClient{
		config: config,
		client: &http.Client{},
	}
}
