// Simple application healthiness service
package statez

import (
	"sync"
)

// Statez is the application wide Service registry
type Statez struct {
	name string
	// Registry
	registry   []ServiceIf
	registryMu sync.RWMutex
}

func NewStatez(name string) *Statez {
	return &Statez{
		name,
		make([]ServiceIf, 0),
		sync.RWMutex{},
	}
}

type ServiceIf interface {
	GetState() ServiceState
	GetName() string
	// IsHealthy() int for now disabled
}

// RegisterService adds Services to the application registry
func (s *Statez) RegisterService(svc ...ServiceIf) {
	s.registryMu.Lock()
	defer s.registryMu.Unlock()
	s.registry = append(s.registry, svc...)
}

// GetReadyState checks the registered services and returns false if any one of the services has StateNotReady set. Services with a StateIgnore set are skipped
func (s *Statez) GetReadyState() bool {
	s.registryMu.RLock()
	defer s.registryMu.RUnlock()

	for _, svc := range s.registry {
		if svc.GetState() == StateIgnore {
			continue
		} else if svc.GetState() == StateNotReady {
			return false
		}
	}
	return true
}
