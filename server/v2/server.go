package server

import (
	"fmt"
	"net/http"

	store "github.com/Natan5533/server/v2/ports"
)

func Server(store store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, store.Fetch())
	}
}
