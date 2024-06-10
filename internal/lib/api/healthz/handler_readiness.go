package healthz

import "net/http"

func HandlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, struct{}{})
}
