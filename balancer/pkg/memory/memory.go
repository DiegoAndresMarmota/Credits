package memory

import (
	"context"
	"sync"
	"time"
)

type Service string

type InstanceID string

// ServiceInstance define una instancia de un servicio
type ServiceInstance struct {
	Host     string
	Active   bool
	LastTime time.Time
}

// Registry define un registro de servicio en memoria
type Registry struct {
	sync.RWMutex
	serviceAddrs map[Service]map[InstanceID]*ServiceInstance
}

// NewRegistry crea un nuevo registro de servicio
func NewRegistry() *Registry {
	return &Registry{
		serviceAddrs: make(map[Service]map[InstanceID]*ServiceInstance),
	}
}

// Register registra una instancia de un servicio
func (re *Registry) Register(ctx context.Context, service string, instanceID string, host string) error {
	re.Lock()
	defer re.Unlock()

	if _, ok := re.serviceAddrs[Service(service)]; !ok {
		re.serviceAddrs[Service(service)] = make(map[InstanceID]*ServiceInstance)
	}
	re.serviceAddrs[Service(service)][InstanceID(instanceID)] = &ServiceInstance{
		Host:     host,
		Active:   true,
		LastTime: time.Now(),
	}
	return nil
}

// Deregister desregistra una instancia de un servicio
func (re *Registry) Deregister(ctx context.Context, service string, instanceID string) error {
	re.Lock()
	defer re.Unlock()

	if _, ok := re.serviceAddrs[Service(service)]; ok {
		delete(re.serviceAddrs[Service(service)], InstanceID(instanceID))
		return nil
	}
	return nil
}
