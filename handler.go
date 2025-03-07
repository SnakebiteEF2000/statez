package statez

import (
	"encoding/json"
	"net/http"
)

func (s *Statez) ReadinessHandler(w http.ResponseWriter, _ *http.Request) {
	if len(s.registry) <= 0 {
		http.Error(w, "no service registered", http.StatusInternalServerError)
		return
	}

	serviceState := s.GetReadyState()

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
		ApplicationReady bool         `json:"application_ready"`
		Services         []mslService `json:"services"`
	}{
		s.name,
		serviceState,
		make([]mslService, 0),
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
