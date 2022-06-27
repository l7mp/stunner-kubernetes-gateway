package event

import (
	"fmt"

	"github.com/l7mp/stunner-gateway-operator/internal/store"
)

// render event

type EventUpdate struct {
	Type           EventType
	GatewayClasses *store.GatewayClassStore
	Gateways       *store.GatewayStore
	UDPRoutes      *store.UDPRouteStore
	ConfigMaps     store.Store
}

// NewEvent returns an empty event
func NewEventUpdate() *EventUpdate {
	return &EventUpdate{
		Type:           EventTypeUpdate,
		GatewayClasses: store.NewGatewayClassStore(),
		Gateways:       store.NewGatewayStore(),
		UDPRoutes:      store.NewUDPRouteStore(),
		ConfigMaps:     store.NewStore(),
	}
}

func (e *EventUpdate) GetType() EventType {
	return e.Type
}

func (e *EventUpdate) String() string {
	return fmt.Sprintf("%s: #gway-classes: %d, #gways: %d, #udp-routes: %d, #configmaps: %d",
		e.Type.String(), e.GatewayClasses.Len(), e.Gateways.Len(), e.UDPRoutes.Len(),
		e.ConfigMaps.Len())
}
