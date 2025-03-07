package statez

import (
	"sync/atomic"
)

type ServiceState int32

const (
	StateIgnore ServiceState = iota - 1
	StateNotReady
	StateReady
)

var readyStateName = map[ServiceState]string{
	StateIgnore:   "ignored",
	StateNotReady: "not ready",
	StateReady:    "ready",
}

func (ss ServiceState) String() string {
	return readyStateName[ss]
}

type Service interface {
	GetState() ServiceState
	GetName() string
	// IsHealthy() int for now disabled
}

type StateHandler struct {
	Name  string `json:"name"`
	Ready int32  `json:"state"` // 0 = no ready; 1 = ready; -1 = ignore state
}

func NewServiceHandlerWithOpts(name string) *StateHandler { // add the opts part not just a name variable
	return &StateHandler{Name: name, Ready: 0}
}

func (s *StateHandler) StateReady() {
	atomic.StoreInt32(&s.Ready, 1)
}

func (s *StateHandler) StateNotReady() {
	atomic.StoreInt32(&s.Ready, 0)
}

func (s *StateHandler) StateIgnore() {
	atomic.StoreInt32(&s.Ready, -1)
}

func (s *StateHandler) GetState() ServiceState {
	return ServiceState(atomic.LoadInt32(&s.Ready))
}

func (s *StateHandler) GetName() string {
	return s.Name
}

func (s ServiceState) MarshalText() ([]byte, error) {
	return []byte(s.String()), nil
}
