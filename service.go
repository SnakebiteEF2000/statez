package statez

import "sync/atomic"

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

func (s *ServiceHandler) StatusReady() {
	atomic.StoreInt32(&s.Ready, 1)
}

func (s *ServiceHandler) StatusNotReady() {
	atomic.StoreInt32(&s.Ready, 0)
}

func (s *ServiceHandler) StatusIgnore() {
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
