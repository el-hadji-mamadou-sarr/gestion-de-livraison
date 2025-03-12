package utils

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/el-hadji-mamadou-sarr/gestion-de-livraison.git/pkg/models"

	"github.com/manifoldco/promptui"
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
	fmt.Println("loading truck ...")
	if truck.Capacity+packages > 100 {
		return errors.New("dépassement de la capacité")
	}
	truck.Capacity += packages
	return nil
}

const statusFile = "status.log"

func LogStatus(status string) {
	f, err := os.OpenFile(statusFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error writing to status file:", err)
		return
	}
	defer f.Close()
	_, err = f.WriteString(status + "\n")
	if err != nil {
		fmt.Println("Error writing status:", err)
	}
}

// Reads the log file contents
func ReadLog(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// Clears the log file after reading
func ClearLog(filePath string) {
	os.WriteFile(filePath, []byte{}, 0644)
}

// Interactive menu for transport selection
func InteractiveMenu() string {
	prompt := promptui.Select{
		Label: "Select Transport Method",
		Items: []string{"Truck 🚚", "Drone ✈️", "Boat ⛵"},
	}
	_, result, err := prompt.Run()
	if err != nil {
		fmt.Println("Prompt failed", err)
		return ""
	}
	return strings.ToLower(strings.Split(result, " ")[0])
}
