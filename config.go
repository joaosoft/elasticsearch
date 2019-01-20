package elastic

import (
	"fmt"

	"github.com/joaosoft/manager"
)

// AppConfig ...
type AppConfig struct {
	Elastic ElasticConfig `json:"elastic"`
}

// ElasticConfig ...
type ElasticConfig struct {
	Log struct {
		Level string `json:"level"`
	} `json:"log"`
	Endpoint string `json:"endpoint"`
}

// NewConfig ...
func NewConfig() (*AppConfig, manager.IConfig, error) {
	appConfig := &AppConfig{}
	simpleConfig, err := manager.NewSimpleConfig(fmt.Sprintf("/config/app.%s.json", GetEnv()), appConfig)

	return appConfig, simpleConfig, err
}
