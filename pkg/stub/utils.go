package stub

import (
	"fmt"

	"github.com/operator-framework/operator-sdk/pkg/sdk"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

func newDeployment(name string, roleName string, replicas int32, t *metav1.TypeMeta, o *metav1.ObjectMeta) *appsv1.Deployment {
	l := map[string]string{
		"app": name,
		"role": roleName,
		"cluster": o.Name,
	}

	deployment := &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "Deployment",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: o.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: l,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: l,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Name:    roleName,
						Image:   "busybox",
						Command: []string{"sleep", "999999999"},
					}},
				},
			},
		},
	}

	addOwnerRefToObject(deployment, asOwner(t, o))
	return deployment
}

func updateDeploymentReplicas(deployment *appsv1.Deployment, replicas int32) error {
	err := sdk.Get(deployment)
	if err != nil {
		return fmt.Errorf("failed to get deployment: %v", err)
	}

	if *deployment.Spec.Replicas != replicas {
		deployment.Spec.Replicas = &replicas
		err = sdk.Update(deployment)
		if err != nil {
			return fmt.Errorf("failed to update deployment: %v", err)
		}
	}
	
	return nil
}

func getPodNames(roleName string, o *metav1.ObjectMeta) ([]string, error) {
	l := map[string]string{
		"cluster": o.Name,
		"role": roleName,
	}

	labelSelector := labels.SelectorFromSet(l).String()
	listOptions := &metav1.ListOptions{LabelSelector: labelSelector}

	podList := &corev1.PodList{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
	}
	
	err := sdk.List(o.Namespace, podList, sdk.WithListOptions(listOptions))
	if err != nil {
		return nil, fmt.Errorf("failed to list pods: %v", err)
	}

	podNames := make([]string, len(podList.Items))
	for i, pod := range podList.Items {
		podNames[i] = pod.Name
	}
	return podNames, nil
}

func addOwnerRefToObject(obj metav1.Object, ownerRef metav1.OwnerReference) {
	obj.SetOwnerReferences(append(obj.GetOwnerReferences(), ownerRef))
}

func asOwner(t *metav1.TypeMeta, o *metav1.ObjectMeta) metav1.OwnerReference {
	trueVar := true
	return metav1.OwnerReference{
		APIVersion: t.APIVersion,
		Kind:       t.Kind,
		Name:       o.Name,
		UID:        o.UID,
		Controller: &trueVar,
	}
}
