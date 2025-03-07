package statez

import (
	"sync/atomic"
)

// Service is the shipped implementation of the ServiceIF interface that Statez requires
type Service struct {
	Name  string `json:"name"`
	Ready int32  `json:"state"` // 0 = no ready; 1 = ready; -1 = ignore state
}

func NewService(name string) *Service { // add the opts part not just a name variable
	return &Service{Name: name, Ready: 0}
}

// StateReady sets the ready value to 1
func (s *Service) StateReady() {
	atomic.StoreInt32(&s.Ready, 1)
}

// StateNotReady sets the ready value to 0
func (s *Service) StateNotReady() {
	atomic.StoreInt32(&s.Ready, 0)
}

// StateIgnore sets the ready value to -1
func (s *Service) StateIgnore() {
	atomic.StoreInt32(&s.Ready, -1)
}

// GetState retrives the current state as ServiceState
func (s *Service) GetState() ServiceState {
	return ServiceState(atomic.LoadInt32(&s.Ready))
}

// GetName returns the name of the Service
func (s *Service) GetName() string {
	return s.Name
}

func (s ServiceState) MarshalText() ([]byte, error) {
	return []byte(s.String()), nil
}
