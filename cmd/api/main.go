package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/el-hadji-mamadou-sarr/gestion-de-livraison.git/pkg/factory"
	"github.com/el-hadji-mamadou-sarr/gestion-de-livraison.git/pkg/tracking"
	"github.com/el-hadji-mamadou-sarr/gestion-de-livraison.git/pkg/utils"
)

const statusFile = "status.log"

func main() {
	fmt.Println("📦 Delivery Tracking Simulation starting...")

	ch := make(chan string, 3) // Buffered channel
	var wg sync.WaitGroup      // WaitGroup to wait for deliveries

	transports := []string{"truck", "drone", "boat"}

	// Launch tracking for each transport
	for _, t := range transports {
		transport, err := factory.GetTransportMethod(t)
		if err != nil {
			fmt.Println("Erreur création transport:", err)
			continue
		}
		wg.Add(1)
		go func(transportType string) {
			defer wg.Done()
			tracking.TrackDelivery(transport, "Destination "+transportType, ch)
		}(t)
	}

	// Collect delivery results
	go func() {
		for range transports {
			fmt.Println(<-ch)
		}
	}()

	// Wait for all deliveries to complete
	wg.Wait()
	fmt.Println("📦 Delivery Tracking Simulation ended.")

	utils.ClearLog(statusFile)
	// Now start reading status updates
	readStatus()
}

// Reads delivery updates periodically after simulation ends
func readStatus() {
	fmt.Println("📦 Delivery Tracking System Running...")

	for {
		data, err := utils.ReadLog(statusFile)
		if err == nil && len(data) > 0 {
			fmt.Println("\n🚀 Live Delivery Updates 🚀")
			fmt.Println(data)
			utils.ClearLog(statusFile) // Prevent duplicate messages
		}
		time.Sleep(5 * time.Second)
	}
}
