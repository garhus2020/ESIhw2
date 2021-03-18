package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type PlantHandler struct {
	plantmRepository *PlantmRepository
	plantRepository  *PlantRepository
}

func NewPlantHandler(plantmRepository *PlantmRepository, plantRepository *PlantRepository) *PlantHandler {
	return &PlantHandler{
		plantmRepository: plantmRepository,
		plantRepository:  plantRepository,
	}
}

func (h *PlantHandler) CreatePlantm(w http.ResponseWriter, r *http.Request) {
	log.Printf("received request %v", r)
	plantm := &Plantm{}
	err := json.NewDecoder(r.Body).Decode(plantm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	createdPlantm, err := h.plantmRepository.Create(plantm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(&createdPlantm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *PlantHandler) CreatePlant(w http.ResponseWriter, r *http.Request) {
	log.Printf("received request %v", r)
	plant := &Plant{}
	err := json.NewDecoder(r.Body).Decode(plant)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	createdPlant, err := h.plantRepository.Create(plant)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(&createdPlant)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *PlantHandler) GetPlants(w http.ResponseWriter, r *http.Request) {
	log.Printf("received request %v", r)
	plantms, err := h.plantmRepository.GetAll()
	plants, _ := h.plantRepository.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&plantms)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = json.NewEncoder(w).Encode(&plants)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func (h *PlantHandler) GetPrice(w http.ResponseWriter, r *http.Request) {
	log.Printf("received request %v", r)

	var data map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	i := data["ident"].(string)
	//_, err = fmt.Sscan(data["ident"].(string), &i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	price, err := h.plantRepository.GetOne(i)
	if err != nil {
		w.Write([]byte("feck it"))
		price, err := h.plantmRepository.GetOne(i)
		w.Write([]byte(price))
		if err != nil {
			w.Write([]byte("feck it"))

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		// return
	}

	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(&price)

	w.Write([]byte(price))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
