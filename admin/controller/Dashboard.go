package controller

import (
	"blogproject/admin/helpers"
	"blogproject/admin/models"
	"fmt"
	"github.com/gosimple/slug"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"
)

type Dashboard struct {
}

func (dashboard Dashboard) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles(helpers.Include("dashboard/list")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Posts"] = models.Post{}.GetAll()
	view.ExecuteTemplate(w, "index", data)
}

func (dashboard Dashboard) NewItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles(helpers.Include("dashboard/add")...)
	if err != nil {
		fmt.Println(err)
		return
	}

	view.ExecuteTemplate(w, "index", nil)
}

func (dashboard Dashboard) Add(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	title := r.FormValue("blog-title")
	slug := slug.Make(title)
	description := r.FormValue("blog-desc")
	categoryID, _ := strconv.Atoi(r.FormValue("category"))
	content := r.FormValue("blog-content")

	//upload
	r.ParseMultipartForm(10 << 20)
	file, header, err := r.FormFile("blog-picture")
	if err != nil {
		fmt.Println(err)
		return
	}
	f, err := os.OpenFile("uploads/"+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = io.Copy(f, file)
	//Upload End
	if err != nil {
		fmt.Println(err)
		return
	}

	models.Post{
		Title:       title,
		Slug:        slug,
		Description: description,
		CategoryID:  categoryID,
		Content:     content,
		PictureURL:  "uploads/" + header.Filename,
	}.Add()
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
	//TODO alert
}
