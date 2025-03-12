package utils

import (
	"testing"

	"github.com/el-hadji-mamadou-sarr/gestion-de-livraison.git/pkg/models"
	"github.com/stretchr/testify/assert"
)

// Test de la mise à jour de la météo du bateau
func TestUpdateBoatWeather(t *testing.T) {
	boat := &models.Boat{ID: "B123", Weather: "Clear"}
	UpdateBoatWeather(boat, "Stormy")

	assert.Equal(t, "Stormy", boat.Weather)
}

// Test de la recharge de la batterie du drone
func TestRechargeDrone(t *testing.T) {
	drone := &models.Drone{ID: "D456", Battery: 50}
	RechargeDrone(drone, 30)

	assert.Equal(t, 80, drone.Battery)

	RechargeDrone(drone, 50)
	assert.Equal(t, 100, drone.Battery)
}

// Test du chargement du camion
func TestLoadTruck(t *testing.T) {
	truck := &models.Truck{ID: "T123", Capacity: 50}

	err := LoadTruck(truck, 40)
	assert.NoError(t, err)
	assert.Equal(t, 90, truck.Capacity)

	err = LoadTruck(truck, 20)
	assert.Error(t, err)
	assert.Equal(t, "dépassement de la capacité", err.Error())
}
