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
	r.GET("/admin/delete/:id", admin.Dashboard{}.Delete)
	r.GET("/admin/edit/:id", admin.Dashboard{}.Edit)
	r.POST("/admin/update/:id", admin.Dashboard{}.Update)
	//Serve Files
	r.ServeFiles("/admin/assets/*filepath", http.Dir("admin/assets"))
	r.ServeFiles("/uploads/*filepath", http.Dir("uploads"))

	//Userops
	r.GET("/admin/login", admin.Userops{}.Index)
	return r
}
