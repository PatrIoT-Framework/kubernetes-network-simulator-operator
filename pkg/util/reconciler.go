package util

import (
	"context"
	"github.com/go-logr/logr"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type ReconcilerBase struct {
	Client client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

func NewRencolicerBase(client client.Client, log logr.Logger, scheme *runtime.Scheme) ReconcilerBase {
	return ReconcilerBase{
		Client: client,
		Log:    log,
		Scheme: scheme,
	}
}

func (r *ReconcilerBase) IsValid(obj metav1.Object) (bool, error) {
	return true, nil
}

func (r *ReconcilerBase) IsInitialized(obj metav1.Object) bool {
	return true
}

func (r *ReconcilerBase) Reconcile(ctx context.Context, req ctrl.Request) (reconcile.Result, error) {
	return reconcile.Result{}, nil
}

func (r *ReconcilerBase) GetClient() client.Client {
	return r.Client
}

func (r *ReconcilerBase) GetScheme() *runtime.Scheme {
	return r.Scheme
}

// getNamespace returns namespace for given name
func (r ReconcilerBase) GetNamespace(name string, ctx context.Context) (*v1.Namespace, error) {
	namespacedName := types.NamespacedName{
		Name: name,
	}

	var namespace v1.Namespace
	if err := r.GetClient().Get(ctx, namespacedName, &namespace); err != nil {
		return nil, err
	}
	return &namespace, nil
}
