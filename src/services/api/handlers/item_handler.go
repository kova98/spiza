package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kova98/spiza/services/api/data"
)

type ItemHandler struct {
	l    *log.Logger
	repo *data.ItemRepo
}

func NewItemHandler(l *log.Logger, r *data.ItemRepo) *ItemHandler {
	return &ItemHandler{l, r}
}

func (h *ItemHandler) CreateItem(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Handle POST Item")

	vars := mux.Vars(r)
	idString := vars["id"]
	categoryId, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		http.Error(rw, "Invalid id: "+idString, http.StatusInternalServerError)
	}

	item := &data.Item{}
	err = data.FromJSON(item, r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}
	item.CategoryId = categoryId

	id, err := h.repo.CreateItem(item)
	if err != nil {
		h.l.Println(err)
		http.Error(rw, "Unable to create item", http.StatusInternalServerError)
	}

	rw.WriteHeader(http.StatusCreated)

	item.Id = id
	data.ToJSON(item, rw)
}

func (h *ItemHandler) DeleteItem(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Handle DELETE Item")

	vars := mux.Vars(r)
	idString := vars["id"]
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		http.Error(rw, "Invalid id: "+idString, http.StatusInternalServerError)
	}

	err = h.repo.DeleteItem(id)
	if err != nil {
		h.l.Println(err)
		http.Error(rw, "Unable to delete item", http.StatusInternalServerError)
	}
}
