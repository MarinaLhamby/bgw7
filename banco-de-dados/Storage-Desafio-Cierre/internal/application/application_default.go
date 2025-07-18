package application

import (
	"app/internal/handler"
	"app/internal/repository"
	"app/internal/service"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type ApplicationDefault struct {
	appConfig ApplicationConfig
	// cfgAddr is the server address.
	cfgAddr string
	// router is the chi router.
	router *chi.Mux
}

// NewApplicationDefault creates a new ApplicationDefault.
func NewApplicationDefault(config *ConfigApplicationDefault) *ApplicationDefault {
	// default values
	defaultCfg := &ConfigApplicationDefault{
		Db:   nil,
		Addr: ":8080",
	}
	if config != nil {
		if config.Db != nil {
			defaultCfg.Db = config.Db
		}
		if config.Addr != "" {
			defaultCfg.Addr = config.Addr
		}
	}

	appConfig := ApplicationConfig{
		cfgDb: defaultCfg.Db,
	}
	return &ApplicationDefault{
		appConfig: appConfig,
		cfgAddr:   defaultCfg.Addr,
	}
}

// SetUp sets up the application.
func (a *ApplicationDefault) SetUp() (err error) {
	err = a.appConfig.SetUpDb()
	if err != nil {
		return err
	}
	// - repository
	rpCustomer := repository.NewCustomersMySQL(a.appConfig.db)
	rpProduct := repository.NewProductsMySQL(a.appConfig.db)
	rpInvoice := repository.NewInvoicesMySQL(a.appConfig.db)
	rpSale := repository.NewSalesMySQL(a.appConfig.db)
	// - service
	svCustomer := service.NewCustomersDefault(rpCustomer)
	svProduct := service.NewProductsDefault(rpProduct)
	svInvoice := service.NewInvoicesDefault(rpInvoice)
	svSale := service.NewSalesDefault(rpSale)
	// - handler
	hdCustomer := handler.NewCustomersDefault(svCustomer)
	hdProduct := handler.NewProductsDefault(svProduct)
	hdInvoice := handler.NewInvoicesDefault(svInvoice)
	hdSale := handler.NewSalesDefault(svSale)

	// routes
	// - router
	a.router = chi.NewRouter()
	// - middlewares
	a.router.Use(middleware.Logger)
	a.router.Use(middleware.Recoverer)
	// - endpoints
	a.router.Route("/customers", func(r chi.Router) {
		// - GET /customers
		r.Get("/", hdCustomer.GetAll())
		// - POST /customers
		r.Post("/", hdCustomer.Create())

		r.Get("/top-buyers", hdCustomer.GetTopBuyers())
	})
	a.router.Route("/products", func(r chi.Router) {
		// - GET /products
		r.Get("/", hdProduct.GetAll())
		// - POST /products
		r.Post("/", hdProduct.Create())

		r.Get("/best-sellers", hdProduct.GetTopFiveBestSellers())
	})
	a.router.Route("/invoices", func(r chi.Router) {
		// - GET /invoices
		r.Get("/", hdInvoice.GetAll())
		// - POST /invoices
		r.Post("/", hdInvoice.Create())
		r.Get("/total-by-conditions", hdInvoice.GetTotalAndCustomerCondition())
	})
	a.router.Route("/sales", func(r chi.Router) {
		// - GET /sales
		r.Get("/", hdSale.GetAll())
		// - POST /sales
		r.Post("/", hdSale.Create())
	})

	return
}

// Run runs the application.
func (a *ApplicationDefault) Run() (err error) {
	defer a.appConfig.db.Close()

	err = http.ListenAndServe(a.cfgAddr, a.router)
	return
}
