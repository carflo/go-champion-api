package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

type championHandler struct {
	sync.Mutex
	store map[string]Champion
}

func (h *championHandler) champions(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.get(w, r)
		return
	case "POST":
		h.post(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
	}
}

func (h *championHandler) get(w http.ResponseWriter, r *http.Request) {
	champions := []Champion{}
	for _, champion := range h.store {
		champions = append(champions, champion)
	}

	result, err := json.Marshal(champions)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("content-type", "application/json")
	w.Write(result)
}

func (h *championHandler) getChampion(w http.ResponseWriter, r *http.Request) {
	urlParts := strings.Split(r.URL.String(), "/")
	if len(urlParts) != 3 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	id := urlParts[2]
	champion, ok := h.store[id]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	result, err := json.Marshal(champion)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("content-type", "application/json")
	w.Write([]byte(result))
}

func (h *championHandler) post(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	reqContentType := r.Header.Get("content-type")
	if reqContentType != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte(fmt.Sprintf("need content-type 'application/json', but got '%s'", reqContentType)))
		return
	}

	newChampion := &Champion{
		ID: fmt.Sprintf("%d", time.Now().UnixNano()),
	}

	err = json.Unmarshal(reqBody, newChampion)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	h.Lock()
	h.store[newChampion.ID] = *newChampion
	h.Unlock()
	w.WriteHeader(http.StatusCreated)
}

func main() {
	championHandler := &championHandler{
		store: map[string]Champion{},
	}

	http.HandleFunc("/champions", championHandler.champions)
	http.HandleFunc("/champions/", championHandler.getChampion)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Panic(err)
	}
}
