package tracking

import (
	"fmt"

	"github.com/el-hadji-mamadou-sarr/gestion-de-livraison.git/pkg/interfaces"
)

func TrackDelivery(transport interfaces.TransportMethod, destination string, ch chan string) {
	status, err := transport.DeliverPackage(destination)
	if err != nil {
		ch <- fmt.Sprintf("ERREUR (%T %s): %v", transport, transport.GetStatus(), err)
		return
	}
	ch <- fmt.Sprintf("SUCCÈS: %s | %s", status, transport.GetStatus())
}
