package memory

import (
	"sync"
	"time"
)

type service string

type instanceID string

// ServiceInstance define una instancia de un servicio
type ServiceInstance struct {
	Host     string
	Active   bool
	LastTime time.Time
}

// Registry define un registro de servicio en memoria
type Registry struct {
	sync.RWMutex
	serviceAddrs map[service]map[instanceID]*ServiceInstance
}

// NewRegistry crea un nuevo registro de servicio
func NewRegistry() *Registry {
	return &Registry{
		serviceAddrs: make(map[service]map[instanceID]*ServiceInstance),
	}
}
