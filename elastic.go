package elastic

import (
	"sync"

	"github.com/joaosoft/logger"
	"github.com/joaosoft/manager"
)

type Elastic struct {
	config        *ElasticConfig
	isLogExternal bool
	logger logger.ILogger
	pm            *manager.Manager
	mux           sync.Mutex
}

// NewElastic ...
func NewElastic(options ...ElasticOption) *Elastic {
	config, simpleConfig, err := NewConfig()
	log := logger.NewLogDefault("elastic", logger.DebugLevel)

	service := &Elastic{
		pm:     manager.NewManager(manager.WithRunInBackground(false)),
		config: &config.Elastic,
		logger: log,
	}

	if err != nil {
		service.logger.Error(err.Error())
	} else {
		service.pm.AddConfig("config_app", simpleConfig)
		level, _ := logger.ParseLevel(config.Elastic.Log.Level)
		service.logger.Debugf("setting log level to %s", level)
		service.logger.Reconfigure(logger.WithLevel(level))
	}

	if service.isLogExternal {
		service.pm.Reconfigure(manager.WithLogger(log))
	}

	service.Reconfigure(options...)

	return service
}

func (e *Elastic) Count() *CountService {
	return NewCountService(e)
}

func (e *Elastic) Create() *CreateService {
	return NewCreateService(e)
}

func (e *Elastic) Update() *UpdateService {
	return NewUpdateService(e)
}

func (elastic *Elastic) Delete() *DeleteService {
	return NewDeleteService(elastic)
}

func (e *Elastic) Search() *SearchService {
	return NewSearchService(e)
}

func (e *Elastic) ExistsIndex() *ExistsIndexService {
	return NewExistsIndexService(e)
}

func (e *Elastic) CreateIndex() *CreateIndexService {
	return NewCreateIndexService(e)
}

func (e *Elastic) UpdateIndex() *CreateIndexService {
	return NewCreateIndexService(e)
}

func (e *Elastic) DeleteIndex() *DeleteIndexService {
	return NewDeleteIndexService(e)
}

func (e *Elastic) Bulk() *BulkService {
	return NewBulkService(e)
}
