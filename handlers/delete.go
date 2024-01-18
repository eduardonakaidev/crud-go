package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/eduardonakaidev/crud-go/models"
	"github.com/go-chi/chi/v5"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r,"id"))
	if err != nil{
		log.Printf("Erro ao fazer parse do json: %v",err)
		http.Error(w,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
		return
	}

	rows, err := models.Delete(int64(id))
	if err != nil {
		log.Printf("Erro ao remover registro: %v",err)
		http.Error(w,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
		return
	}
	if rows > 1 {
		log.Printf("Error: foram removidos %d registros",rows)
	}
	resp := map[string]any{
		"Error": false,
		"Message": "registro removidos com sucesso",
	}
	w.Header().Add("Content-Type","application/json")
	json.NewEncoder(w).Encode(resp)

}