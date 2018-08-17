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
		w.Write([]byte("\n<html>\n<head>\n<title>Idea Dashboard</title>\n<meta property=\"og:image\" content=\"https://idea-dash.com/static/naveed.png\" /> <meta name=\"twitter:card\" content=\"image_of_thing\" ></head>\n<body>\n<br /><br /><br /><marquee><h1 style=\"color: red;\">Hello Naveed-san.</h1></marquee>\n</body>\n</html> "))
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
