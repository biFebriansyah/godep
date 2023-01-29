package products

import "fmt"

type prod_repo struct {
	name string
}

func NewRepo(name string) *prod_repo {
	return &prod_repo{name}
}

func (r *prod_repo) FindAll(msg string) string {
	return fmt.Sprintf("the message %s and the string %s", msg, r.name)
}
