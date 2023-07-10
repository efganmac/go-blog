package controller

import (
	"blogproject/admin/helpers"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

type Userops struct {
}

func (userops Userops) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles(helpers.Include("/userops/login")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	view.ExecuteTemplate(w, "index", nil)
}
