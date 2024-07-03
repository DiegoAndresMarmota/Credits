package memory

import (
	"context"
	"errors"
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

// Services verifica si los servicios estan dispuestos para el servicio.
func (re *Registry) Online(instanceID string, service string) error {
	re.Lock()
	defer re.Unlock()

	if _, ok := re.serviceAddrs[Service(service)]; !ok {
		return errors.New("el servicio no esta disponible")
	}
	if _, ok := re.serviceAddrs[Service(service)][InstanceID(instanceID)]; !ok {
		return errors.New("la instancia no esta disponible")
	}
	if !re.serviceAddrs[Service(service)][InstanceID(instanceID)].Active {
		return errors.New("la instancia no esta disponible")
	}
	re.serviceAddrs[Service(service)][InstanceID(instanceID)].LastTime = time.Now()
	return nil
}

// ListServices lista los servicios que estan disponibles, segun puerto, disponibilidad y hora de verificación.
func (re *Registry) ListServices(ctx context.Context, service string) ([]string, error) {
	re.RLock()
	defer re.RUnlock()

	if len(re.serviceAddrs[Service(service)]) == 0 {
		return nil, errors.New("el servicio no esta disponible")
	}

	var res []string
	for _, i := range re.serviceAddrs[Service(service)] {
		if i.LastTime.Before(time.Now().Add(-5 * time.Second)) {
			continue
		}
		res = append(res, i.Host)
	}
	return res, nil
}
