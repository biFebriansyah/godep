package products

import "github.com/gorilla/mux"

func New(rt *mux.Router) {
	route := rt.PathPrefix("/product").Subrouter()

	repo := NewRepo("ebiebi")
	ctrl := NewCtrl(repo)

	route.HandleFunc("", ctrl.Get).Methods("GET")
}
