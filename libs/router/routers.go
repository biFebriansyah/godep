package router

import (
	"github.com/biFebriansyah/demo_flyio/module/v1/products"
	"github.com/gorilla/mux"
)

func NewApp() *mux.Router {
	mainRoute := mux.NewRouter()

	products.New(mainRoute)
	return mainRoute
}
