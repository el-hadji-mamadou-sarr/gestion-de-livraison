package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test de la livraison réussie avec un camion
func TestDeliverPackage_Truck_Success(t *testing.T) {
	truck := Truck{ID: "T123", Capacity: 10}

	destination := "Thiès"
	result, err := truck.DeliverPackage(destination)

	assert.NoError(t, err)
	assert.Contains(t, result, "Colis livré à Thiès")
}

// Test du statut du camion
func TestGetStatus_Truck(t *testing.T) {
	truck := Truck{ID: "T123", Capacity: 15}

	status := truck.GetStatus()

	expectedStatus := " 🚛 Camion: T123 En route (Capacité: 15)"
	assert.Equal(t, expectedStatus, status)
}
