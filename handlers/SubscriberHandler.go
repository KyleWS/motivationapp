package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/KyleWS/motivationapp/models/subscribe"
)

func (ctx *Context) SubHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Println("GET received! No functionality here yet...")
		fmt.Println(r.Header)
		w.Write([]byte("kinex were way better than legos"))
	case "POST":
		ns := &subscribe.NewSubscriber{}
		if err := json.NewDecoder(r.Body).Decode(ns); err != nil {
			http.Error(w, fmt.Sprintf("error decoding JSON: %v", err), http.StatusBadRequest)
			return
		}

		sub, err := ctx.subscriberStore.Insert(ns)
		if err != nil {
			http.Error(w, fmt.Sprintf("error inserting task: %v", err), http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(sub); err != nil {
			http.Error(w, fmt.Sprintf("error encoding response value to JSON: %v", err), http.StatusInternalServerError)
		}
	default:
		http.Error(w, "method must be GET or POST", http.StatusMethodNotAllowed)
		return
	}
}
