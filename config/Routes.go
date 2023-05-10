package config

import (
	admin "blogproject/admin/controller"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Routes() *httprouter.Router {
	r := httprouter.New()
	r.GET("/admin", admin.Dashboard{}.Index)

	r.ServeFiles("/admin/assets/*filepath", http.Dir("admin/assets"))
	return r
}
