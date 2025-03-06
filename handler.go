package statez

import (
	"encoding/json"
	"net/http"
)

// write a /ready and /healthz handler that first of all handles the checks correctly and expose 200/503

func (s *Statez) ReadynessHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	isReady := s.CheckServiceReadyState()
	if isReady {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
	}

	type Response struct {
		RegisteredServices []ServiceHandler `json:"registered_services"`
		SystemReady        bool             `json:"system_ready"`
	}

	services := make([]ServiceHandler, 0, len(s.registry))
	s.registryMu.RLock()
	for _, v := range s.registry {
		services = append(services, v.GetInfo())
	}
	s.registryMu.RUnlock()

	response := Response{
		RegisteredServices: services,
		SystemReady:        isReady,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
