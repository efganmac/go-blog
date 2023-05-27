package config

import (
	admin "blogproject/admin/controller"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Routes() *httprouter.Router {
	r := httprouter.New()
	//Admin
	//Blog Posts
	r.GET("/admin", admin.Dashboard{}.Index)
	r.GET("/admin/add-new", admin.Dashboard{}.NewItem)
	r.POST("/admin/add", admin.Dashboard{}.Add)
	//Serve Files
	r.ServeFiles("/admin/assets/*filepath", http.Dir("admin/assets"))
	return r
}
