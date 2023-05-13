package main

import (
	admin_models "blogproject/admin/models"
	"blogproject/config"
	"net/http"
)

func main() {
	admin_models.Post{}.Migrate()
	http.ListenAndServe(":8080", config.Routes())
}
