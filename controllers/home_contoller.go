package controllers

import (
	"bpk-dx/controllers/render"
	"fmt"
	"net/http"
)

type Home struct{}

func (h Home) HomeMain(resp http.ResponseWriter, r *http.Request) {
	if err := render.View(resp, "home_main", nil); err != nil {
		fmt.Fprintln(resp, "HomeMain error!", err)
	}
}
