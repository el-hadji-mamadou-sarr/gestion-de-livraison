package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test de la livraison réussie avec météo favorable
func TestDeliverPackage_Boat_Success(t *testing.T) {
	boat := Boat{ID: "B123", Weather: "Clear"}

	destination := "Gorée"
	result, err := boat.DeliverPackage(destination)

	assert.NoError(t, err)
	assert.Contains(t, result, "Colis livré à Gorée")
}

// Test de la livraison échouée à cause de mauvaises conditions météo
func TestDeliverPackage_Boat_BadWeather(t *testing.T) {
	boat := Boat{ID: "B123", Weather: "Stormy"}

	_, err := boat.DeliverPackage("Gorée")

	assert.Error(t, err)
	assert.Equal(t, "conditions météo défavorables", err.Error())
}

// Test du statut du bateau
func TestGetStatus_Boat(t *testing.T) {
	boat := Boat{ID: "B123", Weather: "Clear"}

	status := boat.GetStatus()

	expectedStatus := "Bateau: B123 - Météo actuelle: Clear"
	assert.Equal(t, expectedStatus, status)
}
