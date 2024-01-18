package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/eduardonakaidev/crud-go/models"
	"github.com/go-chi/chi/v5"
)

func Get(w http.ResponseWriter,r *http.Request){
	id,err := strconv.Atoi(chi.URLParam(r,"id"))
	if err != nil {
		log.Printf("Erro ao fazer o parse do id: %v",err)
		http.Error(w,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
		return
	}
	todo,err := models.Get(int64(id))
	if err != nil {
		log.Printf("Erro ao procurar o id o registro: %v",err)
		http.Error(w,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type","application/json")
	json.NewEncoder(w).Encode(todo)
}