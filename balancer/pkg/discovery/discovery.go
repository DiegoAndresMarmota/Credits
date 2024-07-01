package discovery

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Registry define un servicio de registro para la creación de las instancias.
type Registry interface {
	Instance(ctx context.Context, instanceID string, serviceName string, hostPort string) error
	NotInstance(ctx context.Context, instanceID string, serviceName string) error
	ServiceAddresses(ctx context.Context, serviceID string) ([]string, error)
	ReportHealthyState(instanceID string, serviceName string) error
}

var ErrNotFound = errors.New("no service addresses found")

func GenerateInstanceID(serviceName string) string {
	return fmt.Sprintf("%s-%d", serviceName, rand.New(rand.NewSource(time.Now().UnixNano())).Int())
}
