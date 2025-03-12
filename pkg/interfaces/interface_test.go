package interfaces

import (
	"testing"

	"github.com/el-hadji-mamadou-sarr/gestion-de-livraison.git/pkg/models"
	"github.com/stretchr/testify/assert"
)

// Test de l'implémentation de l'interface TransportMethod par Truck
func TestTruck_ImplementsTransportMethod(t *testing.T) {
	var transport TransportMethod = &models.Truck{ID: "T123", Capacity: 10}
	assert.NotNil(t, transport)
}

// Test de l'implémentation de l'interface TransportMethod par Drone
func TestDrone_ImplementsTransportMethod(t *testing.T) {
	var transport TransportMethod = &models.Drone{ID: "D456", Battery: 100}
	assert.NotNil(t, transport)
}

// Test de l'implémentation de l'interface TransportMethod par Boat
func TestBoat_ImplementsTransportMethod(t *testing.T) {
	var transport TransportMethod = &models.Boat{ID: "B789", Weather: "Clear"}
	assert.NotNil(t, transport)
}
