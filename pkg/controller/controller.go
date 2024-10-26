package controller

import (
  "github.com/TranThang-2804/k8s-pod-identity-controller/pkg/controller/reconciler"
)

type CloudRoleController struct {
  serviceAccountReconciler *ServiceAccountReconciler
}

func NewCloudRoleController(mgr ctrl.Manager)
