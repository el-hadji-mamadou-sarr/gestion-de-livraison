package factory

import (
	"testing"

	"github.com/el-hadji-mamadou-sarr/gestion-de-livraison.git/pkg/interfaces"
	"github.com/stretchr/testify/assert"
)

// Test de la création d'un camion
func TestGetTransportMethod_Truck(t *testing.T) {
	transport, err := GetTransportMethod("truck")

	assert.NoError(t, err)
	assert.NotNil(t, transport)
	assert.Implements(t, (*interfaces.TransportMethod)(nil), transport)
}

// Test de la création d'un drone
func TestGetTransportMethod_Drone(t *testing.T) {
	transport, err := GetTransportMethod("drone")

	assert.NoError(t, err)
	assert.NotNil(t, transport)
	assert.Implements(t, (*interfaces.TransportMethod)(nil), transport)
}

// Test de la création d'un bateau
func TestGetTransportMethod_Boat(t *testing.T) {
	transport, err := GetTransportMethod("boat")

	assert.NoError(t, err)
	assert.NotNil(t, transport)
	assert.Implements(t, (*interfaces.TransportMethod)(nil), transport)
}

// Test d'une méthode inconnue
func TestGetTransportMethod_Unknown(t *testing.T) {
	transport, err := GetTransportMethod("plane")

	assert.Error(t, err)
	assert.Nil(t, transport)
	assert.Equal(t, "unknown transport method", err.Error())
}
