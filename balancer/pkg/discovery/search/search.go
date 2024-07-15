package search

import (
	"context"
	"errors"
	"strconv"
	"strings"

	consul "github.com/hashicorp/consul/api"
)

// Registry define el servicio de registro basado en Hashicorp/Consul.
type Registry struct {
	client consul.Client
}

// NewRegistry crea una nueva instancia de un servicio de registro basado en Hashicorp/Consul.
func NewRegistry(address string) (*Registry, error) {
	config := consul.DefaultConfig()
	config.Address = address
	client, err := consul.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &Registry{client: *client}, nil
}

// Register crea un servicio en el registro.
func (re *Registry) Register(ctx context.Context, instanceID string, serviceName string, hostPort string) error {
	split := strings.Split(hostPort, ":")
	if len(split) != 2 {
		return errors.New("hostPort must be in a form of host/port, example: localhost:8081")
	}
	port, err := strconv.Atoi(split[1])
	if err != nil {
		return err
	}
	return re.client.Agent().ServiceRegister(&consul.AgentServiceRegistration{
		Address: split[0],
		ID:      instanceID,
		Name:    serviceName,
		Port:    port,
		Check: &consul.AgentServiceCheck{
			CheckID: instanceID,
			Name:    "Check",
			TTL:     "10s",
		},
	})
}

// Deregister elimina un servicio del registro.
func (re *Registry) Deregister(ctx context.Context, instanceID string, _ string) error {
	return re.client.Agent().ServiceDeregister(instanceID)
}

// ServiceAddresses retorna una lista de direcciones de instancias activas para un servicio dado.
func (re *Registry) ServiceAddresses(ctx context.Context, serviceName string) ([]string, error) {
	entries, _, err := re.client.Health().Service(serviceName, "", true, nil)
	if err != nil {
		return nil, err
	} else if len(entries) == 0 {
		return nil, errors.New("service not found")
	}
	var res []string
	for _, entry := range entries {
		res = append(res, entry.Service.Address+":"+strconv.Itoa(entry.Service.Port))
	}
	return res, nil
}

// ReportHeathyState verifica el estado de salud de reporte de cada servicio.
func (re *Registry) ReportHeathyState(InstanceID string, _ string) error {
	return re.client.Agent().PassTTL(InstanceID, "")
}
