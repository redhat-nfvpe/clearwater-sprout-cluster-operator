package stub

import (
	"context"
	"fmt"
	"reflect"

	"github.com/redhat-nfvpe/clearwater-sprout-cluster-operator/pkg/apis/projectclearwater/v1"

	"github.com/operator-framework/operator-sdk/pkg/sdk"
	appsv1 "k8s.io/api/apps/v1"
	"github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/api/errors"
)

type Handler struct {
}

func NewHandler() sdk.Handler {
	return &Handler{}
}

func (h *Handler) Handle(ctx context.Context, event sdk.Event) error {
	if sproutCluster, ok := event.Object.(*v1.SproutCluster); ok {
		// Ignore the delete event since the garbage collector will clean up all secondary resources for the CR
		// All secondary resources must have the CR set as their OwnerReference for this to be the case
		if event.Deleted {
			return nil
		}

		shards := sproutCluster.Spec.Shards
		for s := int32(0); s < shards; s++ {
			deployment := newShardDeployment(s, sproutCluster)		
		
			if err := sdk.Create(deployment); (err != nil) && !errors.IsAlreadyExists(err) {
				logrus.Errorf("failed to create shard deployment: %v", err)
				return err
			}

			if err := updateShardDeployment(deployment, sproutCluster); err != nil {
				return err
			}

			deployment = newBonoDeployment(s, sproutCluster)		

			if err := sdk.Create(deployment); (err != nil) && !errors.IsAlreadyExists(err) {
				logrus.Errorf("failed to create bono deployment: %v", err)
				return err
			}

			var shardNames []string
			var err error
			if shardNames, err = getPodNames("sprout", &sproutCluster.ObjectMeta); err != nil {
				return err
			}
			var bonoNames []string
			if bonoNames, err = getPodNames("bono", &sproutCluster.ObjectMeta); err != nil {
				return err
			}

			if !reflect.DeepEqual(shardNames, sproutCluster.Status.ShardNodes) || !reflect.DeepEqual(bonoNames, sproutCluster.Status.BonoNodes) {
				sproutCluster.Status.ShardNodes = shardNames;
				sproutCluster.Status.BonoNodes = bonoNames;
				if err := sdk.Update(sproutCluster); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

//
// Shard deployment
//

func newShardDeployment(index int32, sproutCluster *v1.SproutCluster) *appsv1.Deployment {
	name := fmt.Sprintf("%s-shard-%d", sproutCluster.Name, index)
	return newDeployment(name, "sprout", sproutCluster.Spec.Scale, &sproutCluster.TypeMeta, &sproutCluster.ObjectMeta)
}

func updateShardDeployment(deployment *appsv1.Deployment, sproutCluster *v1.SproutCluster) error {
	return updateDeploymentReplicas(deployment, sproutCluster.Spec.Scale)
}

//
// Bono deployment
//

func newBonoDeployment(index int32, sproutCluster *v1.SproutCluster) *appsv1.Deployment {
	name := fmt.Sprintf("%s-shard-%d-edge-proxy", sproutCluster.Name, index)
	return newDeployment(name, "bono", int32(1), &sproutCluster.TypeMeta, &sproutCluster.ObjectMeta)
}
