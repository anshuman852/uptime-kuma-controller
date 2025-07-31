// SPDX-License-Identifier: MIT

package controllers

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	apiv1alpha1 "uptime-kuma-controller/api/v1alpha1"
)

// UptimeKumaInstanceReconciler reconciles a UptimeKumaInstance object
type UptimeKumaInstanceReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=uptimekuma.uptime-kuma-controller,resources=uptimekumainstances,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=uptimekuma.uptime-kuma-controller,resources=uptimekumainstances/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=uptimekuma.uptime-kuma-controller,resources=uptimekumainstances/finalizers,verbs=update

/*
Reconcile reconciles a UptimeKumaInstance resource.
This method is called when a change is detected for the resource.
*/
func (r *UptimeKumaInstanceReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    logger := log.FromContext(ctx)

	var instance apiv1alpha1.UptimeKumaInstance
	err := r.Get(ctx, req.NamespacedName, &instance)
	if err != nil {
		// Resource deleted or not found
		logger.Info("UptimeKumaInstance deleted or not found", "name", req.NamespacedName)
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	logger.Info(fmt.Sprintf("Reconciling UptimeKumaInstance: %s/%s", instance.Namespace, instance.Name))
	// No actual provisioning/connection logic yet

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *UptimeKumaInstanceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&apiv1alpha1.UptimeKumaInstance{}).
		Complete(r)
}