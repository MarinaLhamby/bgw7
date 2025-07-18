package application

import (
	"app/internal"
	"app/internal/repository"
	"app/internal/service"
)

const (
	filePath = "./docs/db/json"
)

type ApplicationLoad struct {
	config             ApplicationConfig
	applicationService internal.ServiceDataLoad
}

func NewLoadApplication(config *ConfigApplicationDefault) *ApplicationLoad {
	defaultCfg := &ConfigApplicationDefault{
		Db: nil,
	}
	if config != nil {
		if config.Db != nil {
			defaultCfg.Db = config.Db
		}
	}
	appConfig := ApplicationConfig{
		cfgDb: defaultCfg.Db,
	}
	return &ApplicationLoad{
		config: appConfig,
	}
}

func (l *ApplicationLoad) SetUp() (err error) {
	err = l.config.SetUpDb()
	if err != nil {
		return err
	}
	// - repository
	rpCustomer := repository.NewCustomersMySQL(l.config.db)
	rpProduct := repository.NewProductsMySQL(l.config.db)
	rpInvoice := repository.NewInvoicesMySQL(l.config.db)
	rpSlle := repository.NewSalesMySQL(l.config.db)

	service := service.NewDataLoadDefault(rpCustomer, rpInvoice, rpProduct, rpSlle)

	l.applicationService = service
	return nil
}

func (l *ApplicationLoad) Run() (err error) {
	defer l.config.db.Close()
	err = l.applicationService.Load(filePath)
	return err
}
