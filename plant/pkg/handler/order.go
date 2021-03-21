package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/garhus2020/ESIhw2/plant/pkg/repository"
	"github.com/garhus2020/ESIhw2/plant/pkg/domain"
)

type OrderHandler struct {
	orderRepository *repository.OrderRepository
}

func NewOrderHandler(orderRepository *repository.OrderRepository) *OrderHandler {
	return &OrderHandler{
		orderRepository: orderRepository,
	}
}

func (h *OrderHandler) GetStatus(w http.ResponseWriter, r *http.Request) {
	log.Printf("received request %v", r)

	var data map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	i := data["ident"].(string)
	s := data["start"].(string)
	e := data["end"].(string)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	num, err := h.orderRepository.GetNumIntersects(i, s, e)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

	if num == "0" {
		w.Write([]byte("available"))
		response := &domain.Check{Availability: "available"}
		err = json.NewEncoder(w).Encode(&response)
	} else {
		w.Write([]byte("Not available"))
		response := &domain.Check{Availability: "Not available"}
		err = json.NewEncoder(w).Encode(&response)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
