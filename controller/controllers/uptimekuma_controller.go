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

// UptimeKumaMonitorReconciler reconciles a UptimeKumaMonitor object
type UptimeKumaMonitorReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=uptimekuma.uptime.kuma,resources=uptimekumamonitors,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=uptimekuma.uptime.kuma,resources=uptimekumamonitors/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=uptimekuma.uptime.kuma,resources=uptimekumamonitors/finalizers,verbs=update

/*
Reconcile reconciles a UptimeKumaMonitor resource.
This method is called when a change is detected for the resource.
*/
func (r *UptimeKumaMonitorReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    logger := log.FromContext(ctx)

	var monitor apiv1alpha1.UptimeKumaMonitor
	if err := r.Get(ctx, req.NamespacedName, &monitor); err != nil {
		// Resource deleted or not found
		logger.Info("UptimeKumaMonitor deleted or not found", "name", req.NamespacedName)
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// Log creation or update
	logger.Info(fmt.Sprintf("Reconciling UptimeKumaMonitor: %s/%s", monitor.Namespace, monitor.Name),
		"url", monitor.Spec.URL,
		"interval", monitor.Spec.Interval,
		"name", monitor.Spec.Name,
	)

	// No actual Uptime Kuma API interaction yet
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *UptimeKumaMonitorReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&apiv1alpha1.UptimeKumaMonitor{}).
		Complete(r)
}