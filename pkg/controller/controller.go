package controller

import (
	"github.com/TranThang-2804/k8s-pod-identity-controller/pkg/controller/reconciler"
	ctrl "sigs.k8s.io/controller-runtime"
	"os"
)

type CloudRoleController struct {
  mgr ctrl.Manager
  serviceAccountReconciler *reconciler.ServiceAccountReconciler
}

func NewCloudRoleController(mgr ctrl.Manager) (*CloudRoleController, error) {
  serviceAccountReconciler, err := reconciler.NewServiceAccountReconciler(mgr)
  if err != nil {
    return nil, err
  }
  return &CloudRoleController{
    mgr: mgr,
    serviceAccountReconciler: serviceAccountReconciler,
  }, nil
}

func (r *CloudRoleController) Start() error {
  logCtx := ctrl.Log.WithName("start controller")
  if err := r.serviceAccountReconciler.SetupWithManager(r.mgr); err != nil {
		logCtx.Error(err, "unable to setup controller with manager", "controller", "Pod")
		os.Exit(1)
	}
  return nil
}
