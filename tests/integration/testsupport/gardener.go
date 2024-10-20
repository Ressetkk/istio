package testsupport

import (
	"context"
	"github.com/kyma-project/istio/operator/internal/clusterconfig"
	corev1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func RunsOnGardener(ctx context.Context, k8sClient client.Client) (bool, error) {
	cmShootInfo := corev1.ConfigMap{}
	err := k8sClient.Get(ctx, types.NamespacedName{Namespace: clusterconfig.ConfigMapShootInfoNS, Name: clusterconfig.ConfigMapShootInfoName}, &cmShootInfo)

	if k8serrors.IsNotFound(err) {
		return false, nil
	}

	return true, err
}
