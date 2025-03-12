package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/el-hadji-mamadou-sarr/gestion-de-livraison.git/pkg/factory"
	"github.com/el-hadji-mamadou-sarr/gestion-de-livraison.git/pkg/utils"
	"github.com/manifoldco/promptui"
)

// Système de suivi

// func main() {
// 	ch := make(chan string, 3) // Buffer pour éviter les blocages

// 	transports := []string{"truck", "drone", "boat"}

// 	for _, t := range transports {
// 		transport, err := factory.GetTransportMethod(t)
// 		if err != nil {
// 			fmt.Println("Erreur création transport:", err)
// 			continue
// 		}
// 		go tracking.TrackDelivery(transport, "Destination "+t, ch)
// 	}

// 	// Collecte des résultats
// 	for range transports {
// 		fmt.Println(<-ch)
// 	}
// }

func main() {
	for {
		transportType := utils.InteractiveMenu()
		if transportType == "" {
			continue
		}
		transport, err := factory.GetTransportMethod(transportType)
		if err != nil {
			fmt.Println("Error creating transport:", err)
			continue
		}

		statusMessage := fmt.Sprintf("%s is delivering your package...\n", transport.GetStatus())
		fmt.Print(statusMessage)
		utils.LogStatus(statusMessage)

		ch := make(chan string)
		go func() {
			status, err := transport.DeliverPackage("Destination")
			if err != nil {
				ch <- fmt.Sprintf("Delivery failed: %v ❌", err)
			} else {
				ch <- status
			}
		}()

		fmt.Println("Delivery in progress... ⏳")
		time.Sleep(time.Second * 2)

		result := <-ch
		fmt.Println(result)
		utils.LogStatus(result)

		prompt := promptui.Prompt{
			Label: "Do you want to make another delivery? (yes/no)",
		}
		response, _ := prompt.Run()
		if strings.ToLower(response) != "yes" {
			break
		}
	}
}
