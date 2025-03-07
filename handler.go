package statez

import (
	"encoding/json"
	"net/http"
)

// write a /ready and /healthz handler that first of all handles the checks correctly and expose 200/503

func (s *Statez) ReadynessHandler(w http.ResponseWriter, _ *http.Request) {
	if len(s.registry) <= 0 {
		http.Error(w, "no service registered", http.StatusInternalServerError)
		return
	}

	serviceState := s.CheckServiceReadyState()

	if serviceState {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
	}

	type mslService struct {
		Name   string       `json:"name"`
		Status ServiceState `json:"status"`
	}

	dat := struct {
		Application      string       `json:"app"`
		Services         []mslService `json:"services"`
		ApplicationReady bool         `json:"application_ready"`
	}{
		Application:      s.Name,
		ApplicationReady: serviceState,
	}

	for _, v := range s.registry {
		var newMslService mslService
		newMslService.Name = v.GetName()
		newMslService.Status = v.GetState()
		dat.Services = append(dat.Services, newMslService)
	}

	if err := json.NewEncoder(w).Encode(dat); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

}
