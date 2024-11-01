package controller

import (
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"testing"
)

func TestNewCloudRoleController(t *testing.T) {

	scheme := runtime.NewScheme()
	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		LeaderElection:   false,
		LeaderElectionID: "ce329a9c.core.oam.dev",
		// MetricsBindAddress: metricsAddr,
		// Namespace:          namespace,
		// Port:               9443,
		Scheme: scheme,
		// SyncPeriod:         &syncPeriod,
	})
	if err != nil {
		t.Errorf("Can't create manager")
	}

	controller, err := NewCloudRoleController(mgr)
	if err != nil {
		t.Errorf("Can't create manager")
	}

  if controller.Start() != nil {
		t.Errorf("Can't start controller")
  }
}
