package utils

import (
	"errors"

	"github.com/el-hadji-mamadou-sarr/gestion-de-livraison.git/pkg/models"
)

func UpdateBoatWeather(boat *models.Boat, weather string) {
	boat.Weather = weather
}

// Recharge drone battery
func RechargeDrone(drone *models.Drone, amount int) {
	drone.Battery += amount
	if drone.Battery > 100 {
		drone.Battery = 100
	}
}

// Load packages into truck
func LoadTruck(truck *models.Truck, packages int) error {
	if truck.Capacity+packages > truck.Capacity {
		return errors.New("dépassement de la capacité")
	}
	truck.Capacity += packages
	return nil
}
