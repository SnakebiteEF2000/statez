package statez

import (
	"encoding/json"
	"net/http"
)

// write a /ready and /healthz handler that first of all handles the checks correctly and expose 200/503

func (s *Statez) ReadynessHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if s.CheckServiceReadyState() {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
	}

	dat := struct {
		Data []ServiceHandler `json:"registered_services"`
	}{}

	for _, v := range s.registry {
		dat.Data = append(dat.Data, v.GetInfo())
	}

	if err := json.NewEncoder(w).Encode(dat); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
