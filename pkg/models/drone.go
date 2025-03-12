package models

import (
	"errors"
	"fmt"
	"time"
)

type Drone struct {
	ID      string
	Battery int
}

func (d Drone) DeliverPackage(destination string) (string, error) {
	if d.Battery < 20 {
		return "", errors.New("⚡ Batterie trop faible pour décoller")
	}
	time.Sleep(1 * time.Second) // Livraison rapide
	// d.Battery -= 20 // Simulate battery consumption
	return fmt.Sprintf("🚀 Drone :%s - Colis livré à %s - Date: %s", d.ID, destination, time.Now().Format("2006-01-02 15:04:05")), nil
}

func (d Drone) GetStatus() string {
	return fmt.Sprintf("🚀 Drone : %s - Batterie à %d%%", d.ID, d.Battery)
}

