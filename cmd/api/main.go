package main

import (
	"errors"
	"fmt"
	"time"
)

// Interface TransportMethod
type TransportMethod interface {
	DeliverPackage(destination string) (string, error)
	GetStatus() string
}

// Implémentation des transports
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

type Drone struct {
	ID      string
	Battery int
}

func (d Drone) DeliverPackage(destination string) (string, error) {
	if d.Battery < 20 {
		return "", errors.New("batterie trop faible pour décoller")
	}
	time.Sleep(1 * time.Second) // Livraison rapide
	return fmt.Sprintf("Drone %s: Colis livré à %s", d.ID, destination), nil
}

func (d Drone) GetStatus() string {
	return fmt.Sprintf("Drone %s: Batterie à %d%%", d.ID, d.Battery)
}

type Boat struct {
	ID      string
	Weather string
}

func (b Boat) DeliverPackage(destination string) (string, error) {
	if b.Weather != "Clear" {
		return "", errors.New("conditions météo défavorables")
	}
	time.Sleep(5 * time.Second) // Livraison lente
	return fmt.Sprintf("Bateau %s: Colis livré à %s", b.ID, destination), nil
}

func (b Boat) GetStatus() string {
	return fmt.Sprintf("Bateau %s: Météo actuelle - %s", b.ID, b.Weather)
}

// Fabrique de transport
func GetTransportMethod(method string) (TransportMethod, error) {
	switch method {
	case "truck":
		return Truck{ID: "T123", Capacity: 10}, nil
	case "drone":
		return Drone{ID: "D456", Battery: 15}, // Testez avec 15% pour voir l'erreur
	case "boat":
		return Boat{ID: "B789", Weather: "Stormy"}, // Testez avec "Clear" pour succès
	default:
		return nil, errors.New("méthode de transport inconnue")
	}
}

// Système de suivi
func TrackDelivery(transport TransportMethod, destination string, ch chan string) {
	status, err := transport.DeliverPackage(destination)
	if err != nil {
		ch <- fmt.Sprintf("ERREUR (%T %s): %v", transport, transport.GetStatus(), err)
		return
	}
	ch <- fmt.Sprintf("SUCCÈS: %s | %s", status, transport.GetStatus())
}

func main() {
	ch := make(chan string, 3) // Buffer pour éviter les blocages

	transports := []string{"truck", "drone", "boat"}
	
	for _, t := range transports {
		transport, err := GetTransportMethod(t)
		if err != nil {
			fmt.Println("Erreur création transport:", err)
			continue
		}
		go TrackDelivery(transport, "Destination "+t, ch)
	}

	// Collecte des résultats
	for range transports {
		fmt.Println(<-ch)
	}
}