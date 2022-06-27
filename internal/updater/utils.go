package updater

// updater uploads client updates
import (
	// "context"
	"fmt"
	// "reflect"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	// "k8s.io/apimachinery/pkg/runtime"
	// "sigs.k8s.io/controller-runtime/pkg/manager"
	// ctlr "sigs.k8s.io/controller-runtime"
	// "sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/client"
	// corev1 "k8s.io/api/core/v1"
	ctrlutil "sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	gatewayv1alpha2 "sigs.k8s.io/gateway-api/apis/v1alpha2"

	// "github.com/l7mp/stunner-gateway-operator/internal/event"
	// "github.com/l7mp/stunner-gateway-operator/internal/operator"
	"github.com/l7mp/stunner-gateway-operator/internal/store"
)

func (u *Updater) updateGatewayClass(o client.Object) (ctrlutil.OperationResult, error) {
	u.log.V(2).Info("updating gateway class", "resource",
		store.GetObjectKey(o))

	gc, ok := o.(*gatewayv1alpha2.GatewayClass)
	if ok == false {
		return ctrlutil.OperationResultNone,
			fmt.Errorf("internal error: cannot cast object %q to gateway-class",
				store.GetObjectKey(o))
	}

	client := u.manager.GetClient()

	// we use create-or-update: the create branch is useless but at least it does the
	// get/update cycle for us
	current := &gatewayv1alpha2.GatewayClass{ObjectMeta: metav1.ObjectMeta{
		Name:      gc.GetName(),
		Namespace: gc.GetNamespace(),
	}}
	op, err := ctrlutil.CreateOrUpdate(u.ctx, client, current, func() error {
		// the only thing we change on gatewayclasses is the status: copy
		gc.Status.DeepCopyInto(&current.Status)
		return nil
	})

	if err != nil {
		return ctrlutil.OperationResultNone, fmt.Errorf("cannot update gatewayclass %q: %w",
			store.GetObjectKey(gc), err)
	}

	u.log.V(1).Info("gateway-class updated", "resource", store.GetObjectKey(gc), "result",
		current)

	return op, nil
}

func (u *Updater) updateGateway(o client.Object) (ctrlutil.OperationResult, error) {
	u.log.V(2).Info("updating gateway", "resource", store.GetObjectKey(o))

	gw, ok := o.(*gatewayv1alpha2.Gateway)
	if ok == false {
		return ctrlutil.OperationResultNone,
			fmt.Errorf("internal error: cannot cast %q object to gateway",
				store.GetObjectKey(o))
	}

	client := u.manager.GetClient()

	// we use create-or-update: the create branch is useless but at least it does the
	// get/update cycle for us
	current := &gatewayv1alpha2.Gateway{ObjectMeta: metav1.ObjectMeta{
		Name:      gw.GetName(),
		Namespace: gw.GetNamespace(),
	}}
	op, err := ctrlutil.CreateOrUpdate(u.ctx, client, current, func() error {
		// the only thing we change on gateways is the status: copy
		gw.Status.DeepCopyInto(&current.Status)
		return nil
	})

	if err != nil {
		return ctrlutil.OperationResultNone, fmt.Errorf("cannot update gateway %q: %w",
			store.GetObjectKey(gw), err)
	}

	u.log.V(1).Info("gateway updated", "resource", store.GetObjectKey(gw), "result",
		current)

	return op, nil
}

func (u *Updater) updateUDPRoute(o client.Object) (ctrlutil.OperationResult, error) {
	u.log.V(2).Info("updating UDP-route", "resource", store.GetObjectKey(o))

	ro, ok := o.(*gatewayv1alpha2.UDPRoute)
	if ok == false {
		return ctrlutil.OperationResultNone,
			fmt.Errorf("internal error: cannot cast object %q to UDP route",
				store.GetObjectKey(o))
	}

	client := u.manager.GetClient()

	// we use create-or-update: the create branch is useless but at least it does the
	// get/update cycle for us
	current := &gatewayv1alpha2.UDPRoute{ObjectMeta: metav1.ObjectMeta{
		Name:      ro.GetName(),
		Namespace: ro.GetNamespace(),
	}}
	op, err := ctrlutil.CreateOrUpdate(u.ctx, client, current, func() error {
		// the only thing we change on UDP routes is the status: copy
		ro.Status.DeepCopyInto(&current.Status)
		return nil
	})

	if err != nil {
		return ctrlutil.OperationResultNone, fmt.Errorf("cannot update UDP-route %q: %w",
			store.GetObjectKey(ro), err)
	}

	u.log.V(1).Info("UDP-route updated", "resource", store.GetObjectKey(ro), "result",
		current)

	return op, nil
}

func (u *Updater) updateConfigMap(o client.Object) (ctrlutil.OperationResult, error) {
	u.log.V(2).Info("updating config-map", "resource", store.GetObjectKey(o))

	cm, ok := o.(*corev1.ConfigMap)
	if ok == false {
		return ctrlutil.OperationResultNone,
			fmt.Errorf("internal error: cannot cast object %q to config-map",
				store.GetObjectKey(o))
	}

	client := u.manager.GetClient()

	current := &corev1.ConfigMap{ObjectMeta: metav1.ObjectMeta{
		Name:      cm.GetName(),
		Namespace: cm.GetNamespace(),
	}}

	op, err := ctrlutil.CreateOrUpdate(u.ctx, client, current, func() error {
		// thing that might have been changed by the controler: the owner ref and the data

		u.log.Info("before", "cm", fmt.Sprintf("%#v\n", current))

		current.SetOwnerReferences(cm.GetOwnerReferences())
		current.Data = make(map[string]string)
		for k, v := range cm.Data {
			current.Data[k] = v
		}

		u.log.Info("after", "cm", fmt.Sprintf("%#v\n", current))
		return nil
	})

	if err != nil {
		return ctrlutil.OperationResultNone, fmt.Errorf("cannot update STUNNer config-map %q: %w",
			store.GetObjectKey(cm), err)
	}

	u.log.V(1).Info("STUNner config-map  updated", "resource", store.GetObjectKey(cm), "result",
		current)

	return op, nil
}