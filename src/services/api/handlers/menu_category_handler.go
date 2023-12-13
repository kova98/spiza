package handlers

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"

	"github.com/kova98/spiza/services/api/data"
)

type MenuCategoryHandler struct {
	l    *log.Logger
	repo *data.MenuCategoryRepo
}

func NewMenuCategoryHandler(l *log.Logger, r *data.MenuCategoryRepo) *MenuCategoryHandler {
	return &MenuCategoryHandler{l, r}
}

func (mch *MenuCategoryHandler) CreateMenuCategory(rw http.ResponseWriter, r *http.Request) {
	mch.l.Println("Handle POST MenuCategory")

	menuCategory := &data.MenuCategory{}
	err := data.FromJSON(menuCategory, r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
		return
	}

	id, err := mch.repo.CreateMenuCategory(menuCategory)
	if err != nil {
		mch.l.Println(err)
		http.Error(rw, "Unable to create menu category", http.StatusInternalServerError)
		return
	}

	menuCategory.Id = id
	err = data.ToJSON(menuCategory, rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (mch *MenuCategoryHandler) DeleteMenuCategory(rw http.ResponseWriter, r *http.Request) {
	mch.l.Println("Handle DELETE MenuCategory")

	vars := mux.Vars(r)
	idString := vars["id"]
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		mch.l.Println(err)
		http.Error(rw, "Invalid id: "+idString, http.StatusInternalServerError)
		return
	}

	if err := mch.repo.DeleteMenuCategory(id); err != nil {
		mch.l.Println(err)
		http.Error(rw, "Unable to delete menu category", http.StatusInternalServerError)
	}
}
