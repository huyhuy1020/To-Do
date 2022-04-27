package models

import (
	"fmt"
	"net/http"
)

type Employee struct { //item
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
type Employee__ struct { //itemslist
	Em []Employee `json:"Employee"`
}

func (i *Employee) Bind(r *http.Request) error {
	if i.Name == "" {
		return fmt.Errorf("name is a required field")
	}
	return nil
}
func (*Employee__) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (*Employee) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
