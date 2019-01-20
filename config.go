package elastic

import (
	"fmt"

	gomanager "github.com/joaosoft/manager"
)

// AppConfig ...
type AppConfig struct {
	Elastic *ElasticConfig `json:"elastic"`
}

// ElasticConfig ...
type ElasticConfig struct {
	Log struct {
		Level string `json:"level"`
	} `json:"log"`
	Endpoint string `json:"endpoint"`
}

// NewConfig ...
func NewConfig() (*AppConfig, gomanager.IConfig, error) {
	appConfig := &AppConfig{}
	simpleConfig, err := gomanager.NewSimpleConfig(fmt.Sprintf("/config/app.%s.json", GetEnv()), appConfig)
	if err != nil {
		log.Error(err.Error())

		appConfig.Elastic = &ElasticConfig{
			Endpoint: DefaultURL,
		}
	}

	return appConfig, simpleConfig, err
}
