package products

import (
	"encoding/json"
	"net/http"
)

type prod_ctrl struct {
	repo *prod_repo
}

func NewCtrl(repo *prod_repo) *prod_ctrl {
	return &prod_ctrl{repo}
}

func (c *prod_ctrl) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	vars := r.URL.Query()
	msg := vars["message"][0]

	respone := c.repo.FindAll(msg)

	json.NewEncoder(w).Encode(respone)
	return
}
