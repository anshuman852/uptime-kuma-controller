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

// MonitorReconciler reconciles a Monitor object
type MonitorReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=uptimekuma.uptime.kuma,resources=monitors,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=uptimekuma.uptime.kuma,resources=monitors/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=uptimekuma.uptime.kuma,resources=monitors/finalizers,verbs=update

func (r *MonitorReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	var monitor apiv1alpha1.Monitor
	if err := r.Get(ctx, req.NamespacedName, &monitor); err != nil {
		// Resource deleted or not found
		logger.Info(fmt.Sprintf("Monitor resource deleted or not found: %s", req.NamespacedName))
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	logger.Info("Reconciling Monitor resource",
		"name", monitor.Name,
		"namespace", monitor.Namespace,
		"url", monitor.Spec.URL,
		"type", monitor.Spec.Type,
		"interval", monitor.Spec.Interval,
		"uptimeKumaInstanceRef", monitor.Spec.UptimeKumaInstanceRef,
	)

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *MonitorReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&apiv1alpha1.Monitor{}).
		Complete(r)
}