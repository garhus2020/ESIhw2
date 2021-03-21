package handler

import (
	"encoding/json"
	"fmt"
	"github.com/garhus2020/ESIhw2/plant/pkg/domain"
	"github.com/garhus2020/ESIhw2/plant/pkg/repository"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type PlantHandler struct {
	plantmRepository *repository.PlantmRepository
	plantRepository  *repository.PlantRepository
	cacheRepository  *repository.CacheRepository
}

func NewPlantHandler(plantmRepository *repository.PlantmRepository, plantRepository *repository.PlantRepository, cacheRepository *repository.CacheRepository) *PlantHandler {
	return &PlantHandler{
		plantmRepository: plantmRepository,
		plantRepository:  plantRepository,
		cacheRepository:  cacheRepository,
	}
}

func (h *PlantHandler) CacheRequest(r *http.Request) (string, error) {
	log.Printf("handling request %v", r)
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "Nouuuuu", fmt.Errorf("could not insert request, err %v", err)
	}
	bs := string(b)
	result := &domain.Request{Body: bs}
	_, err = h.cacheRepository.Create(result)
	if err != nil {
		return "Nouuuuu", fmt.Errorf("could not insert request, err %v", err)
	}
	return "Done", nil
}

func (h *PlantHandler) CreatePlantm(w http.ResponseWriter, r *http.Request) {
	log.Printf("received request %v", r)
	plantm := &domain.Plantm{}
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

	plant := &domain.Plant{}
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
	answer, err := h.CacheRequest(r)
	w.Write([]byte(answer))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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
	req_cache := r
	var data map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	i := data["ident"].(string)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	start, _ := strconv.Atoi(data["start"].(string))
	end, _ := strconv.Atoi(data["end"].(string))
	price, err := h.plantRepository.GetOne(i)
	p, _ := strconv.Atoi(price)
	if err != nil {
		price, err := h.plantmRepository.GetOne(i)
		p, _ := strconv.Atoi(price)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		res := &domain.Cost{Ident: i, Price: strconv.Itoa(p * (end - start))} //// calculating perdiod mult by price
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(&res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	s := strconv.Itoa(p * (end - start))
	res := &domain.Cost{Ident: i, Price: s}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	_, err = h.CacheRequest(req_cache)
	//w.Write([]byte(answer))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (h *PlantHandler) GetCache(w http.ResponseWriter, r *http.Request) {
	log.Printf("received request %v", r)
	requests, err := h.cacheRepository.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// write success response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&requests)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// _, err = h.CacheRequest(r)
	// //w.Write([]byte(answer))
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
}
