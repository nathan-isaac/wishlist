package server

import (
	"net/http"
	"wishlist/internal/views"
)

func HelloWebHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	name := r.FormValue("name")
	component := views.HelloPost(name)
	component.Render(r.Context(), w)
}
