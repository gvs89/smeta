package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func CreateEstimate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var estimate EstimateRequest
	err := json.NewDecoder(r.Body).Decode(&estimate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pdfFileName := fmt.Sprintf("%s.pdf", estimate.Title)
	pdfGenerator(pdfFileName, estimate)

	sendToWhatsApp(pdfFileName)

	response := map[string]string{"message": "Смета успешно создана"}
	jsonResponse, _ := json.Marshal(response)
	w.Write(jsonResponse)
}

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/create-estimate", CreateEstimate).Methods("POST")
	return r
}
