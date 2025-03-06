// Simple application healthiness service
package statez

import (
	"sync"
)

type Statez struct {
	// Registry
	registry   []Service
	registryMu sync.RWMutex
}

func (s *Statez) RegisterService(svc ...Service) {
	s.registryMu.Lock()
	defer s.registryMu.Unlock()
	s.registry = append(s.registry, svc...)
}

func (s *Statez) CheckServiceReadyState() bool {
	s.registryMu.RLock()
	defer s.registryMu.RUnlock()

	for _, svc := range s.registry {
		if svc.IsReady() == -1 {
			// -1 is an ignored service that was registered but is currently disabled i.e. by a sheduler
			continue
		} else if svc.IsReady() == 0 {
			return false
		}
	}
	return true
}
