package main

import (
	"os"
	"time"

	"github.com/TranThang-2804/k8s-pod-identity-controller/pkg/controller"
	"github.com/spf13/pflag"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/util/feature"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/klog/v2"
	"k8s.io/klog/v2/klogr"
	ctrl "sigs.k8s.io/controller-runtime"
)

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
)

func init() {
	_ = clientgoscheme.AddToScheme(scheme)
}

func main() {
	var metricsAddr string
	var enableLeaderElection bool
	var syncPeriod time.Duration
	var namespace string
	var controllerNamespace string

	pflag.BoolVar(&enableLeaderElection, "enable-leader-election", false, "Enable leader election for controller manager, this will ensure there is only one active controller manager.")
	pflag.DurationVar(&syncPeriod, "informer-re-sync-interval", 10*time.Second, "controller shared informer lister full re-sync period")
	pflag.StringVar(&metricsAddr, "metrics-addr", ":38080", "The address the metric endpoint binds to.")
	pflag.StringVar(&namespace, "namespace", "", "Namespace hehe to watch for resources, defaults to all namespaces")
	pflag.StringVar(&controllerNamespace, "controller-namespace", "", "Namespace to run the terraform jobs")
	feature.DefaultMutableFeatureGate.AddFlag(pflag.CommandLine)

	// embed klog
	klog.InitFlags(nil)
	pflag.Parse()

	ctrl.SetLogger(klogr.New())

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		LeaderElection:   enableLeaderElection,
		LeaderElectionID: "ce329a9c.core.oam.dev",
		// MetricsBindAddress: metricsAddr,
		// Namespace:          namespace,
		// Port:               9443,
		Scheme: scheme,
		// SyncPeriod:         &syncPeriod,
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

  controller, err := controller.NewCloudRoleController(mgr)
	if err != nil {
		setupLog.Error(err, "unable to create controller instance", "controller", "Pod")
		os.Exit(1)
	}
  
  controller.Start()

	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}

	setupLog.Info("starting manager")

	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}
