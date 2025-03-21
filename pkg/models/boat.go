package models

import (
	"errors"
	"fmt"
	"time"
)

type Boat struct {
	ID      string
	Weather string
}

func (b Boat) DeliverPackage(destination string) (string, error) {
	if b.Weather != "Clear" {
		return "", errors.New("conditions météo défavorables")
	}
	time.Sleep(5 * time.Second) // Livraison lente
	return fmt.Sprintf("⛴ Bateau %s: Colis livré à %s Date: %s", b.ID, destination, time.Now().Format("2006-01-02 15:04:05")), nil
}

func (b Boat) GetStatus() string {
	return fmt.Sprintf("Bateau: %s - Météo actuelle: %s", b.ID, b.Weather)
}
