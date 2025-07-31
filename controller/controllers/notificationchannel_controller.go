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

// NotificationChannelReconciler reconciles a NotificationChannel object
type NotificationChannelReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=uptimekuma.uptime.kuma,resources=notificationchannels,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=uptimekuma.uptime.kuma,resources=notificationchannels/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=uptimekuma.uptime.kuma,resources=notificationchannels/finalizers,verbs=update

/*
Reconcile reconciles a NotificationChannel resource.
This method is called when a change is detected for the resource.
*/
func (r *NotificationChannelReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    logger := log.FromContext(ctx)

	var channel apiv1alpha1.NotificationChannel
	if err := r.Get(ctx, req.NamespacedName, &channel); err != nil {
		// Resource deleted or not found
		logger.Info(fmt.Sprintf("NotificationChannel resource deleted or not found: %s", req.NamespacedName))
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	logger.Info("Reconciling NotificationChannel resource",
		"name", channel.Name,
		"namespace", channel.Namespace,
		"type", channel.Spec.Type,
		"channelName", channel.Spec.Name,
		"uptimeKumaInstanceRef", channel.Spec.UptimeKumaInstanceRef,
	)

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *NotificationChannelReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&apiv1alpha1.NotificationChannel{}).
		Complete(r)
}