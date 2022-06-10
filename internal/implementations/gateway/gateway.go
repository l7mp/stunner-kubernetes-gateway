package implementation

import (
	"github.com/go-logr/logr"
	"sigs.k8s.io/gateway-api/apis/v1alpha2"

	"github.com/l7mp/stunner-kubernetes-gateway/internal/config"
	"github.com/l7mp/stunner-kubernetes-gateway/pkg/sdk/v1alpha2"
)

type gatewayImplementation struct {
	conf config.Config
}

func NewGatewayImplementation(conf config.Config) sdk.GatewayImpl {
	return &gatewayImplementation{
		conf: conf,
	}
}

func (impl *gatewayImplementation) Logger() logr.Logger {
	return impl.conf.Logger
}

func (impl *gatewayImplementation) ControllerName() string {
	return impl.conf.GatewayCtlrName
}

func (impl *gatewayImplementation) Upsert(gw *v1alpha2.Gateway) {
	if gw.Name == impl.ControllerName() {
		impl.Logger().Info("Found correct Gateway resource",
			"name", gw.Name,
		)
		return
	}
}

func (impl *gatewayImplementation) Remove(key string) {
	impl.Logger().Info("Gateway resource was removed",
		"name", key,
	)
}
