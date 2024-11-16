package reconciler

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/TranThang-2804/k8s-pod-identity-controller/pkg/constants"
	"github.com/TranThang-2804/k8s-pod-identity-controller/pkg/provider"
	"github.com/TranThang-2804/k8s-pod-identity-controller/pkg/utils"
	"k8s.io/client-go/dynamic"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// Reconciler reconciles a SA object
type ServiceAccountReconciler struct {
	client.Client
	Scheme        *runtime.Scheme
	DynamicClient dynamic.Interface
}

func NewServiceAccountReconciler(mgr ctrl.Manager) (*ServiceAccountReconciler, error) {
	dynamicClient, err := dynamic.NewForConfig(mgr.GetConfig())
	if err != nil {
		return nil, fmt.Errorf("failed to create dynamic client: %w", err)
	}

	return &ServiceAccountReconciler{
		Client:        mgr.GetClient(),
		Scheme:        mgr.GetScheme(),
		DynamicClient: dynamicClient,
	}, nil
}

// Reconcile is part of the main Kubernetes reconciliation loop
func (r *ServiceAccountReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	sa := &corev1.ServiceAccount{}

	err := r.Get(ctx, req.NamespacedName, sa)
	if err != nil {
		if client.IgnoreNotFound(err) == nil {
			logger.Info("Service account not found. Ignoring since object must be deleted.", "service account", req.NamespacedName)
			return ctrl.Result{}, nil
		}
		logger.Error(err, "Failed to get Service Account", "service account", req.NamespacedName)
		return ctrl.Result{}, err
	}

	logger.Info("serviceAccountName", sa.Name)

	r.handleRoleIntegration(ctx, sa.Annotations)

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ServiceAccountReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&corev1.ServiceAccount{}).
		Complete(r)
}

func (r *ServiceAccountReconciler) handleRoleIntegration(ctx context.Context, annotations map[string]string) {
	logger := log.FromContext(ctx)

	providerList := annotations[constants.ENABLE_CLOUD_PROVIDER_Annotation]

  var providerClient provider.ProviderClient

	for _, cloudProvider := range utils.SplitAndRemoveWhitespace(&providerList) {
		switch cloudProvider {
    case string(constants.AWS):
      awsProviderClient, err := provider.NewAwsProviderClient(annotations[constants.AWS_ROLE_ANNOTATION])
      if err != nil {
        logger.Error(err, "Failed to create AWS provider client")
      }
      providerClient = awsProviderClient
    default:
      logger.Info("Invalid cloud provider", "cloudProvider", cloudProvider)
		}
	}

  providerClient.AssumeRole()
}
