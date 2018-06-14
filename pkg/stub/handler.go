package stub

import (
	"context"

	"github.com/redhat-nfvpe/clearwater-sprout-cluster-operator/pkg/apis/projectclearwater/v1"

	"github.com/operator-framework/operator-sdk/pkg/sdk"
	"github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func NewHandler() sdk.Handler {
	return &Handler{}
}

type Handler struct {
	// Fill me
}

func (h *Handler) Handle(ctx context.Context, event sdk.Event) error {
	switch o := event.Object.(type) {
	case *v1.SproutCluster:
		//clusters : = o.Spec.Clusters;
		scale := o.Spec.Scale;

		for i := 0; i < scale; i++ {
			err := sdk.Create(newSproutPod(i, o))
			if err != nil && !errors.IsAlreadyExists(err) {
				logrus.Errorf("Failed to create sprout pod : %v", err)
				return err
			}
		}
	}
	return nil
}

func newSproutPod(index int, cr *v1.SproutCluster) *corev1.Pod {
	labels := map[string]string{
		"app": "sprout",
	}

	return &corev1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "sprout",
			Namespace: cr.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(cr, schema.GroupVersionKind{
					Group:   v1.SchemeGroupVersion.Group,
					Version: v1.SchemeGroupVersion.Version,
					Kind:    "SproutCluster",
				}),
			},
			Labels: labels,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:    "sprout",
					Image:   "busybox",
					Command: []string{"sleep", "3600"},
				},
			},
		},
	}
}

func newBonoPod(index int, cr *v1.SproutCluster) *corev1.Pod {
	labels := map[string]string{
		"app": "bono",
	}

	return &corev1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "bono",
			Namespace: cr.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(cr, schema.GroupVersionKind{
					Group:   v1.SchemeGroupVersion.Group,
					Version: v1.SchemeGroupVersion.Version,
					Kind:    "SproutCluster",
				}),
			},
			Labels: labels,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:    "bono",
					Image:   "busybox",
					Command: []string{"sleep", "3600"},
				},
			},
		},
	}
}
