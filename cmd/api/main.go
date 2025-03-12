package main

import (
	"fmt"

	"github.com/el-hadji-mamadou-sarr/gestion-de-livraison.git/pkg/factory"
	"github.com/el-hadji-mamadou-sarr/gestion-de-livraison.git/pkg/models"
	"github.com/el-hadji-mamadou-sarr/gestion-de-livraison.git/pkg/tracking"
	"github.com/el-hadji-mamadou-sarr/gestion-de-livraison.git/pkg/utils"
)

// Système de suivi

func main() {
	ch := make(chan string, 3) // Buffer pour éviter les blocages

	transports := []string{"truck", "drone", "boat"}

	for _, t := range transports {

		transport, err := factory.GetTransportMethod(t)

		if err != nil {
			fmt.Println("Erreur création transport:", err)
			continue
		}

		if truck, ok := transport.(*models.Truck); ok {
			utils.LoadTruck(truck, 20)
		}

		go tracking.TrackDelivery(transport, "Destination "+t, ch)

	}

	// Collecte des résultats
	for range transports {
		fmt.Println(<-ch)
	}
}
