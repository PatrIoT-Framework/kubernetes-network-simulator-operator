/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"github.com/go-logr/logr"
	v1 "k8s.io/api/core/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	networksimulatorv1 "github.com/jsmadis/kubernetes-network-simulator-operator/api/v1"
)

// DeviceReconciler reconciles a Device object
type DeviceReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=network-simulator.patriot-framework.io,resources=devices,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=network-simulator.patriot-framework.io,resources=devices/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=network-simulator.patriot-framework.io,resources=devices/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Device object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.7.0/pkg/reconcile
func (r *DeviceReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := r.Log.WithValues("device", req.NamespacedName)

	var device networksimulatorv1.Device
	if err := r.Get(ctx, req.NamespacedName, &device); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	constructPodForDevice := func(device *networksimulatorv1.Device) (*v1.Pod, error) {
		name := device.Name + "-pod"
		pod := &v1.Pod{
			ObjectMeta: v12.ObjectMeta{
				Labels:      make(map[string]string),
				Annotations: make(map[string]string),
				Name:        name,
				Namespace:   device.Namespace,
			},
			Spec: *device.Spec.PodTemplate.Spec.DeepCopy(),
		}
		if err := ctrl.SetControllerReference(device, pod, r.Scheme); err != nil {
			return nil, err
		}
		return pod, nil
	}

	if !device.Spec.Active {
		// TODO: suspend device?
		log.V(1).Info("Not active device")
		return ctrl.Result{}, nil
	}

	pod, err := constructPodForDevice(&device)

	if err != nil {
		log.V(1).Info("Failed to create pod for device")
		return ctrl.Result{}, err
	}

	if err := r.Create(ctx, pod); err != nil {
		log.Error(err, "unable to create Pod for device", "pod", pod)
		return ctrl.Result{}, err
	}

	log.V(1).Info("Created pod")

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *DeviceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&networksimulatorv1.Device{}).
		Complete(r)
}
