package reconciler

import (
	"context"
	ctrl "sigs.k8s.io/controller-runtime"
)

type Reconciler interface {
	Reconcile(cts context.Context, req ctrl.Request) (ctrl.Result, error)
  SetupWithManager(mgr ctrl.Manager) error
}
