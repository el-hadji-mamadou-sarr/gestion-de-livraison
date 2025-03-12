package models

import (
	"fmt"
	"time"
)

type Truck struct {
	ID       string
	Capacity int
}

func (t Truck) DeliverPackage(destination string) (string, error) {
	time.Sleep(3 * time.Second) // Simulation délai de livraison
	return fmt.Sprintf("Camion %s: Colis livré à %s", t.ID, destination), nil
}

func (t Truck) GetStatus() string {
	return fmt.Sprintf("Camion %s: En route (Capacité: %d)", t.ID, t.Capacity)
}
