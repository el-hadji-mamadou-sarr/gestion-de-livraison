package factory

import (
	"errors"

	"github.com/el-hadji-mamadou-sarr/gestion-de-livraison.git/pkg/interfaces"
	"github.com/el-hadji-mamadou-sarr/gestion-de-livraison.git/pkg/models"
)
// permet de créer dynamiquement des instances de méthodes de transport
func GetTransportMethod(method string) (interfaces.TransportMethod, error) {
	//pour les différents cas
	switch method {
	case "truck":
		return &models.Truck{ID: "T123", Capacity: 10}, nil
	case "drone":
		return &models.Drone{ID: "D456", Battery: 100}, nil
	case "boat":
		return &models.Boat{ID: "B789", Weather: "Clear"}, nil
	default:
		return nil, errors.New("unknown transport method")
	}
}
