package healthz

import "net/http"

func HandlerError(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusBadRequest, "Something went wrong")
}
