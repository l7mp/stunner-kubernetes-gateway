package implementation

import (
	"github.com/go-logr/logr"
	"sigs.k8s.io/gateway-api/apis/v1alpha2"

	"github.com/l7mp/stunner-kubernetes-gateway/internal/config"
	"github.com/l7mp/stunner-kubernetes-gateway/pkg/sdk/v1alpha2"
)

type gatewayClassImplementation struct {
	conf config.Config
}

func NewGatewayClassImplementation(conf config.Config) sdk.GatewayClassImpl {
	return &gatewayClassImplementation{
		conf: conf,
	}
}

func (impl *gatewayClassImplementation) Logger() logr.Logger {
	return impl.conf.Logger
}

func (impl *gatewayClassImplementation) ControllerName() string {
	return impl.conf.GatewayCtlrName
}

func (impl *gatewayClassImplementation) Upsert(gc *v1alpha2.GatewayClass) {
	if string(gc.Spec.ControllerName) != impl.ControllerName() {
		impl.Logger().Info("Wrong ControllerName in the GatewayClass resource",
			"expected", impl.ControllerName(),
			"got", gc.Spec.ControllerName)
		return
	}

	impl.Logger().Info("Processing GatewayClass resource",
		"name", gc.Name)
}

func (impl *gatewayClassImplementation) Remove(key string) {
	impl.Logger().Info("GatewayClass resource was removed",
		"name", key)
}
