package factory

import (
	"errors"

	"github.com/el-hadji-mamadou-sarr/gestion-de-livraison.git/pkg/interfaces"
	"github.com/el-hadji-mamadou-sarr/gestion-de-livraison.git/pkg/models"
)

func GetTransportMethod(method string) (interfaces.TransportMethod, error) {
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
