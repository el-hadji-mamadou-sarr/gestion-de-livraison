package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test de la livraison réussie avec une batterie suffisante
func TestDeliverPackage_Success(t *testing.T) {
	drone := Drone{ID: "D123", Battery: 50}

	destination := "Dakar"
	result, err := drone.DeliverPackage(destination)

	assert.NoError(t, err)
	assert.Contains(t, result, "Colis livré à Dakar")
}

// Test de la livraison échouée à cause d'une batterie faible
func TestDeliverPackage_LowBattery(t *testing.T) {
	drone := Drone{ID: "D123", Battery: 10}

	_, err := drone.DeliverPackage("Dakar")

	assert.Error(t, err)
	assert.Equal(t, "⚡ Batterie trop faible pour décoller", err.Error())
}

// Test du statut du drone
func TestGetStatus(t *testing.T) {
	drone := Drone{ID: "D123", Battery: 75}

	status := drone.GetStatus()

	expectedStatus := "🚀 Drone : D123 - Batterie à 75%"
	assert.Equal(t, expectedStatus, status)
}
