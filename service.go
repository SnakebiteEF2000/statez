package statez

import (
	"encoding/json"
	"sync/atomic"
)

const (
	StateNotReady = 0
	StateReady    = 1
	StateIgnore   = -1
)

type Service interface {
	GetName() string
	IsReady() int
	GetInfo() ServiceHandler
	// IsHealthy() int for now disabled
}

type ServiceHandler struct {
	Name  string `json:"name"`
	Ready int32  `json:"state"` // 0 = no ready; 1 = ready; -1 = ignore state
}

func (s *ServiceHandler) StateReady() {
	atomic.StoreInt32(&s.Ready, 1)
}

func (s *ServiceHandler) StateNotReady() {
	atomic.StoreInt32(&s.Ready, 0)
}

func (s *ServiceHandler) StateIgnore() {
	atomic.StoreInt32(&s.Ready, -1)
}

func (s *ServiceHandler) IsReady() int {
	return int(atomic.LoadInt32(&s.Ready))
}

func (s *ServiceHandler) GetName() string {
	return s.Name
}

func (s *ServiceHandler) GetInfo() ServiceHandler {
	return *s
}

func NewServiceHandlerWithOpts(name string) *ServiceHandler { // add the opts part not just a name variable
	return &ServiceHandler{Name: name, Ready: 0}
}

func (s *ServiceHandler) MarshalJSON() ([]byte, error) {
	type Alias ServiceHandler

	var readableState string
	switch s.Ready {
	case StateReady:
		readableState = "ready"
	case StateNotReady:
		readableState = "not_ready"
	case StateIgnore:
		readableState = "ignored"
	default:
		readableState = "unknown"
	}

	return json.MarshalIndent(&struct {
		*Alias
		Ready string `json:"state"`
	}{
		Alias: (*Alias)(s),
		Ready: readableState,
	}, "", "  ")
}
