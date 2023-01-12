package handler

import (
	"encoding/json"
	makves "github.com/cucumberjaye/makves_testovoe"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (h *Handler) GetItem(w http.ResponseWriter, r *http.Request) {
	strId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		http.Error(w, "id is empty or not a number", http.StatusBadRequest)
		return
	}

	var item makves.Item
	item, err = h.service.GetItem(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := json.Marshal(item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(output)
}

func (h *Handler) GetItems(w http.ResponseWriter, r *http.Request) {
	strStartId := chi.URLParam(r, "start_id")
	strEndId := chi.URLParam(r, "end_id")
	startId, err1 := strconv.Atoi(strStartId)
	endId, err2 := strconv.Atoi(strEndId)
	if err1 != nil || err2 != nil {
		http.Error(w, "id is empty or not a number", http.StatusBadRequest)
		return
	}

	var items []makves.Item
	items, err := h.service.GetItems(startId, endId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := json.Marshal(items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(output)
}
