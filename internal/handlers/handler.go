package handlers

import (
	"net/http"
	"encoding/json"
	"log"
	"ramengo-go/internal/models"
	"ramengo-go/internal/client"
)


func GetBroths(w http.ResponseWriter, r *http.Request) {

	body, err := client.GetRequest("broths")
	if err != nil {
		log.Printf("Erro ao fazer requisição para a RV: %s", err)
		http.Error(w, "could not send request", http.StatusInternalServerError)
		return
	}
	defer body.Close()

	var broths []models.Broth

	if err := json.NewDecoder(body).Decode(&broths); err != nil {
		log.Printf("Erro ao encodar a resposta da requisição: %s", err)
		http.Error(w, "could not send request", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(broths)
}

func GetProteins(w http.ResponseWriter, r *http.Request) {

	body, err := client.GetRequest("proteins")
	if err != nil {
		log.Printf("Erro ao fazer requisição para a RV: %s", err)
		http.Error(w, "could not send request", http.StatusInternalServerError)
		return
	}
	defer body.Close()

	var proteins []models.Protein

	if err := json.NewDecoder(body).Decode(&proteins); err != nil {
		log.Printf("Erro ao encodar a resposta da requisição: %s", err)
		http.Error(w, "could not send request", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(proteins)
}

func PostOrder(w http.ResponseWriter, r *http.Request) {

	body, err := client.PostRequest("order", r.Body)
	if err != nil {
		log.Printf("Erro ao fazer requisição para a RV: %s", err)
		http.Error(w, "could not place order", http.StatusInternalServerError)
		return
	}
	defer body.Close()

	var order models.Order

	if err := json.NewDecoder(body).Decode(&order); err != nil {
		log.Printf("Erro ao encodar a resposta da requisição: %s", err)
		http.Error(w, "could not place order", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(order)
}
